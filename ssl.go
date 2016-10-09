package main

import (
	"bytes"
)

type SSL string

// address to proxy to
func (t SSL) Address() string {
	return string(t)
}

// identify header as one of TCP
func (t SSL) Identify(header []byte) bool {
	// this is a dummy protocol handler used for the default
	if bytes.Equal(header[:2], []byte{0x16, 0x03}) && header[2] >= 0x00 && header[2] <= 0x03 {
		return true
	}
	return false
}
