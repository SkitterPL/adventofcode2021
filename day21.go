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

func (dice *DeterministicDice) roll() int {
	dice.currentValue++
	if dice.currentValue > 100 {
		dice.currentValue = 1
	}
	dice.rollCount++
	return dice.currentValue
}

func day21() (int, int) {
	mostWins := 0
	data := fileToStringArray("input/21/input.txt")
	player1Position, _ := strconv.Atoi(strings.Fields(data[0])[4])
	player2Position, _ := strconv.Atoi(strings.Fields(data[1])[4])
	losingPlayerScore, diceRolls := playWithDeterministicDice(&Player{player1Position, 0}, &Player{player2Position, 0})
	player1Wins, player2Wins := playWithDiracDice(Player{player1Position, 0}, Player{player2Position, 0}, 1, true)
	if player1Wins > player2Wins {
		mostWins = player1Wins
	} else {
		mostWins = player2Wins
	}
	return losingPlayerScore * diceRolls, mostWins
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

//3 1+1+1
//4 1+1+2 1+2+1 2+1+1
//5 1+2+2 1+1+3 1+3+1 3+1+1 2+1+2 2+2+1
//6 2+2+2 1+2+3 3+2+1 2+3+1 2+1+3 1+3+2 3+1+2
//7 1+3+3 3+1+3 3+3+1 2+2+3 2+3+2 3+2+2
//8 3+3+2 2+3+3 3+2+3
//9 3+3+3
func playWithDiracDice(player1 Player, player2 Player, universes int, player1Moves bool) (int, int) {
	possibleDiracDiceRollResults := map[int]int{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}
	if player1.score >= 21 {
		return universes, 0
	} else if player2.score >= 21 {
		return 0, universes
	}

	var currentPlayer Player
	var player1Wins, player2Wins, subPlayer1Wins, subPlayer2Wins int

	for rollSum, possibilities := range possibleDiracDiceRollResults {
		if player1Moves {
			currentPlayer = player1
		} else {
			currentPlayer = player2
		}
		currentPlayer.boardPosition = (currentPlayer.boardPosition + rollSum%10) % 10
		if currentPlayer.boardPosition == 0 {
			currentPlayer.boardPosition = 10
		}
		currentPlayer.score += currentPlayer.boardPosition

		if player1Moves {
			subPlayer1Wins, subPlayer2Wins = playWithDiracDice(currentPlayer, player2, universes*possibilities, !player1Moves)
		} else {
			subPlayer1Wins, subPlayer2Wins = playWithDiracDice(player1, currentPlayer, universes*possibilities, !player1Moves)
		}

		player1Wins += subPlayer1Wins
		player2Wins += subPlayer2Wins
	}

	return player1Wins, player2Wins
}
