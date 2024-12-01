package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	lines := strings.Split(input, "\n")

	arr1 := make([]int, len(lines))
	arr2 := make([]int, len(lines))

	for i := 0; i < len(lines); i++ {
		a1, _ := strconv.Atoi(strings.Split(lines[i], "   ")[0])
		a2, _ := strconv.Atoi(strings.Split(lines[i], "   ")[1])

		arr1[i] = a1
		arr2[i] = a2
	}

	occurances := make(map[int]int)

	for i := 0; i < len(lines); i++ {
		num := arr2[i]
		prev := occurances[num]
		occurances[num] = prev + 1
	}

	similarity := 0

	for i := 0; i < len(lines); i++ {
		times := occurances[arr1[i]]

		similarity += times * arr1[i]
	}

	println(similarity)
}
