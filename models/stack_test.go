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

func TestChangesStack_Push_Length(t *testing.T) {
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

func TestChangesStack_Pop(t *testing.T) {
	expectedInteger := 12
	received := NewStack()
	received.Push(420)
	received.Push(12)

	if expectedInteger != received.Pop() {
		t.Errorf(fmt.Sprintf("expected pop %d, got %d", expectedInteger, received.Pop()))
	}

	type TestStruct struct {
		id   int
		name string
	}

	expectedStruct := TestStruct{
		99,
		"Robin Williams",
	}
	receivedStruct := NewStack()
	receivedStruct.Push(TestStruct{15, "Michael Faraday"})
	receivedStruct.Push(expectedStruct)

	if expectedStruct != receivedStruct.Pop() {
		t.Errorf(fmt.Sprintf("expected pop %v, got %v", expectedStruct, receivedStruct.Pop()))
	}
}
