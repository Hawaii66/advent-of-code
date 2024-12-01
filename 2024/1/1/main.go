package main

import (
	"fmt"
	"math"
	"os"
	"sort"
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

	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})

	sum := 0
	for i := 0; i < len(lines); i++ {
		sum += (int(math.Abs(float64(arr1[i] - arr2[i]))))
	}

	fmt.Println(sum)
}
