package main

import (
	utils "hawaiidev/advent-of-code/go"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadLines("../input.txt")

	arr1 := make([]int, len(lines))
	occurances := make(map[int]int)

	for i := 0; i < len(lines); i++ {
		a1, _ := strconv.Atoi(strings.Split(lines[i], "   ")[0])
		a2, _ := strconv.Atoi(strings.Split(lines[i], "   ")[1])

		arr1[i] = a1

		prev := occurances[a2]
		occurances[a2] = prev + 1
	}

	similarity := 0

	for i := 0; i < len(lines); i++ {
		times := occurances[arr1[i]]

		similarity += times * arr1[i]
	}

	println(similarity)
}
