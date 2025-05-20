// Copyright 2025 Karel Van Hecke
// SPDX-License-Identifier: Apache-2.0

package policy

import "net"

type Node struct {
	DNSNames    []string
	IPAddresses []net.IP
}
