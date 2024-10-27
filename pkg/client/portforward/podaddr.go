package portforward

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"google.golang.org/grpc/resolver"
)

type podAddress struct {
	fromSvc   bool
	name      string
	namespace string
	port      uint16
}

func parseAddr(addr string) (kind, name, namespace, port string, err error) {
	if slash := strings.Index(addr, "/"); slash < 0 {
		kind = "pod"
	} else {
		kind = addr[:slash]
		addr = addr[slash+1:]
	}
	if name, port, err = net.SplitHostPort(addr); err == nil {
		var namespace string
		if dot := strings.LastIndex(name, "."); dot > 0 {
			namespace = name[dot+1:]
			name = name[:dot]
		}
		return kind, name, namespace, port, nil
	}
	return "", "", "", "", fmt.Errorf("%q is not a valid [<kind>/]<name[.namespace]>:<port-number>", addr)
}

func parsePodAddr(addr string) (podAddress, error) {
	kind, name, namespace, port, err := parseAddr(addr)
	if err != nil {
		return podAddress{}, err
	}
	if kind == "pod" {
		if pn, err := strconv.ParseUint(port, 10, 16); err == nil {
			return podAddress{
				name:      name,
				namespace: namespace,
				port:      uint16(pn),
			}, nil
		}
	}
	return podAddress{}, fmt.Errorf("%q is not a valid pod port address", addr)
}

func (pa *podAddress) String() string {
	return fmt.Sprintf("%s.%s:%d", pa.name, pa.namespace, pa.port)
}

func (pa *podAddress) state() resolver.State {
	return resolver.State{Addresses: []resolver.Address{{Addr: pa.String()}}}
}
