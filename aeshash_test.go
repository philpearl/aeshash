package aeshash_test

import (
	"testing"

	"github.com/philpearl/aeshash"
)

func TestAESHash(t *testing.T) {
	val := aeshash.AESHash("cheese")
	if val != 1315767268 {
		t.Errorf("Expected 1315767268, got %d", val)
	}
}
