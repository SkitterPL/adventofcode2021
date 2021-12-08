package main

import (
	"strconv"
	"strings"
	"sync"
)

//https://adventofcode.com/2021/day/8

func day8() (int, int) {
	data := fileToStringArray("input/8/input.txt")
	return day8Task1(data), day8Task2(data)

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

func day8Task2(data []string) int {
	result := 0
	recordsToSum := len(data)
	channel := make(chan int, len(data))
	wg := sync.WaitGroup{}
	wg.Add(recordsToSum)
	for _, line := range data {
		go decodeNumbersAsync(line, &wg, channel)
	}
	wg.Wait()
	for i := 0; i < recordsToSum; i++ {
		result += <-channel
	}
	return result
}

func decodeNumbersAsync(line string, wg *sync.WaitGroup, channel chan int) {
	defer wg.Done()
	encodedNumbers := strings.Split(line, "|")
	decodedNumbers := decodeNumbers(strings.Fields(encodedNumbers[0]))
	encodedSumToCalcualte := strings.Fields(encodedNumbers[1])
	sum := ""
	for _, encodedNumber := range encodedSumToCalcualte {
		for key, phrase := range decodedNumbers {
			if len(getStringDiff(phrase, encodedNumber)) == 0 {
				sum += strconv.Itoa(key)
				break
			}
		}
	}
	lineResult, _ := strconv.Atoi(sum)
	channel <- lineResult
}

func decodeNumbers(numbers []string) [10]string {
	decodedNumbers := [10]string{}

	getObviousNumbers(&decodedNumbers, numbers)
	get6And0And9(&decodedNumbers, numbers)
	get5And2And3(&decodedNumbers, numbers)

	return decodedNumbers
}

func getObviousNumbers(decodedNumbers *[10]string, data []string) {
	for _, codedNumber := range data {
		len := len(codedNumber)
		//1
		if len == 2 {
			decodedNumbers[1] = codedNumber
			continue
		}
		//4
		if len == 4 {
			decodedNumbers[4] = codedNumber
			continue
		}
		//7
		if len == 3 {
			decodedNumbers[7] = codedNumber
			continue
		}
		//8
		if len == 7 {
			decodedNumbers[8] = codedNumber
			continue
		}
	}
}

func get6And0And9(decodedNumbers *[10]string, data []string) {
	fourAndOneDiff := getStringDiff(decodedNumbers[1], decodedNumbers[4])
	for _, codedNumber := range data {
		if len(codedNumber) != 6 {
			continue
		}
		//0
		if len(getStringDiff(codedNumber, fourAndOneDiff)) != 4 {
			decodedNumbers[0] = codedNumber
		} else {
			//9
			if len(getStringDiff(codedNumber, decodedNumbers[4])) == 2 {
				decodedNumbers[9] = codedNumber
			} else { //6
				decodedNumbers[6] = codedNumber
			}

		}

	}
}

func get5And2And3(decodedNumbers *[10]string, data []string) {
	eightAndSixDiff := getStringDiff(decodedNumbers[8], decodedNumbers[6])
	for _, codedNumber := range data {
		codedNumberLength := len(codedNumber)
		if codedNumberLength != 5 {
			continue
		}
		//5
		if len(getStringDiff(codedNumber, eightAndSixDiff)) == 5 {
			decodedNumbers[5] = codedNumber
			break
		}
	}
	fiveAndEightDiff := getStringDiff(decodedNumbers[5], decodedNumbers[8])
	for _, codedNumber := range data {
		if len(codedNumber) != 5 {
			continue
		}
		diff := len(getStringDiff(codedNumber, fiveAndEightDiff))
		//5
		if diff == 4 {
			decodedNumbers[3] = codedNumber
		} else if diff == 3 {
			decodedNumbers[2] = codedNumber
		}
	}
}

func getStringDiff(a string, b string) string {
	var longer, shorter string
	if len(a) < len(b) {
		longer = b
		shorter = a
	} else {
		longer = a
		shorter = b
	}
	length := len(longer)

	diff := ""
	for i := 0; i < length; i++ {
		hasLetter := false
		for j := 0; j < len(shorter); j++ {
			if longer[i] == shorter[j] {
				hasLetter = true
				break
			}
		}
		if !hasLetter {
			diff += string(longer[i])
		}
	}
	return diff
}
