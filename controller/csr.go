// Copyright 2025 Karel Van Hecke
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"net"
	"strings"

	"github.com/karelvanhecke/kubelet-serving-approver/policy"
	certv1 "k8s.io/api/certificates/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

type CSRReconciler struct {
	client.Client
	Policy policy.Policy
}

func (cr *CSRReconciler) Reconcile(ctx context.Context, req ctrl.Request) (res ctrl.Result, err error) {
	csr := &certv1.CertificateSigningRequest{}

	if err := cr.Get(ctx, req.NamespacedName, csr); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	csrDER, _ := pem.Decode(csr.Spec.Request)

	csrData, err := x509.ParseCertificateRequest(csrDER.Bytes)
	if err != nil {
		return ctrl.Result{}, err
	}

	node := &corev1.Node{}

	if err := cr.Get(ctx, types.NamespacedName{Name: strings.Split(csr.Spec.Username, ":")[2]}, node); err != nil {
		return ctrl.Result{}, err
	}

	policyName := cr.Policy.Name()

	policyRequest := policy.Request{
		CSR:  csrData,
		Node: policy.Node{},
	}

	for _, a := range node.Status.Addresses {
		if a.Type == corev1.NodeHostName ||
			a.Type == corev1.NodeExternalDNS ||
			a.Type == corev1.NodeInternalDNS {
			policyRequest.Node.DNSNames = append(policyRequest.Node.DNSNames, a.Address)
		}
		if a.Type == corev1.NodeExternalIP ||
			a.Type == corev1.NodeInternalIP {
			ip := net.ParseIP(a.Address)
			if ip != nil {
				policyRequest.Node.IPAddresses = append(policyRequest.Node.IPAddresses, ip)
			}
		}
	}

	ok, err := cr.Policy.Approve(policyRequest)
	if err != nil {
		return ctrl.Result{}, err
	}

	if ok {
		csr.Status.Conditions = []certv1.CertificateSigningRequestCondition{
			{
				Type:    certv1.CertificateApproved,
				Status:  corev1.ConditionTrue,
				Reason:  "ApprovedBy" + policyName + "Policy",
				Message: "Request approved by policy: " + policyName,
			},
		}
	} else {
		csr.Status.Conditions = []certv1.CertificateSigningRequestCondition{
			{
				Type:    certv1.CertificateDenied,
				Status:  corev1.ConditionTrue,
				Reason:  "DeniedBy" + policyName + "Policy",
				Message: "Request denied by " + policyName + " policy",
			},
		}
	}

	return ctrl.Result{}, cr.SubResource("approval").Update(ctx, csr)
}

func (cr *CSRReconciler) SetupWithManager(m ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(m).For(&certv1.CertificateSigningRequest{},
		builder.WithPredicates(predicate.Funcs{
			CreateFunc: func(e event.CreateEvent) bool {
				return requiresApproval(e.Object)
			},
			UpdateFunc: func(e event.UpdateEvent) bool {
				return requiresApproval(e.ObjectNew)
			},
			DeleteFunc: func(event.DeleteEvent) bool {
				return false
			},
			GenericFunc: func(event.GenericEvent) bool {
				return false
			},
		})).Complete(cr)
}

func requiresApproval(o client.Object) bool {
	if csr, ok := o.(*certv1.CertificateSigningRequest); !ok ||
		csr.Spec.SignerName != "kubernetes.io/kubelet-serving" ||
		csr.Status.Conditions != nil {
		return false
	}

	return true
}
