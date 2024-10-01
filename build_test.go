package main

import (
	"github.com/wailsapp/wails/v2/pkg/options"
	"testing"
)

func TestBuild(t *testing.T) {
	expected := Build{
		nil,
		options.App{},
		true,
	}
	received, err := Build{}.build()

	if err != nil {
		t.Errorf(RError()+" unable initialize Build : %f", err)
	}

	if received.AppInstance == expected.AppInstance {
		t.Errorf(RError()+" expected Build.AppInstance to be initialized with a pointer, but got %f", received.AppInstance)
	}

	if received.TempDirInitialized != expected.TempDirInitialized {
		t.Errorf(RError() + " expected Build.TempDirInitialized to be true, but got false")
	}
}

func TestTempDirHandling(t *testing.T) {
	err := initializeTemporaryDir()

	if err != nil {
		t.Errorf(RError()+" unable to initialize temporary dir : %s", err)
	}

	err = removeTemporaryDir()

	if err != nil {
		t.Errorf(RError()+" unable to remove temporary dir : %s", err)
	}
}
