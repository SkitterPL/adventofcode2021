package main

import (
	"bufio"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileToStringArray(path string) []string {
	file, err := os.Open(path)
	check(err)

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func fileToIntArray(path string) []int {
	file, err := os.Open(path)
	check(err)

	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}
	return numbers
}

func fileToColumnStringArray(path string) []string {
	lines := fileToStringArray(path)
	var columns = make([]string, len(lines[0]))
	for i := 0; i < len(lines[0]); i++ {
		for _, line := range lines {
			columns[i] += string(line[i])
		}
	}
	return columns
}
