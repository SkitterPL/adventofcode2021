package main

import (
	"math"
	"strings"
	"sync"
)

//https://adventofcode.com/2021/day/7

func day7() (int, int) {
	data := strings.Split(fileToStringArray("input/7/input.txt")[0], ",")
	floatData := stringToFloatArray(data)
	return calculateFuelUsage(floatData, false), calculateFuelUsage(floatData, true)
}

func calculateFuelUsage(data []float64, sumAsSequenceSum bool) int {
	wg := sync.WaitGroup{}
	crabLen := len(data)
	channel := make(chan float64, crabLen)
	wg.Add(crabLen)
	for i := 0; i < crabLen; i++ {
		if sumAsSequenceSum {
			go calculateAdvancedCrabsFuelUsage(float64(i), data, &wg, channel)
		} else {
			go calculateCrabsFuelUsage(float64(i), data, &wg, channel)
		}
	}
	wg.Wait()
	smallestFuelUsage := 1000000000.0
	var value float64
	for i := 0; i < crabLen; i++ {
		value = <-channel
		if value < smallestFuelUsage {
			smallestFuelUsage = value
		}
	}
	return int(smallestFuelUsage)
}

func calculateAdvancedCrabsFuelUsage(value float64, data []float64, waitGroup *sync.WaitGroup, channel chan float64) {
	defer waitGroup.Done()
	var sum float64
	for _, crabPosition := range data {
		diff := math.Abs(crabPosition - value)
		sum += diff * (diff + 1.0) / 2.0
	}
	channel <- sum
}

func calculateCrabsFuelUsage(value float64, data []float64, waitGroup *sync.WaitGroup, channel chan float64) {
	defer waitGroup.Done()
	var sum float64
	for _, crabPosition := range data {
		sum += math.Abs(crabPosition - value)
	}
	channel <- sum
}
