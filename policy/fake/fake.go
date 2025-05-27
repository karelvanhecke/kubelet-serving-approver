// Copyright 2025 Karel Van Hecke
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"errors"

	"github.com/karelvanhecke/kubelet-serving-approver/policy"
)

type Fake struct {
	Approved   bool
	ThrowError bool
}

func (f *Fake) Name() string {
	return "Fake"
}

func (f *Fake) Approve(req policy.Request) (ok bool, err error) {
	if f.ThrowError {
		return false, errors.New("fake policy error")
	}

	if f.Approved {
		return true, nil
	}

	return false, nil
}
