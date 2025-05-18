package main

import (
	"errors"
	"net"
	"os"
	"regexp"

	"github.com/karelvanhecke/kubelet-serving-approver/controller"
	p "github.com/karelvanhecke/kubelet-serving-approver/policy"
	"github.com/karelvanhecke/kubelet-serving-approver/policy/policies"
	flag "github.com/spf13/pflag"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"sigs.k8s.io/controller-runtime/pkg/metrics/server"
)

const (
	namespacePath    = "/var/run/secrets/kubernetes.io/serviceaccount/namespace"
	leaderElectionID = "kubelet-serving-approver"
)

var (
	policy               *string      = flag.StringP("policy", "p", "Static", "Select a policy for the approver. Policies: Static")
	staticAllowedDomains *[]string    = flag.StringSlice("static-allowed-domains", nil, "Static policy: list of allowed domains")
	staticAllowedSubnets *[]net.IPNet = flag.IPNetSlice("static-allowed-subnets", nil, "Static policy: list of allowed subnets")
	staticMatchHost      *string      = flag.String("static-match-host", "", "Static policy: regex to match hosts against")
	leaderElection       *bool        = flag.BoolP("leader-election", "l", true, "Enable leader election")
)

func main() {
	logger := zap.New()
	log.SetLogger(logger)
	klog.SetLogger(logger)

	flag.Parse()

	ns, err := os.ReadFile(namespacePath)
	if err != nil {
		logger.Error(err, "Could not detect namespace")
		os.Exit(1)
	}

	m, err := manager.New(config.GetConfigOrDie(), manager.Options{
		LeaderElection:          *leaderElection,
		LeaderElectionID:        leaderElectionID,
		LeaderElectionNamespace: string(ns),
		Metrics:                 server.Options{},
		HealthProbeBindAddress:  ":9080",
	})
	if err != nil {
		logger.Error(err, "Failed to setup manager")
		os.Exit(1)
	}

	if err := m.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		logger.Error(err, "Could not setup health check")
		os.Exit(1)
	}
	if err := m.AddReadyzCheck("ready", healthz.Ping); err != nil {
		logger.Error(err, "Could not setup ready check")
		os.Exit(1)
	}

	r := &controller.CSRReconciler{Client: m.GetClient()}

	switch *policy {
	case "Static":
		r.Policy = staticPolicy()
	default:
		logger.Error(errors.New("unknown policy"), "No valid policy has been provided")
		os.Exit(1)
	}

	if err := (r).SetupWithManager(m); err != nil {
		logger.Error(err, "Failed to setup CSR reconciler")
		os.Exit(1)
	}

	if err := m.Start(signals.SetupSignalHandler()); err != nil {
		logger.Error(err, "Failed to start manager")
		os.Exit(1)
	}
}

func staticPolicy() p.Policy {
	s := policies.Static{
		AllowedDomains: *staticAllowedDomains,
		AllowedSubnets: *staticAllowedSubnets,
	}

	if *staticMatchHost != "" {
		s.MatchHost = regexp.MustCompile(*staticMatchHost)
	}

	return s
}
