package main

import (
	"fmt"
	utils "hawaiidev/advent-of-code/go"
	"math"
)

func main() {
	lines := utils.ReadLines("../input.txt")

	good := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]

		numbers := utils.StringToIntSlice(utils.SplitLineOnSpaces(line))

		isGood := isGoodLineWithDampener(numbers)
		if isGood {
			good += 1
		}
	}

	fmt.Println(good)
}

func isGoodLineWithDampener(line []int) bool {
	if isGoodLine(line) {
		return true
	}

	for i := 0; i < len(line); i++ {
		curr := make([]int, len(line))
		copy(curr, line)
		if isGoodLine(remove(curr, i)) {
			return true
		}
	}

	return false
}

func isGoodLine(line []int) bool {
	isIncreasing := false
	prev := line[0]

	for i := 1; i < len(line); i++ {
		curr := line[i]
		diff := curr - prev
		if i == 1 {
			if diff > 0 {
				isIncreasing = true
			}
		}

		if diff == 0 || math.Abs(float64(diff)) > 3 {
			return false
		}
		if isIncreasing && diff < 0 {
			return false
		}
		if !isIncreasing && diff > 0 {
			return false
		}

		prev = curr
	}

	return true
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
