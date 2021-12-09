package main

import (
	"strconv"
	"strings"
	"sync"
)

//https://adventofcode.com/2021/day/9

func day9() (int, int) {
	data := fileToStringArray("input/9/input.txt")
	return day9Task1(data), day9Task2(data)
}

func day9Task1(data []string) int {
	linesNumber := len(data)
	result := 0
	ch := make(chan int, linesNumber)
	wg := sync.WaitGroup{}
	wg.Add(linesNumber)
	go calculateLowestPointsRisk(data[0], data[1], "", ch, &wg)
	for i := 1; i < linesNumber-1; i++ {
		go calculateLowestPointsRisk(data[i], data[i+1], data[i-1], ch, &wg)
	}
	go calculateLowestPointsRisk(data[linesNumber-1], "", data[linesNumber-2], ch, &wg)
	wg.Wait()
	for i := 0; i < linesNumber; i++ {
		result += <-ch
	}
	return result
}

func calculateLowestPointsRisk(
	currentLine string,
	nextLine string,
	previousLine string,
	result chan int,
	waitGroup *sync.WaitGroup) {

	defer waitGroup.Done()
	higherByte := ":"[0]
	lineLen := len(currentLine)
	sum := 0
	if previousLine == "" {
		previousLine = strings.Repeat(":", lineLen)
	} else if nextLine == "" {
		nextLine = strings.Repeat(":", lineLen)
	}
	for key := range currentLine {
		if key == 0 {
			sum += isLowest(currentLine[key], higherByte, currentLine[key+1], previousLine[key], nextLine[key])
		} else if key == lineLen-1 {
			sum += isLowest(currentLine[key], currentLine[key-1], higherByte, previousLine[key], nextLine[key])
		} else {
			sum += isLowest(currentLine[key], currentLine[key-1], currentLine[key+1], previousLine[key], nextLine[key])
		}
	}
	result <- sum
}

func isLowest(
	current byte,
	left byte,
	right byte,
	up byte,
	down byte) int {
	if current < left &&
		current < right &&
		current < up &&
		current < down {
		value, _ := strconv.Atoi(string(current))
		return 1 + value
	}
	return 0
}

func day9Task2(data []string) int {
	result := 0
	return result
}
