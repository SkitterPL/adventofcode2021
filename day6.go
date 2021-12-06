package main

import (
	"strconv"
	"strings"
)

const (
	OLD_REPRODUCE_CYCLE   int = 6
	FIRST_REPRODUCE_CYCLE int = 8
)

const DAYS_NUMBER = 80

type Lanternfish struct {
	daysToReproduce int
	old             bool
}

func (fish *Lanternfish) liveOneDay() {
	fish.daysToReproduce--
	if fish.daysToReproduce == -1 {
		fish.daysToReproduce = OLD_REPRODUCE_CYCLE
		fish.old = true
	}
}

func (fish *Lanternfish) canReproduce() bool {
	return fish.old && fish.daysToReproduce == OLD_REPRODUCE_CYCLE
}

//https://adventofcode.com/2021/day/6

func day6() (int, int) {
	return day6Task1(), day6Task2()
}

func day6Task1() int {
	data := strings.Split(fileToStringArray("input/6/input.txt")[0], ",")
	lanternfishes := make([]*Lanternfish, 0)
	var days int
	for _, daysToReproduce := range data {
		days, _ = strconv.Atoi(daysToReproduce)
		lanternfishes = append(lanternfishes, &Lanternfish{days, true})
	}

	for i := 0; i < DAYS_NUMBER; i++ {
		lanternfishes = liveOneDay(lanternfishes)
	}

	return len(lanternfishes)
}

func day6Task2() int {
	return 0
}

func liveOneDay(lanternfishes []*Lanternfish) []*Lanternfish {
	lanternfishNumber := len(lanternfishes)
	for i := 0; i < lanternfishNumber; i++ {
		lanternfishes[i].liveOneDay()
		if lanternfishes[i].canReproduce() {
			lanternfishes = append(lanternfishes, &Lanternfish{FIRST_REPRODUCE_CYCLE, false})
		}
	}
	return lanternfishes
}
