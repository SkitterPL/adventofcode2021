package main

import "testing"

//go test -count 5 -run '^$' -bench . -memprofile=v1.mem.pprof -cpuprofile=v1.cpu.pprof > v1.txt
func BenchmarkAdventOfCode(b *testing.B) {
	data := fileTo2DIntArray("input/11/input.txt")
	for n := 0; n < b.N; n++ {
		getThroughCavern(data)
	}
}
