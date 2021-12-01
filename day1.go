package main

//https://adventofcode.com/2021/day/1

func day1() (int, int) {
	return day1Task1(), day1Task2()
}

func day1Task1() int {
	data := fileToIntArray("input/1/input.txt")
	return calculateIncreasedNumber(data)
}

func day1Task2() int {
	var transformedData []int
	var numberOfElementsGrouped = 3

	data := fileToIntArray("input/1/input.txt")
	for key := range data {
		groupedSlice := data[key : key+numberOfElementsGrouped]
		sumOfElements := groupedSlice[0] + groupedSlice[1] + groupedSlice[2]
		transformedData = append(transformedData, sumOfElements)
	}
	return calculateIncreasedNumber(transformedData)
}

func calculateIncreasedNumber(data []int) int {
	var previousData, increasedNumber int = data[0], 0
	for _, currentData := range data[1:] {
		if currentData > previousData {
			increasedNumber++
		}
		previousData = currentData
	}
	return increasedNumber
}
