package main

import (
	"strconv"
	"strings"
)

//https://adventofcode.com/2021/day/21

type Player struct {
	boardPosition int
	score         int
}

type Dice interface {
	roll() int
}

type DeterministicDice struct {
	currentValue int
	rollCount    int
}

func day21() (int, int) {
	data := fileToStringArray("input/21/test_input.txt")
	player1Position, _ := strconv.Atoi(strings.Fields(data[0])[4])
	player2Position, _ := strconv.Atoi(strings.Fields(data[1])[4])
	losingPlayerScore, diceRolls := playWithDeterministicDice(&Player{player1Position, 0}, &Player{player2Position, 0})
	return losingPlayerScore * diceRolls, 0
}

func playWithDeterministicDice(player1 *Player, player2 *Player) (int, int) {
	currentPlayer := player1
	var roll, losingPoints int
	dice := DeterministicDice{0, 0}
	for {
		roll = dice.roll() + dice.roll() + dice.roll()
		currentPlayer.boardPosition = (currentPlayer.boardPosition + roll%10) % 10
		if currentPlayer.boardPosition == 0 {
			currentPlayer.boardPosition = 10
		}
		currentPlayer.score += currentPlayer.boardPosition
		if currentPlayer.score >= 1000 {
			break
		}
		if currentPlayer == player1 {
			currentPlayer = player2
		} else {
			currentPlayer = player1
		}

	}
	if player1.score >= 1000 {
		losingPoints = player2.score
	} else {
		losingPoints = player1.score
	}
	return losingPoints, dice.rollCount
}

func (dice *DeterministicDice) roll() int {
	dice.currentValue++
	if dice.currentValue > 100 {
		dice.currentValue = 1
	}
	dice.rollCount++
	return dice.currentValue
}
