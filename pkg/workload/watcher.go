package workload

import (
	"context"
	"math"
	"sync"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	apps "k8s.io/api/apps/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	"k8s.io/kubectl/pkg/util/deployment"

	"github.com/datawire/dlib/dlog"
	"github.com/datawire/k8sapi/pkg/k8sapi"
	"github.com/telepresenceio/telepresence/v2/pkg/informer"
)

type EventType int

const (
	EventTypeAdd = iota
	EventTypeUpdate
	EventTypeDelete
)

type WorkloadEvent struct {
	Type     EventType
	Workload k8sapi.Workload
}

func (e EventType) String() string {
	switch e {
	case EventTypeAdd:
		return "add"
	case EventTypeUpdate:
		return "update"
	case EventTypeDelete:
		return "delete"
	default:
		return "unknown"
	}
}

type Watcher interface {
	Subscribe(ctx context.Context) <-chan []WorkloadEvent
}

type watcher struct {
	sync.Mutex
	namespace       string
	subscriptions   map[uuid.UUID]chan<- []WorkloadEvent
	timer           *time.Timer
	events          []WorkloadEvent
	rolloutsEnabled bool
}

func NewWatcher(ctx context.Context, ns string, rolloutsEnabled bool) (Watcher, error) {
	w := new(watcher)
	w.namespace = ns
	w.rolloutsEnabled = rolloutsEnabled
	w.subscriptions = make(map[uuid.UUID]chan<- []WorkloadEvent)
	w.timer = time.AfterFunc(time.Duration(math.MaxInt64), func() {
		w.Lock()
		ss := make([]chan<- []WorkloadEvent, len(w.subscriptions))
		i := 0
		for _, sub := range w.subscriptions {
			ss[i] = sub
			i++
		}
		events := w.events
		w.events = nil
		w.Unlock()
		for _, s := range ss {
			select {
			case <-ctx.Done():
				return
			case s <- events:
			}
		}
	})

	err := w.addEventHandler(ctx, ns)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func hasValidReplicasetOwner(wl k8sapi.Workload, rolloutsEnabled bool) bool {
	for _, ref := range wl.GetOwnerReferences() {
		if ref.Controller != nil && *ref.Controller {
			switch ref.Kind {
			case "Deployment":
				return true
			case "Rollout":
				if rolloutsEnabled {
					return true
				}
			}
		}
	}
	return false
}

func (w *watcher) Subscribe(ctx context.Context) <-chan []WorkloadEvent {
	ch := make(chan []WorkloadEvent, 1)
	initialEvents := make([]WorkloadEvent, 0, 100)
	id := uuid.New()
	kf := informer.GetFactory(ctx, w.namespace)
	ai := kf.GetK8sInformerFactory().Apps().V1()
	dlog.Debugf(ctx, "workload.Watcher producing initial events for namespace %s", w.namespace)
	if dps, err := ai.Deployments().Lister().Deployments(w.namespace).List(labels.Everything()); err == nil {
		for _, obj := range dps {
			if wl, ok := FromAny(obj); ok && !hasValidReplicasetOwner(wl, w.rolloutsEnabled) {
				initialEvents = append(initialEvents, WorkloadEvent{
					Type:     EventTypeAdd,
					Workload: wl,
				})
			}
		}
	}
	if rps, err := ai.ReplicaSets().Lister().ReplicaSets(w.namespace).List(labels.Everything()); err == nil {
		for _, obj := range rps {
			if wl, ok := FromAny(obj); ok && !hasValidReplicasetOwner(wl, w.rolloutsEnabled) {
				initialEvents = append(initialEvents, WorkloadEvent{
					Type:     EventTypeAdd,
					Workload: wl,
				})
			}
		}
	}
	if sps, err := ai.StatefulSets().Lister().StatefulSets(w.namespace).List(labels.Everything()); err == nil {
		for _, obj := range sps {
			if wl, ok := FromAny(obj); ok && !hasValidReplicasetOwner(wl, w.rolloutsEnabled) {
				initialEvents = append(initialEvents, WorkloadEvent{
					Type:     EventTypeAdd,
					Workload: wl,
				})
			}
		}
	}
	if w.rolloutsEnabled {
		ri := kf.GetArgoRolloutsInformerFactory().Argoproj().V1alpha1()
		if sps, err := ri.Rollouts().Lister().Rollouts(w.namespace).List(labels.Everything()); err == nil {
			for _, obj := range sps {
				if wl, ok := FromAny(obj); ok && !hasValidReplicasetOwner(wl, w.rolloutsEnabled) {
					initialEvents = append(initialEvents, WorkloadEvent{
						Type:     EventTypeAdd,
						Workload: wl,
					})
				}
			}
		}
	}
	ch <- initialEvents

	w.Lock()
	w.subscriptions[id] = ch
	w.Unlock()
	go func() {
		<-ctx.Done()
		close(ch)
		w.Lock()
		delete(w.subscriptions, id)
		w.Unlock()
	}()
	return ch
}

func compareOptions() []cmp.Option {
	return []cmp.Option{
		// Ignore frequently changing fields of no interest
		cmpopts.IgnoreFields(meta.ObjectMeta{}, "Namespace", "ResourceVersion", "Generation", "ManagedFields"),

		// Only the Conditions are of interest in the DeploymentStatus.
		cmp.Comparer(func(a, b apps.DeploymentStatus) bool {
			// Only compare the DeploymentCondition's type and status
			return cmp.Equal(a.Conditions, b.Conditions, cmp.Comparer(func(c1, c2 apps.DeploymentCondition) bool {
				return c1.Type == c2.Type && c1.Status == c2.Status
			}))
		}),

		// Treat a nil map or slice as empty.
		cmpopts.EquateEmpty(),

		// Ignore frequently changing annotations of no interest.
		cmpopts.IgnoreMapEntries(func(k, _ string) bool {
			return k == AnnRestartedAt || k == deployment.RevisionAnnotation
		}),
	}
}

func (w *watcher) watch(ix cache.SharedIndexInformer, ns string, hasValidController func(k8sapi.Workload) bool) error {
	_, err := ix.AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj any) {
				if wl, ok := FromAny(obj); ok && ns == wl.GetNamespace() && !hasValidController(wl) {
					w.handleEvent(WorkloadEvent{Type: EventTypeAdd, Workload: wl})
				}
			},
			DeleteFunc: func(obj any) {
				if wl, ok := FromAny(obj); ok {
					if ns == wl.GetNamespace() && !hasValidController(wl) {
						w.handleEvent(WorkloadEvent{Type: EventTypeDelete, Workload: wl})
					}
				} else if dfsu, ok := obj.(*cache.DeletedFinalStateUnknown); ok {
					if wl, ok = FromAny(dfsu.Obj); ok && ns == wl.GetNamespace() && !hasValidController(wl) {
						w.handleEvent(WorkloadEvent{Type: EventTypeDelete, Workload: wl})
					}
				}
			},
			UpdateFunc: func(oldObj, newObj any) {
				if wl, ok := FromAny(newObj); ok && ns == wl.GetNamespace() && !hasValidController(wl) {
					if oldWl, ok := FromAny(oldObj); ok {
						if cmp.Equal(wl, oldWl, compareOptions()...) {
							return
						}
						// Replace the cmp.Equal above with this to view the changes that trigger an update:
						//
						// diff := cmp.Diff(wl, oldWl, compareOptions()...)
						// if diff == "" {
						//   return
						// }
						// dlog.Debugf(ctx, "DIFF:\n%s", diff)
						w.handleEvent(WorkloadEvent{Type: EventTypeUpdate, Workload: wl})
					}
				}
			},
		})
	return err
}

func (w *watcher) addEventHandler(ctx context.Context, ns string) error {
	kf := informer.GetFactory(ctx, ns)
	hvc := func(wl k8sapi.Workload) bool {
		return hasValidReplicasetOwner(wl, w.rolloutsEnabled)
	}

	ai := kf.GetK8sInformerFactory().Apps().V1()
	if err := w.watch(ai.Deployments().Informer(), ns, hvc); err != nil {
		return err
	}
	if err := w.watch(ai.ReplicaSets().Informer(), ns, hvc); err != nil {
		return err
	}
	if err := w.watch(ai.StatefulSets().Informer(), ns, hvc); err != nil {
		return err
	}
	if !w.rolloutsEnabled {
		dlog.Infof(ctx, "Argo Rollouts is disabled, Argo Rollouts will not be watched")
	} else {
		ri := kf.GetArgoRolloutsInformerFactory().Argoproj().V1alpha1()
		if err := w.watch(ri.Rollouts().Informer(), ns, hvc); err != nil {
			return err
		}
	}
	return nil
}

func (w *watcher) handleEvent(we WorkloadEvent) {
	w.Lock()
	w.events = append(w.events, we)
	w.Unlock()

	// Defers sending until things been quiet for a while
	w.timer.Reset(5 * time.Millisecond)
}
