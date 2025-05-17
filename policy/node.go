package policy

import "net"

type Node struct {
	DNSNames    []string
	IPAddresses []net.IP
}
