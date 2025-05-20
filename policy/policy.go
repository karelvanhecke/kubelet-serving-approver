// Copyright 2025 Karel Van Hecke
// SPDX-License-Identifier: Apache-2.0

package policy

type Policy interface {
	Name() string
	Approve(Request) (ok bool, err error)
}
