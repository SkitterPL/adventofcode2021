package day1

import (
	"adventofcode/utils"
)

func Day1() (int, int) {
	return Task1(), Task2()
}

func Task1() int {
	data := utils.FileToIntArray("input/1/input.txt")
	return calculateIncreasedNumber(data)
}

func Task2() int {
	var transformedData []int
	var numberOfElementsGrouped = 3

	data := utils.FileToIntArray("input/1/input.txt")
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
