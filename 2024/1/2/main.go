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
