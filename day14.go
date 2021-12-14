package main

import (
	"strings"
)

//https://adventofcode.com/2021/day/14

func day14() (int, int) {
	data := fileToStringArray("input/14/input.txt")
	template := data[0]
	instructions := make(map[string]string, len(data[2:]))
	for _, line := range data[2:] {
		instructionParsed := strings.Split(line, " -> ")
		instructions[instructionParsed[0]] = instructionParsed[1]
	}
	return calculateOccurences(template, instructions, 10), calculateOccurences(template, instructions, 40)
}

func calculateOccurences(template string, instructions map[string]string, stepNumber int) int {
	occurences := map[string]int{}
	pairsOccurences := map[string]int{}
	templateLen := len(template)

	for i := 0; i < templateLen; i++ {
		occurences[string(template[i])]++
		if i < templateLen-1 {
			pairsOccurences[template[i:i+2]]++
		}
	}

	for i := 0; i < stepNumber; i++ {
		nextPairOccurences := make(map[string]int)
		for pair, numberOfOccurences := range pairsOccurences {
			newElement := instructions[pair]
			occurences[newElement] += numberOfOccurences
			nextPairOccurences[string(pair[0])+string(newElement)] += numberOfOccurences
			nextPairOccurences[string(newElement)+string(pair[1])] += numberOfOccurences
		}
		pairsOccurences = nextPairOccurences
	}

	min := 10000000000000
	max := -1
	for _, value := range occurences {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	return max - min
}
