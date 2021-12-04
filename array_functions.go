package main

import (
	"strconv"
)

func linesToColumnsArray(lines []string) []string {
	var columns = make([]string, len(lines[0]))
	for i := 0; i < len(lines[0]); i++ {
		for _, line := range lines {
			columns[i] += string(line[i])
		}
	}
	return columns
}

func stringToIntArray(stringNumbers []string) []int {
	var numbers []int
	for _, stringNumber := range stringNumbers {
		number, err := strconv.Atoi(stringNumber)
		check(err)
		numbers = append(numbers, number)
	}
	return numbers
}

func sum(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}
