package main

import (
	"strings"
	"sync"
)

//https://adventofcode.com/2021/day/4

type BingoBoard struct {
	notMarkedNumbers     []int
	possibleCombinations []*Combination
	winner               bool
	lastDraw             int
}

type Combination struct {
	numbers []int
}

func day4() (int, int) {
	return day4Task1(), day4Task2()
}

func day4Task1() int {
	data := fileToStringArray("input/4/input.txt")
	drawedNumbers := stringToIntArray(strings.Split(data[0], ","))
	channel := playBingo(data)

	smallestDrawNumber := len(drawedNumbers) + 1
	var fastestWinningBoard BingoBoard
	boardsNumber := len(channel)
	for i := 0; i < boardsNumber; i++ {
		board := <-channel
		if board.winner && board.lastDraw < smallestDrawNumber {
			fastestWinningBoard = board
			smallestDrawNumber = board.lastDraw
		}
	}

	return sum(fastestWinningBoard.notMarkedNumbers...) * drawedNumbers[fastestWinningBoard.lastDraw-1]
}

func day4Task2() int {
	data := fileToStringArray("input/4/input.txt")
	drawedNumbers := stringToIntArray(strings.Split(data[0], ","))
	channel := playBingo(data)

	largestDrawNumber := 0
	var slowestWinningBoard BingoBoard
	boardsNumber := len(channel)
	for i := 0; i < boardsNumber; i++ {
		board := <-channel
		if board.winner && board.lastDraw > largestDrawNumber {
			slowestWinningBoard = board
			largestDrawNumber = board.lastDraw
		}
	}

	return sum(slowestWinningBoard.notMarkedNumbers...) * drawedNumbers[slowestWinningBoard.lastDraw-1]
}

func playBingo(data []string) chan BingoBoard {
	drawedNumbers := stringToIntArray(strings.Split(data[0], ","))
	bingoBoards := transformToBingoBoards(data[2:])
	channel := make(chan BingoBoard, len(bingoBoards))
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(len(bingoBoards))
	for _, board := range bingoBoards {
		go simulateGameAsync(board, drawedNumbers, &wg, &mu, channel)
	}
	wg.Wait()
	return channel
}

func simulateGameAsync(board BingoBoard, drawedNumbers []int, wg *sync.WaitGroup, mu *sync.Mutex, channel chan BingoBoard) {
	board.simulateGame(drawedNumbers)
	mu.Lock()
	defer mu.Unlock()
	channel <- board
	wg.Done()
}

// BingoBoard functions

func (board *BingoBoard) simulateGame(drawedNumbers []int) {
	var alreadyDrawedNumbers []int
	for _, drawedNumber := range drawedNumbers {
		alreadyDrawedNumbers = append(alreadyDrawedNumbers, drawedNumber)
		board.lastDraw++
		board.mark(drawedNumber)
		for _, combination := range board.possibleCombinations {
			if combination.checkIsWinner(alreadyDrawedNumbers) {
				board.winner = true
				return
			}
		}
	}
}

func (board *BingoBoard) mark(numberToMark int) {
	for key, number := range board.notMarkedNumbers {
		if numberToMark == number {
			board.notMarkedNumbers[key] = board.notMarkedNumbers[len(board.notMarkedNumbers)-1]
			board.notMarkedNumbers = board.notMarkedNumbers[:len(board.notMarkedNumbers)-1]
		}
	}
}

//Combination functions

func (combination *Combination) checkIsWinner(drawedNumbers []int) bool {
	neededChecks := len(combination.numbers)
	checks := 0
	for _, number := range drawedNumbers {
		for _, combinationNumber := range combination.numbers {
			if number == combinationNumber {
				checks++
				if checks == neededChecks {
					return true
				}
			}
		}
	}
	return false
}

//Transform data to BingoBoards

func transformToBingoBoards(data []string) []BingoBoard {
	var numbers [][]int
	var boards []BingoBoard
	for _, line := range data {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			boards = append(boards, createBingoBoard(numbers))
			numbers = nil
		} else {
			numbers = append(numbers, stringToIntArray(fields))
		}
	}
	boards = append(boards, createBingoBoard(numbers))
	return boards
}

func createBingoBoard(rowArray [][]int) BingoBoard {
	var columRawCombinations = make([][]int, len(rowArray))
	for i := 0; i < len(rowArray); i++ {
		columRawCombinations[i] = make([]int, len(rowArray[0]))
	}

	var combinations []*Combination
	var allNumbers []int
	for rowKey, row := range rowArray {
		combinations = append(combinations, &Combination{row})
		for colKey, number := range row {
			columRawCombinations[colKey][rowKey] = number
		}
		allNumbers = append(allNumbers, row...)
	}
	for _, column := range columRawCombinations {
		combinations = append(combinations, &Combination{column})
	}
	return BingoBoard{allNumbers, combinations, false, 0}
}
