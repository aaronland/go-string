package dsn

import (
	"testing"
)

func TestDSN(t *testing.T) {

	str_dsn := "foo=bar hello=world"

	d, err := StringToDSN(str_dsn)

	if err != nil {
		t.Fatalf("Failed to parse '%s', %v", str_dsn, err)
	}

	if d.String() != str_dsn {
		t.Fatalf("Unexpected string value: %s", d.String())
	}

	d, err = StringToDSNWithKeys(str_dsn, "hello", "foo")

	if err != nil {
		t.Fatalf("Failed to parse '%s' with keys, %v", str_dsn, err)
	}

	_, err = StringToDSNWithKeys(str_dsn, "hello", "world")

	if err == nil {
		t.Fatalf("Expected '%s' to fail parsing with keys but did not", str_dsn)
	}

}
