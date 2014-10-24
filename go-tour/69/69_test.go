package main

import (
	"testing"

	"../tree"
)

func TestSame(t *testing.T) {
	if Same(tree.New(1), tree.New(1)) != true {
		t.Errorf("Same depths trees should be same")
	}
}

func TestDifferent(t *testing.T) {
	if Same(tree.New(1), tree.New(2)) != false {
		t.Errorf("Different depths trees should be not same")
	}
}
