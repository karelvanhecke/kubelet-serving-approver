// Copyright 2025 Karel Van Hecke
// SPDX-License-Identifier: Apache-2.0

package policies_test

import (
	"crypto/x509"
	"net"
	"regexp"
	"testing"

	"github.com/karelvanhecke/kubelet-serving-approver/policy"
	"github.com/karelvanhecke/kubelet-serving-approver/policy/policies"
)

func TestStaticAllowedDomainApproval(t *testing.T) {
	dnsNames := []string{"host.example.org"}
	allowedDomains := []string{"example.org"}

	req := policy.Request{
		CSR:  &x509.CertificateRequest{DNSNames: dnsNames},
		Node: policy.Node{DNSNames: dnsNames},
	}

	s := &policies.Static{AllowedDomains: allowedDomains}
	ok, err := s.Approve(req)
	if !ok || err != nil {
		t.Fail()
	}
}

func TestStaticAllowedDomainDenial(t *testing.T) {
	dnsNames := []string{"host.example.org"}
	allowedDomains := []string{"example.com"}

	req := policy.Request{
		CSR:  &x509.CertificateRequest{DNSNames: dnsNames},
		Node: policy.Node{DNSNames: dnsNames},
	}

	s := &policies.Static{AllowedDomains: allowedDomains}
	ok, err := s.Approve(req)
	if ok || err != nil {
		t.Fail()
	}
}

func TestStaticMatchHostApproval(t *testing.T) {
	dnsNames := []string{"host.example.org", "host"}

	req := policy.Request{
		CSR:  &x509.CertificateRequest{DNSNames: dnsNames},
		Node: policy.Node{DNSNames: dnsNames},
	}

	s := &policies.Static{MatchHost: regexp.MustCompile("^host$")}
	ok, err := s.Approve(req)
	if !ok || err != nil {
		t.Fail()
	}
}

func TestStaticMatchHostDenial(t *testing.T) {
	dnsNames := []string{"host.example.org", "host2"}

	req := policy.Request{
		CSR:  &x509.CertificateRequest{DNSNames: dnsNames},
		Node: policy.Node{DNSNames: dnsNames},
	}

	s := &policies.Static{MatchHost: regexp.MustCompile("^host$")}
	ok, err := s.Approve(req)
	if ok || err != nil {
		t.Fail()
	}
}

func TestMissingReportedDNSName(t *testing.T) {
	req := policy.Request{
		CSR:  &x509.CertificateRequest{DNSNames: []string{"host.example.org"}},
		Node: policy.Node{DNSNames: []string{"host"}},
	}

	s := &policies.Static{}
	ok, err := s.Approve(req)
	if ok || err != nil {
		t.Fail()
	}
}

func TestAllowedIPv4SubnetApproval(t *testing.T) {
	ipAddresses := []net.IP{net.ParseIP("192.0.2.1")}
	_, allowedSubnet, _ := net.ParseCIDR("192.0.2.0/24")

	req := policy.Request{
		CSR:  &x509.CertificateRequest{IPAddresses: ipAddresses},
		Node: policy.Node{IPAddresses: ipAddresses},
	}

	s := &policies.Static{AllowedSubnets: []net.IPNet{*allowedSubnet}}
	ok, err := s.Approve(req)
	if !ok || err != nil {
		t.Fail()
	}
}

func TestAllowedIPv4SubnetDenial(t *testing.T) {
	ipAddresses := []net.IP{net.ParseIP("198.51.100.1")}
	_, allowedSubnet, _ := net.ParseCIDR("192.0.2.0/24")

	req := policy.Request{
		CSR:  &x509.CertificateRequest{IPAddresses: ipAddresses},
		Node: policy.Node{IPAddresses: ipAddresses},
	}

	s := &policies.Static{AllowedSubnets: []net.IPNet{*allowedSubnet}}
	ok, err := s.Approve(req)
	if ok || err != nil {
		t.Fail()
	}
}

func TestAllowedIPv6SubnetApproval(t *testing.T) {
	ipAddresses := []net.IP{net.ParseIP("2001:db8::")}
	_, allowedSubnet, _ := net.ParseCIDR("2001:db8::/32")

	req := policy.Request{
		CSR:  &x509.CertificateRequest{IPAddresses: ipAddresses},
		Node: policy.Node{IPAddresses: ipAddresses},
	}

	s := &policies.Static{AllowedSubnets: []net.IPNet{*allowedSubnet}}
	ok, err := s.Approve(req)
	if !ok || err != nil {
		t.Fail()
	}
}

func TestAllowedIPv6SubnetDenial(t *testing.T) {
	ipAddresses := []net.IP{net.ParseIP("2001:db8::")}
	_, allowedSubnet, _ := net.ParseCIDR("3fff::/20")

	req := policy.Request{
		CSR:  &x509.CertificateRequest{IPAddresses: ipAddresses},
		Node: policy.Node{IPAddresses: ipAddresses},
	}

	s := &policies.Static{AllowedSubnets: []net.IPNet{*allowedSubnet}}
	ok, err := s.Approve(req)
	if ok || err != nil {
		t.Fail()
	}
}

func TestMissingReportedIPAddress(t *testing.T) {
	req := policy.Request{
		CSR:  &x509.CertificateRequest{IPAddresses: []net.IP{net.ParseIP("2001:db8::")}},
		Node: policy.Node{IPAddresses: []net.IP{net.ParseIP("2001:db8::ffff")}},
	}

	s := &policies.Static{}
	ok, err := s.Approve(req)
	if ok || err != nil {
		t.Fail()
	}
}
