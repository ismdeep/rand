package rand

import (
	"crypto/rand"
	"fmt"
)

// Hex generate hex bytes
func Hex(len int) []byte {
	bytes := make([]byte, len)
	_, _ = rand.Read(bytes)
	return bytes
}

// HexStr generate hex string
func HexStr(len int) string {
	return fmt.Sprintf("%x", Hex(len/2+1))[:len]
}
