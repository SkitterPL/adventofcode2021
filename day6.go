package main

import (
	"strconv"
	"strings"
)

const TASK_1_DAYS_NUMBER = 80
const TASK_2_DAYS_NUMBER = 256

//https://adventofcode.com/2021/day/6

func day6() (int, int) {
	data := strings.Split(fileToStringArray("input/6/input.txt")[0], ",")
	return calculateFishNumber(TASK_1_DAYS_NUMBER, data), calculateFishNumber(TASK_2_DAYS_NUMBER, data)
}

func calculateFishNumber(daysNumber int, data []string) int {
	days := make([]int, 9)
	for _, daysToReproduce := range data {
		day, _ := strconv.Atoi(daysToReproduce)
		days[day]++
	}

	temp := 0
	for i := 0; i < daysNumber; i++ {
		for j := 0; j < 8; j++ {
			days[j] = days[j+1]
		}

		days[8] = temp
		days[6] += temp
		temp = days[0]
	}

	return sum(days...)
}
