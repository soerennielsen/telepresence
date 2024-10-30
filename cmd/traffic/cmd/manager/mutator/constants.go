package mutator

import (
	"github.com/telepresenceio/telepresence/v2/pkg/agentconfig"
	"github.com/telepresenceio/telepresence/v2/pkg/workload"
)

const (
	InjectAnnotation       = workload.DomainPrefix + "inject-" + agentconfig.ContainerName
	ServiceNameAnnotation  = workload.DomainPrefix + "inject-service-name"
	ManualInjectAnnotation = workload.DomainPrefix + "manually-injected"
)
