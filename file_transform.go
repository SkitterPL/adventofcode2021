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

func fileTo2DIntArray(path string) [][]int {
	file, err := os.Open(path)
	check(err)

	defer file.Close()

	var numbers [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		len := len(line)
		lineNumbers := make([]int, len)
		for key, number := range line {
			lineNumbers[key], _ = strconv.Atoi(string(number))
		}
		numbers = append(numbers, lineNumbers)
	}
	return numbers
}

func fileToColumnStringArray(path string) []string {
	lines := fileToStringArray(path)
	return linesToColumnsArray(lines)
}
