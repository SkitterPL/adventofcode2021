package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	//fmt.Println(day1())
	//fmt.Println(day2())
	//fmt.Println(day3())
	//fmt.Println(day4())
	fmt.Println(day5())
	elapsed := time.Since(start)
	fmt.Printf("Task took %s\n", elapsed)
}
