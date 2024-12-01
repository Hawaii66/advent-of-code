package main

import (
	"fmt"
	utils "hawaiidev/advent-of-code/go"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadLines("../input.txt")

	arr1 := make([]int, len(lines))
	arr2 := make([]int, len(lines))

	for i := 0; i < len(lines); i++ {
		a1, _ := strconv.Atoi(strings.Split(lines[i], "   ")[0])
		a2, _ := strconv.Atoi(strings.Split(lines[i], "   ")[1])

		arr1[i] = a1
		arr2[i] = a2
	}

	utils.SortSlice(arr1, func(a, b int) bool { return a < b })
	utils.SortSlice(arr2, func(a, b int) bool { return a < b })

	sum := 0
	for i := 0; i < len(lines); i++ {
		sum += (int(math.Abs(float64(arr1[i] - arr2[i]))))
	}

	fmt.Println(sum)
}
