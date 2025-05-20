// Copyright 2025 Karel Van Hecke
// SPDX-License-Identifier: Apache-2.0

package policy

import "crypto/x509"

type Request struct {
	CSR  *x509.CertificateRequest
	Node Node
}
