package main

import "fmt"

const (
	DaysInWeekCount = 7
)

type Name string

func Sum(x, y int64) int64 {
	fmt.Println("ХУЕТА")
	return x + y
}

func getWeeksCount(daysCount int) int {
	return daysCount / DaysInWeekCount
}
