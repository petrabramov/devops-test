package main

import (
	"bytes"
	"testing"
)

func TestHelloMessage(t *testing.T) {
	var result bytes.Buffer
	HelloMessage(&result)

	if result.String() != "Hello, World" {
		t.Errorf("HelloMessage FAILED, expected Hello, World but got %v", result)
	}
}
