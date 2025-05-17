package policy

import "crypto/x509"

type Request struct {
	CSR  *x509.CertificateRequest
	Node Node
}
