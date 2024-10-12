package models

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	expected := ChangesStack{}
	received := NewStack()

	if expected.Empty() != received.Empty() {
		t.Errorf(fmt.Sprintf("expected stack status %v, got %v", expected.Empty(), received.Empty()))
	}

	if expected.Length() != received.Length() {
		t.Errorf(fmt.Sprintf("expected length %d, got %d", expected.Length(), received.Length()))
	}
}
