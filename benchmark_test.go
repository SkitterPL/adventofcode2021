package main

import "testing"

//go test -count 5 -run '^$' -bench . -memprofile=v1.mem.pprof -cpuprofile=v1.cpu.pprof > v1.txt
func BenchmarkAdventOfCode(b *testing.B) {
	data := fileToStringArray("input/10/input.txt")
	for n := 0; n < b.N; n++ {
		day10Task2(data)
	}
}
