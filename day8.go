package main

import "strings"

//https://adventofcode.com/2021/day/8

var uniqueSegments = map[int]int{
	2: 1,
	4: 4,
	3: 7,
	7: 8,
}

func day8() (int, int) {
	data := fileToStringArray("input/8/input.txt")
	return day8Task1(data), 0

}

func day8Task1(data []string) int {
	result := 0
	for _, line := range data {
		numbers := strings.Fields(strings.Split(line, "|")[1])
		for _, codedNumber := range numbers {
			len := len(codedNumber)
			//Unique segment numbers
			if len == 2 || len == 4 || len == 3 || len == 7 {
				result++
			}
		}
	}
	return result
}
