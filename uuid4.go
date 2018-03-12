// Package uuid4 provides functions for generating and parsing uuids
// compliant with RFC 4122
package uuid4

import (
	"crypto/rand"
	"encoding/hex"
)

// UUID4 is a container for an RFC 4122 compliant uuid
type UUID4 struct {
	bytes []byte
}

// New generates a new RFC 4122 compliant uuid
func New() UUID4 {
	bytes := make([]byte, 16)
	rand.Read(bytes)

	bytes[6] = byte(0x40 | (int(bytes[6]) & 0xf))
	bytes[8] = byte(0x80 | (int(bytes[8]) & 0x3f))

	return UUID4{bytes: bytes}
}

// String provides the uuid in a format compliant with RFC 4122
func (uuid UUID4) String() string {
	str := hex.EncodeToString(uuid.bytes)

	return str[:8] + "-" + str[8:12] + "-" + str[12:16] + "-" + str[16:20] + "-" + str[20:]
}

// Bytes provides the bytes of the uuid
func (uuid UUID4) Bytes() []byte {
	val := make([]byte, 16)
	copy(val, uuid.bytes)
	return val
}
