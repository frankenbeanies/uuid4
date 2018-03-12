package uuid4

import (
	"testing"
)

// TestUUID4StringIsCorrectLength tests that the length of a generated uuid4 string is 36 characters
func TestUUID4StringIsCorrectLength(t *testing.T) {
	l := len(New().String())

	if l != 36 {
		t.Errorf("String() length was incorrect, expected 36, got %d", l)
	}
}

// TestUUID4StringContainsCorrectHyphens tests that the uuid4 string contains
// the correct hyphens in compliance with RFC 4122
func TestUUID4StringContainsCorrectHyphens(t *testing.T) {
	uuidStr := New().String()

	if uuidStr[8] != '-' {
		t.Errorf("String()[8] was incorrect, expected '-', got %d", uuidStr[8])
	}

	if uuidStr[13] != '-' {
		t.Errorf("String()[13] was incorrect, expected '-' got %d", uuidStr[13])
	}

	if uuidStr[18] != '-' {
		t.Errorf("String()[18] was incorrect, expected '-' got %d", uuidStr[18])
	}

	if uuidStr[23] != '-' {
		t.Errorf("String()[23] was incorrect, expected '-' got %d", uuidStr[23])
	}
}
