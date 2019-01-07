// Copyright (C) 2018 Sam Olds

package testhelp

import "testing"

type T struct {
	*testing.T
}

func New(t *testing.T) *T {
	return &T{T: t}
}

func (t *T) AssertEqual(a interface{}, b interface{}) {
	t.assert(a == b, "%v == %v", a, b)
}

func (t *T) AssertNotEqual(a interface{}, b interface{}) {
	t.assert(a != b, "%v != %v", a, b)
}

func (t *T) AssertError(err error) {
	t.assert(err != nil, "%v != nil", err)
}

func (t *T) AssertNoError(err error) {
	t.assert(err == nil, "%v == nil", err)
}

func (t *T) AssertNotNil(a interface{}) {
	t.assert(a != nil, "%v != nil", a)
}

func (t *T) Assert(a bool) {
	t.assert(a, "%b == true", a)
}

func (t *T) assert(b bool, fexp string, args ...interface{}) {
	if b {
		return
	}

	t.T.Fatalf("expected: '"+fexp+"'", args...)
}
