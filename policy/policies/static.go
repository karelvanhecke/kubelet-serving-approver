package policies

import (
	"net"
	"regexp"
	"slices"

	"github.com/karelvanhecke/kubelet-serving-approver/policy"
	"github.com/miekg/dns"
)

type Static struct {
	RequireFQDN    bool           `yaml:"requireFQDN"`
	AllowedDomains []string       `yaml:"allowedDomains"`
	AllowedSubnets []net.IPNet    `yaml:"allowedSubnets"`
	MatchHost      *regexp.Regexp `yaml:"matchHost"`
}

func (s Static) Name() string {
	return "Static"
}

func (s Static) Approve(req policy.Request) (ok bool, err error) {
	dnsNames := req.CSR.DNSNames
	ipAddresses := req.CSR.IPAddresses

	reportedDNSNames := req.Node.DNSNames
	reportedIPAddresses := req.Node.IPAddresses

	if ok := s.CheckDNSNames(dnsNames, reportedDNSNames); !ok {
		return false, nil
	}

	if ok := s.CheckIPAddresses(ipAddresses, reportedIPAddresses); !ok {
		return false, nil
	}

	return true, nil
}

func (s Static) CheckDNSNames(requested []string, reported []string) (ok bool) {
	for _, dn := range requested {
		if s.RequireFQDN {
			if !dns.IsFqdn(dn) {
				return false
			}

			allowed := s.AllowedDomains == nil
			for _, ad := range s.AllowedDomains {
				allowed = dns.IsSubDomain(ad, dn)
				if allowed {
					break
				}
			}
			if !allowed {
				return false
			}
		}

		if s.MatchHost != nil {
			if !s.MatchHost.Match([]byte(dns.SplitDomainName(dn)[0])) {
				return false
			}
		}

		if !slices.Contains(reported, dn) {
			return false
		}
	}

	return true
}

func (s Static) CheckIPAddresses(requested []net.IP, reported []net.IP) (ok bool) {
	for _, ip := range requested {
		allowed := s.AllowedSubnets == nil
		for _, sn := range s.AllowedSubnets {
			allowed = sn.Contains(ip)
			if allowed {
				break
			}
		}
		if !allowed {
			return false
		}

		if !slices.ContainsFunc(reported, func(ri net.IP) bool {
			return ri.String() == ip.String()
		}) {
			return false
		}
	}
	return true
}
