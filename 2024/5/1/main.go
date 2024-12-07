package main

import (
	"fmt"
	utils "hawaiidev/advent-of-code/go"
	"strconv"
	"strings"
)

func main() {
	input := utils.ReadLines("../input.txt")

	pairs := make(map[string][]string)

	var lines []string
	for _, line := range input {
		if strings.Contains(line, "|") {
			splitted := strings.Split(line, "|")
			num1 := splitted[0]
			num2 := splitted[1]

			prev := pairs[num1]
			prev = append(prev, num2)
			pairs[num1] = prev
		} else {
			lines = append(lines, line)
		}
	}

	total := 0
	for _, line := range lines {
		nums := strings.Split(line, ",")
		toCheck := []string{nums[0]}
		allowedLine := true
		for _, num := range nums {
			for _, check := range toCheck {
				for _, toComp := range pairs[num] {
					if toComp == check {
						allowedLine = false
						break
					}
				}
			}

			toCheck = append(toCheck, num)
		}

		if allowedLine {
			num, err := strconv.Atoi(nums[len(nums)/2])
			if err != nil {
				panic(err)
			}
			total += num
		}
	}

	fmt.Println(total)
}
