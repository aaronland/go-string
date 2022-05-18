package random

import (
	"regexp"
	"testing"
)

func TestRandom(t *testing.T) {

	opts := DefaultOptions()

	s, err := String(opts)

	if err != nil {
		t.Fatalf("Failed to create string, %v", err)
	}

	if len(s) != opts.Length {
		t.Fatalf("Invalid string length")
	}
}

func TestAlphaNumeric(t *testing.T) {

	opts := DefaultOptions()
	opts.Length = 10
	opts.AlphaNumeric = true

	s, err := String(opts)

	if err != nil {
		t.Fatalf("Failed to create string, %v", err)
	}

	if len(s) != opts.Length {
		t.Fatalf("Invalid string length")
	}

	re, err := regexp.Compile(`^[a-zA-Z0-9]+$`)

	if err != nil {
		t.Fatalf("Failed to compile alphnumeric regexp, %v", err)
	}

	if !re.MatchString(s) {
		t.Fatalf("'%s' is not alphanumeric", s)
	}

}
