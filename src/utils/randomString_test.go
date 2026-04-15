package utils

import (
	"encoding/hex"
	"testing"
)

func TestRandomString_Length(t *testing.T) {
	for _, length := range []int{1, 8, 16, 32} {
		result, err := RandomString(length)
		if err != nil {
			t.Fatalf("RandomString(%d) returned error: %v", length, err)
		}
		// hex encoding doubles the byte length
		expected := length * 2
		if len(result) != expected {
			t.Errorf("RandomString(%d): got len %d, want %d", length, len(result), expected)
		}
	}
}

func TestRandomString_IsHex(t *testing.T) {
	result, err := RandomString(16)
	if err != nil {
		t.Fatalf("RandomString returned error: %v", err)
	}
	if _, err := hex.DecodeString(result); err != nil {
		t.Errorf("RandomString result is not valid hex: %q", result)
	}
}

func TestRandomString_Unique(t *testing.T) {
	a, _ := RandomString(16)
	b, _ := RandomString(16)
	if a == b {
		t.Error("RandomString returned identical values on consecutive calls")
	}
}
