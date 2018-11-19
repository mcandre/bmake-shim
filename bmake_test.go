package bmake_test

import (
	"testing"

	"github.com/mcandre/bmake-shim"
)

func TestVersion(t *testing.T) {
	if bmake.Version == "" {
		t.Errorf("Expected %v to be non-empty", bmake.Version)
	}
}
