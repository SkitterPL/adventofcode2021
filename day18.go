package main

import (
	"encoding/json"
	"fmt"
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

func (number *SnailfishNumber) print() string {
	if number.isRegular() {
		return fmt.Sprint(number.value.value)
	}
	return "[" + fmt.Sprint(number.x.print()+","+number.y.print()) + "]"
}

func recreateFromString(value string) *SnailfishNumber {
	var parsedData []interface{}
	err := json.Unmarshal([]byte(value), &parsedData)
	check(err)
	return newSnailfishNumber(parsedData)
}

func newSnailfishNumber(value []interface{}) *SnailfishNumber {
	var xSnail, ySnail *SnailfishNumber
	var snailFish *SnailfishNumber
	xFloat, ok := value[0].(float64)
	if !ok {
		x := value[0].([]interface{})
		xSnail = newSnailfishNumber(x)
	} else {
		xSnail = &SnailfishNumber{value: &RegularNumber{xFloat}}
	}

	yFloat, ok := value[1].(float64)
	if !ok {
		y := value[1].([]interface{})
		ySnail = newSnailfishNumber(y)
	} else {
		ySnail = &SnailfishNumber{value: &RegularNumber{yFloat}}
	}

	snailFish = &SnailfishNumber{x: xSnail, y: ySnail}
	return snailFish
}

func add(num1 *SnailfishNumber, num2 *SnailfishNumber) *SnailfishNumber {
	newNumber := (&SnailfishNumber{x: num1, y: num2}).copy()
	newNumber.recalculateParents(nil)
	for {
		if newNumber.explode(0) {
			newNumber.recalculateParents(nil)
			continue
		}
		if newNumber.split() {
			newNumber.recalculateParents(nil)
			continue
		}
		break
	}
	return newNumber
}

func (number *SnailfishNumber) explode(nestedLevel int) bool {
	if nestedLevel == 4 && number.x.isRegular() && number.y.isRegular() {
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
	} else {
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
	}
	return false
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
	if number.isRegular() {
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
