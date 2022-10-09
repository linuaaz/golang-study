package ch2

import (
	"runtime"
	"testing"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("skip MacOs")
	}

	if testing.Short() {
		t.Skip("skip because of short")
	}

	t.Log("there a non-skipped log")
}
