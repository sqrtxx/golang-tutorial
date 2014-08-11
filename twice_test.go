package main

import "testing"

func TestTwice(t *testing.T) {
	const in, out = 2, 4
	if x := Twice(in); x != out {
		t.Errorf("Twice(%v) = %v, want %v", in, x, out)
	}
}
