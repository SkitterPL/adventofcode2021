package main

import (
	"encoding/json"
	"math"
)

//https://adventofcode.com/2021/day/18

type SnailfishNumber struct {
	x      *SnailfishNumber
	y      *SnailfishNumber
	value  *RegularNumber
	parent *SnailfishNumber
}

type RegularNumber struct {
	value float64
}

func day18() (float64, float64) {
	data := fileToStringArray("input/18/input.txt")
	number := recreateFromString(data[0])
	snailFishes := make([]*SnailfishNumber, len(data))
	snailFishes[0] = number
	for key, line := range data[1:] {
		snailFishes[key+1] = recreateFromString(line)
		number = add(number, snailFishes[key+1])
	}

	return number.magnitude(), findMaxPossibleMagnitude(snailFishes)
}

func recreateFromString(value string) *SnailfishNumber {
	var parsedData []interface{}
	err := json.Unmarshal([]byte(value), &parsedData)
	check(err)
	return newSnailfishNumber(parsedData)
}

func newSnailfishNumber(value []interface{}) *SnailfishNumber {
	xNumber, yNumber := createNumbersFromJson(value)
	return &SnailfishNumber{x: xNumber, y: yNumber}
}

func createNumbersFromJson(value []interface{}) (*SnailfishNumber, *SnailfishNumber) {
	snailfishNumbers := [2]*SnailfishNumber{}
	for i := 0; i < 2; i++ {
		float, ok := value[i].(float64)
		if ok {
			snailfishNumbers[i] = &SnailfishNumber{value: &RegularNumber{float}}
		} else {
			x := value[i].([]interface{})
			snailfishNumbers[i] = newSnailfishNumber(x)
		}
	}
	return snailfishNumbers[0], snailfishNumbers[1]
}

func add(num1 *SnailfishNumber, num2 *SnailfishNumber) *SnailfishNumber {
	newNumber := (&SnailfishNumber{x: num1, y: num2}).copy()
	for {
		newNumber.recalculateParents(nil)
		if newNumber.explode(0) {
			continue
		}
		if newNumber.split() {
			continue
		}
		break
	}
	return newNumber
}

func (number *SnailfishNumber) explode(nestedLevel int) bool {
	if nestedLevel > 4 {
		return false
	}

	if nestedLevel < 4 {
		if !number.x.isRegular() {
			if number.x.explode(nestedLevel + 1) {
				return true
			}
			if !number.y.isRegular() {
				return number.y.explode(nestedLevel + 1)
			}
		}
		if !number.y.isRegular() {
			return number.y.explode(nestedLevel + 1)
		}
		return false
	}

	parent := number.parent
	var right, left *SnailfishNumber
	if number == parent.x {
		right = parent.findFirstRegularTheRightForX()
		left = parent.findFirstRegularToTheLeftForX()
	} else {
		left = parent.findFirstRegularToTheLeftForY()
		right = parent.findFirstRegularTheRightForY()
	}

	if left != nil {
		left.value.value += number.x.value.value
	}
	if right != nil {
		right.value.value += number.y.value.value
	}

	number.x = nil
	number.y = nil
	number.value = &RegularNumber{0}
	return true
}

func (number *SnailfishNumber) split() bool {
	if number.isRegular() {
		if number.value.value >= 10 {
			number.x = &SnailfishNumber{value: &RegularNumber{math.Floor(number.value.value / 2)}}
			number.y = &SnailfishNumber{value: &RegularNumber{math.Ceil(number.value.value / 2)}}
			number.value = nil
			return true
		}
		return false
	}

	if number.x.split() {
		return true
	}

	if number.y.split() {
		return true
	}
	return false
}

func (number *SnailfishNumber) isRegular() bool {
	return number.value != nil
}

func (number *SnailfishNumber) findFirstRegularTheRightForX() *SnailfishNumber {
	if number.parent == nil {
		return nil
	}

	if number.y.isRegular() {
		return number.y
	}
	if number.parent.x == number {
		return number.y.findFirstRegularToTheLeftForY()
	}
	return getMostLeft(number.y)
}

func (number *SnailfishNumber) findFirstRegularToTheLeftForY() *SnailfishNumber {
	if number.parent == nil {
		return nil
	}
	if number.x.isRegular() {
		return number.x
	}
	if number.parent.y == number {
		return number.parent.x.findFirstRegularTheRightForX()
	}

	return number.parent.findFirstRegularToTheLeftForY()
}

func (number *SnailfishNumber) findFirstRegularTheRightForY() *SnailfishNumber {
	if number.parent == nil {
		return nil
	}
	if number.parent.y == number {
		return number.parent.findFirstRegularTheRightForY()
	}
	if number.parent.y.isRegular() {
		return number.parent.y
	}
	return getMostLeft(number.parent.y)
}

func (number *SnailfishNumber) findFirstRegularToTheLeftForX() *SnailfishNumber {
	if number.parent == nil {
		return nil
	}
	if number.parent.x == number {
		return number.parent.findFirstRegularToTheLeftForX()
	}
	if number.parent.y.isRegular() {
		return number.parent.y
	}
	return getMostRight(number.parent.x)
}

func getMostRight(number *SnailfishNumber) *SnailfishNumber {
	y := number
	for !y.isRegular() {
		y = y.y
	}
	return y
}

func getMostLeft(number *SnailfishNumber) *SnailfishNumber {
	x := number
	for !x.isRegular() {
		x = x.x
	}
	return x
}

func (number *SnailfishNumber) magnitude() float64 {
	if number.isRegular() {
		return number.value.value
	}
	return 3*number.x.magnitude() + 2*number.y.magnitude()
}

func findMaxPossibleMagnitude(numbers []*SnailfishNumber) float64 {
	var maxMagnitude float64
	for _, number1 := range numbers {
		for _, number2 := range numbers {
			if number2 == number1 {
				continue
			}

			mag := add(number1, number2).magnitude()
			if mag > maxMagnitude {
				maxMagnitude = mag
			}
			mag = add(number2, number1).magnitude()
			if mag > maxMagnitude {
				maxMagnitude = mag
			}
		}
	}
	return maxMagnitude
}

func (number *SnailfishNumber) recalculateParents(parent *SnailfishNumber) {
	number.parent = parent
	if number.isRegular() {
		return
	}
	number.x.recalculateParents(number)
	number.y.recalculateParents(number)
}

func (number *SnailfishNumber) copy() *SnailfishNumber {
	if number.isRegular() {
		return &SnailfishNumber{value: &RegularNumber{number.value.value}}
	}
	return &SnailfishNumber{x: number.x.copy(), y: number.y.copy()}
}
