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

// TestUUID4BytesReturnsCorrectBytes tests that the bytes return by Bytes are correct
func TestUUID4BytesReturnsCorrectBytes(t *testing.T) {
	uuid := New()
	bytes := uuid.Bytes()

	if len(bytes) != 16 {
		t.Errorf("Bytes() was the incorrect length, expected 16, got %d", len(bytes))
	}

	for i, b := range bytes {
		if b != uuid.bytes[i] {
			t.Errorf("Bytes()[%d] != bytes[%d]", i, i)
		}
	}
}

//TestUUID4BytesNotInternalSlice tests that Bytes() does not return the same slice used internally
func TestUUID4BytesNotInternalSlice(t *testing.T) {
	uuid := New()
	bytes := uuid.Bytes()

	bytes[3] += 0x1

	if bytes[3] == uuid.bytes[3] {
		t.Errorf("Bytes() and bytes are referencing the same slice")
	}
}
