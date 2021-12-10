package main

import "sort"

//https://adventofcode.com/2021/day/10

type Chunk struct {
	opening rune
	closing rune
}

func (chunk *Chunk) close(char rune) {
	chunk.closing = char
}

func (chunk *Chunk) isCorrupted() bool {
	return !((chunk.opening == '(' && chunk.closing == ')') ||
		(chunk.opening == '[' && chunk.closing == ']') ||
		(chunk.opening == '{' && chunk.closing == '}') ||
		(chunk.opening == '<' && chunk.closing == '>'))
}

const NIL_CLOSING_RUNE = ':'

func day10() (int, int) {
	data := fileToStringArray("input/10/input.txt")
	return day10Task1(data), day10Task2(data)
}

func day10Task1(data []string) int {
	result := 0
	for _, line := range data {
		result += getScoreForPart1(calculateFirstIllegalCharacter(line))
	}
	return result
}

func day10Task2(data []string) int {
	values := 0
	var valueStore []int
	for _, line := range data {
		value := getScoreForPart2(calculateCompletionString(line))
		if value != 0 {
			valueStore = append(valueStore, value)
			values++
		}
	}
	sort.Ints(valueStore)
	return valueStore[values/2]
}

func calculateFirstIllegalCharacter(line string) rune {
	chunksStorage := make([]*Chunk, len(line))
	openingBrackets := 0
	for _, character := range line {
		if isOpeningBracket(character) {
			chunksStorage[openingBrackets] = &Chunk{character, NIL_CLOSING_RUNE}
			openingBrackets++
			continue
		} else {
			chunksStorage[openingBrackets-1].close(character)
			if chunksStorage[openingBrackets-1].isCorrupted() {
				return chunksStorage[openingBrackets-1].closing
			}
			openingBrackets--
		}
	}

	return NIL_CLOSING_RUNE
}

func calculateCompletionString(line string) string {
	chunksStorage := make([]*Chunk, len(line))
	openingBrackets := 0
	for _, character := range line {
		if isOpeningBracket(character) {
			chunksStorage[openingBrackets] = &Chunk{character, NIL_CLOSING_RUNE}
			openingBrackets++
			continue
		} else {
			chunksStorage[openingBrackets-1].close(character)
			if chunksStorage[openingBrackets-1].isCorrupted() {
				return ""
			}
			openingBrackets--
		}
	}
	result := ""
	for _, chunk := range chunksStorage {
		if chunk == nil {
			return reverse(result)
		}
		if chunk.closing == NIL_CLOSING_RUNE {
			result += string(completeChunk(chunk.opening))
		}
	}

	return reverse(result)
}

func isOpeningBracket(char rune) bool {
	return char == '{' || char == '<' || char == '(' || char == '['
}

func getScoreForPart1(char rune) int {
	switch char {
	case '}':
		return 1197
	case ')':
		return 3
	case ']':
		return 57
	case '>':
		return 25137
	case NIL_CLOSING_RUNE:
		return 0
	}
	return 0
}

func getScoreForPart2(brackets string) int {
	sum := 0
	for _, bracket := range brackets {
		sum = sum * 5
		switch bracket {
		case '}':
			sum += 3
		case ')':
			sum += 1
		case ']':
			sum += 2
		case '>':
			sum += 4
		default:
			sum += 0
		}
	}
	return sum
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func completeChunk(openingBracket rune) rune {
	switch openingBracket {
	case '{':
		return '}'
	case '(':
		return ')'
	case '[':
		return ']'
	case '<':
		return '>'
	}
	return NIL_CLOSING_RUNE
}
