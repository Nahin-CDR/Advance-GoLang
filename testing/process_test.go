package main

import (
	"process_test.go/process"
	"testing"
)

func TestCheckEven(t *testing.T) {
	i := 10
	expected := "YES"
	result := process.CheckEven(i)

	if expected != result {
		t.Errorf("Expected : %v, got %v", expected, result)
	}
}
