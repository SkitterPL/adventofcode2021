package main

import (
	"strconv"
)

//https://adventofcode.com/2021/day/16

type Expression struct {
	value          string
	version        int64
	typeId         int64
	literal        bool
	literalValue   int64
	subExpressions []*Expression
}

func newExpression(value string) (*Expression, int) {
	subExpressions := []*Expression{}
	version, _ := strconv.ParseInt(value[0:3], 2, 64)
	typeId, _ := strconv.ParseInt(value[3:6], 2, 64)

	if typeId == 4 {
		literalEnd := false
		index := 6
		literalValue := ""
		for !literalEnd {
			shouldEnd, _ := strconv.ParseInt(value[index:index+1], 2, 64)
			if shouldEnd == 0 {
				literalEnd = true
			}
			literalValue += value[index+1 : index+5]
			index += 5
		}
		intLiteral, _ := strconv.ParseInt(literalValue, 2, 64)
		return &Expression{value: value, version: version, typeId: typeId, literal: true, literalValue: intLiteral}, index
	}
	lengthTypeId, _ := strconv.ParseInt(value[6:7], 2, 64)

	if lengthTypeId == 0 {
		index := 22
		subPacketsNumber, _ := strconv.ParseInt(value[7:index], 2, 64)
		endIndex := index + int(subPacketsNumber)
		for index < endIndex {
			subExpression, newIndex := newExpression(value[index:])
			if subExpression == nil {
				index = endIndex
				break
			}
			index += newIndex
			subExpressions = append(subExpressions, subExpression)
		}
		return &Expression{value: value, version: version, typeId: typeId, literal: false, subExpressions: subExpressions}, index
	}

	index := 18
	subPacketsNumber, _ := strconv.ParseInt(value[7:index], 2, 64)
	for i := 0; i < int(subPacketsNumber); i++ {
		subExpression, newIndex := newExpression(value[index:])
		if subExpression == nil {
			break
		}
		index += newIndex
		subExpressions = append(subExpressions, subExpression)
	}
	return &Expression{value: value, version: version, typeId: typeId, literal: false, subExpressions: subExpressions}, index
}

func day16() (int64, int64) {
	hexString := fileToStringArray("input/16/input.txt")[0]
	mainExpression, _ := newExpression(hex2Bin(hexString))
	return mainExpression.packetVersionSum(), mainExpression.calculateExpression()
}

func hex2Bin(hex string) string {
	bin := ""
	for _, char := range hex {
		i, _ := strconv.ParseInt(string(char), 16, 32)
		number := strconv.FormatInt(i, 2)
		for len(number) != 4 {
			number = "0" + number
		}
		bin += number
	}
	return bin
}

func (expr *Expression) packetVersionSum() int64 {
	sum := expr.version
	for _, subExpr := range expr.subExpressions {
		sum += subExpr.packetVersionSum()
	}
	return sum
}

func (expr *Expression) calculateExpression() int64 {
	switch expr.typeId {
	case 0:
		return expr.sumSubExpressions()
	case 1:
		return expr.multiplySubExpressions()
	case 2:
		return expr.minSubExpression()
	case 3:
		return expr.maxSubExpression()
	case 4:
		return expr.literalValue
	case 5:
		if expr.subExpressions[0].calculateExpression() > expr.subExpressions[1].calculateExpression() {
			return 1
		} else {
			return 0
		}
	case 6:
		if expr.subExpressions[0].calculateExpression() < expr.subExpressions[1].calculateExpression() {
			return 1
		} else {
			return 0
		}
	case 7:
		if expr.subExpressions[0].calculateExpression() == expr.subExpressions[1].calculateExpression() {
			return 1
		} else {
			return 0
		}

	}
	return 0
}

func (expr *Expression) sumSubExpressions() int64 {
	var sum int64
	for _, subExpr := range expr.subExpressions {
		sum += subExpr.calculateExpression()
	}
	return sum
}

func (expr *Expression) multiplySubExpressions() int64 {
	var sum int64
	sum = 1
	for _, subExpr := range expr.subExpressions {
		sum *= subExpr.calculateExpression()
	}
	return sum
}

func (expr *Expression) minSubExpression() int64 {
	min := int64(MaxInt)

	for _, subExpr := range expr.subExpressions {
		value := subExpr.calculateExpression()
		if value < min {
			min = value
		}
	}
	return min
}

func (expr *Expression) maxSubExpression() int64 {
	var max int64

	for _, subExpr := range expr.subExpressions {
		value := subExpr.calculateExpression()
		if value > max {
			max = value
		}
	}
	return max
}
