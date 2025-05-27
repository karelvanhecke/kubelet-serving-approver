// Copyright 2025 Karel Van Hecke
// SPDX-License-Identifier: Apache-2.0

package controller_test

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"net"
	"slices"
	"testing"

	"github.com/karelvanhecke/kubelet-serving-approver/controller"
	policy "github.com/karelvanhecke/kubelet-serving-approver/policy/fake"
	certv1 "k8s.io/api/certificates/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	client "sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func assertCSRCondition(policyResult bool, cType certv1.RequestConditionType) bool {
	nodeHostname := "node.example.org"
	nodeExternalDNS := "node.ext.example.org"
	nodeInternalDNS := "node.int.example.org"
	nodeInternalIP := "192.0.2.1"
	nodeExternalIP := "2001:db8::"
	csrName := "csr-test"

	pk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	csrData, err := x509.CreateCertificateRequest(rand.Reader, &x509.CertificateRequest{
		DNSNames: []string{
			nodeHostname,
			nodeExternalDNS,
			nodeInternalDNS,
		},
		IPAddresses: []net.IP{
			net.ParseIP(nodeInternalIP),
			net.ParseIP(nodeExternalIP),
		},
	}, pk)
	if err != nil {
		panic(err)
	}
	cl := client.NewFakeClient([]runtime.Object{
		&corev1.Node{
			ObjectMeta: metav1.ObjectMeta{
				Name: nodeHostname,
			},
			Status: corev1.NodeStatus{
				Addresses: []corev1.NodeAddress{
					{
						Type:    corev1.NodeHostName,
						Address: nodeHostname,
					},
					{
						Type:    corev1.NodeExternalDNS,
						Address: nodeExternalDNS,
					},
					{
						Type:    corev1.NodeInternalDNS,
						Address: nodeInternalDNS,
					},
					{
						Type:    corev1.NodeInternalIP,
						Address: nodeInternalIP,
					},
					{
						Type:    corev1.NodeExternalIP,
						Address: nodeExternalIP,
					},
				},
			},
		},
		&certv1.CertificateSigningRequest{
			ObjectMeta: metav1.ObjectMeta{
				Name: csrName,
			},
			Spec: certv1.CertificateSigningRequestSpec{
				Request: pem.EncodeToMemory(&pem.Block{
					Type:  "CERTIFICATE REQUEST",
					Bytes: csrData,
				}),
				SignerName: "kubernetes.io/kubelet-serving",
				Username:   "system:node:" + nodeHostname,
			},
		},
	}...)

	cr := controller.CSRReconciler{Client: cl, Policy: &policy.Fake{Approved: policyResult}}
	if _, err := cr.Reconcile(context.TODO(), ctrl.Request{NamespacedName: types.NamespacedName{Name: csrName}}); err != nil {
		panic(err)
	}
	csr := &certv1.CertificateSigningRequest{}
	if err := cl.Get(context.TODO(), types.NamespacedName{Name: csrName}, csr); err != nil {
		panic(err)
	}
	return slices.ContainsFunc(csr.Status.Conditions, func(c certv1.CertificateSigningRequestCondition) bool {
		if c.Type == cType && c.Status == corev1.ConditionStatus(metav1.ConditionTrue) {
			return true
		}
		return false
	})
}

func TestCSRApproved(t *testing.T) {
	if !assertCSRCondition(true, certv1.CertificateApproved) {
		t.Fail()
	}
}

func TestCSRDenied(t *testing.T) {
	if !assertCSRCondition(false, certv1.CertificateDenied) {
		t.Fail()
	}
}
