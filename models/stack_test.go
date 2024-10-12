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

func TestChangesStack_Length(t *testing.T) {
	expected := 0
	received := NewStack()

	if expected != received.Length() {
		t.Errorf(fmt.Sprintf("expected length %d, got %d", expected, received.Length()))
	}

	expected = 2
	received.Push(1000)
	received.Push(80)

	if expected != received.Length() {
		t.Errorf(fmt.Sprintf("expected length %d, got %d", expected, received.Length()))
	}
}
