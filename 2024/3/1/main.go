package main

import (
	"fmt"
	utils "hawaiidev/advent-of-code/go"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	line := utils.ReadLine("../input.txt")

	mulInstruction := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	total := 0

	matches := mulInstruction.FindAllString(line, -1)
	for i := 0; i < len(matches); i++ {
		match := matches[i]

		splitted := strings.Split(match, ",")
		num1Str := strings.Split(splitted[0], "(")[1]
		num2Str := strings.Split(splitted[1], ")")[0]

		num1, err := strconv.Atoi(num1Str)
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(num2Str)
		if err != nil {
			panic(err)
		}

		total += num1 * num2
	}

	fmt.Println(total)
}
