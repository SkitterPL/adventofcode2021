package main

import (
	"testing"
)

//go test -count 5 -run '^$' -bench . -memprofile=v1.mem.pprof -cpuprofile=v1.cpu.pprof > v1.txt
func BenchmarkAdventOfCode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		day20()
	}
}
