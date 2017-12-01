package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	a := 1

	if a == 2 {
		t.Fail()
	}
}
