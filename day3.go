package main

import (
	"strconv"
)

//https://adventofcode.com/2021/day/3

func day3() (int, int) {
	return day3Task1(), day3Task2()
}

type LifeSupportRating string

const (
	Oxygen LifeSupportRating = "0"
	CO2    LifeSupportRating = "1"
)

func day3Task1() int {
	data := fileToColumnStringArray("input/3/input.txt")
	gammaRate, epsilonRate := calculateGammaAndEpsylon(data)
	return gammaRate * epsilonRate
}

func day3Task2() int {
	data := fileToStringArray("input/3/input.txt")
	oxygenGeneratorRating := calculateLifeSupportRating(data, Oxygen)
	CO2Rating := calculateLifeSupportRating(data, CO2)

	return int(CO2Rating) * int(oxygenGeneratorRating)
}

func calculateLifeSupportRating(rowData []string, rating LifeSupportRating) int {
	iterations := len(rowData[0])
	for i := 0; i < iterations; i++ {
		columnData := linesToColumnsArray(rowData)
		rowData = removeRowsByBit(rowData, i, getBitToRemove(columnData[i], rating))
		if len(rowData) == 1 {
			result, err := strconv.ParseInt(rowData[0], 2, 64)
			check(err)
			return int(result)
		}
	}
	panic("Something went wrong")
}

func removeRowsByBit(rowData []string, bitPosition int, bitToRemove string) []string {
	var rowDataToReturn []string
	for _, binaryString := range rowData {
		if string(binaryString[bitPosition]) != bitToRemove {
			rowDataToReturn = append(rowDataToReturn, binaryString)
		}
	}
	return rowDataToReturn
}

func calculateGammaAndEpsylon(data []string) (int, int) {
	var binaryGammaRate, binaryEpsilonRate string

	for _, binaryString := range data {
		mostCommonBit, leastCommonBit := getMostAndLeastCommonBitInPhrase(binaryString)
		binaryGammaRate += mostCommonBit
		binaryEpsilonRate += leastCommonBit
	}

	gammaRate, err := strconv.ParseInt(binaryGammaRate, 2, 64)
	check(err)
	epsilonRate, err := strconv.ParseInt(binaryEpsilonRate, 2, 64)
	check(err)

	return int(gammaRate), int(epsilonRate)
}

func getBitToRemove(phrase string, ratingType LifeSupportRating) string {
	characterNumbers := len(phrase)
	zeroNumber := calculateCharacterNumber('0', phrase)
	oneNumber := characterNumbers - zeroNumber
	if zeroNumber == oneNumber {
		return string(ratingType)
	}
	mostCommonBit, leastCommonBit := getMostAndLeastCommonBitInPhrase(phrase)
	if ratingType == Oxygen {
		return leastCommonBit
	}
	return mostCommonBit
}

func getMostAndLeastCommonBitInPhrase(phrase string) (string, string) {
	characterNumbers := len(phrase)
	zeroNumber := calculateCharacterNumber('0', phrase)
	oneNumber := characterNumbers - zeroNumber
	var mostCommonCharacter, leastCommonCharacter string
	if oneNumber > zeroNumber {
		mostCommonCharacter = "1"
		leastCommonCharacter = "0"
	} else {
		mostCommonCharacter = "0"
		leastCommonCharacter = "1"
	}
	return mostCommonCharacter, leastCommonCharacter
}

func calculateCharacterNumber(character byte, phrase string) int {
	result := 0
	for i := 0; i < len(phrase); i++ {
		if character == phrase[i] {
			result++
		}
	}
	return result
}
