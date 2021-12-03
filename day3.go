package main

import (
	"strconv"
)

//https://adventofcode.com/2021/day/3

func day3() (int, int) {
	return day3Task1(), 0
}

func day3Task1() int {
	data := fileToColumnStringArray("input/3/input.txt")
	binaryGammaRate, binaryEpsilonRate := calculateGammaAndEpsylon(data)
	gammaRate, err := strconv.ParseInt(binaryGammaRate, 2, 64)
	check(err)
	epsilonRate, err := strconv.ParseInt(binaryEpsilonRate, 2, 64)
	check(err)
	return int(gammaRate) * int(epsilonRate)
}

func calculateGammaAndEpsylon(data []string) (string, string) {
	var gammaRate, epsilonRate string

	for _, binaryString := range data {
		mostCommonCharacter := getMostCommonBitInPhrase(binaryString)
		leastCommonCharacter := "1"
		if mostCommonCharacter == leastCommonCharacter {
			leastCommonCharacter = "0"
		}
		gammaRate += mostCommonCharacter
		epsilonRate += leastCommonCharacter
	}

	return gammaRate, epsilonRate
}

func getMostCommonBitInPhrase(phrase string) string {
	characterNumbers := len(phrase)
	zeroNumber := calculateCharacterNumber('0', phrase)
	oneNumber := characterNumbers - zeroNumber
	mostCommonCharacter := "0"
	if oneNumber > zeroNumber {
		mostCommonCharacter = "1"
	}
	return mostCommonCharacter
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
