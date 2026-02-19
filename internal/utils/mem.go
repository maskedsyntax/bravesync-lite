package utils

import "runtime"

// ZeroMem zeroes the byte slice.
func ZeroMem(b []byte) {
	for i := range b {
		b[i] = 0
	}
	runtime.KeepAlive(b)
}
