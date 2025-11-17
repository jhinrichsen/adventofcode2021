package main

import "testing"

// BenchmarkDetectCPU is a minimal benchmark used solely to trigger Go's CPU detection
func BenchmarkDetectCPU(b *testing.B) {
	for range b.N {
		_ = 1 + 1
	}
}
