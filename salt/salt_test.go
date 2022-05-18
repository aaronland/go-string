package salt

import (
	"testing"
)

func TestSalt(t *testing.T) {

	opts := DefaultSaltOptions()

	s, err := NewRandomSalt(opts)

	if err != nil {
		t.Fatalf("Failed to create string, %v", err)
	}

	if len(s.Bytes()) != opts.Length {
		t.Fatalf("Invalid string length")
	}

	is_valid, err := IsValidSalt(s.String())

	if err != nil {
		t.Fatalf("Failed to validate salt, %v", err)
	}

	if !is_valid {
		t.Fatalf("Invalid salt")
	}
}
