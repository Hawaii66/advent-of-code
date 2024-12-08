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

	var notAllowedLines []string
	for _, line := range lines {
		goodLine := isGoodLine(line, pairs)

		if !goodLine {
			notAllowedLines = append(notAllowedLines, line)
		}
	}

	total := 0
	for _, line := range notAllowedLines {
		fixedLine := fixLine(line, pairs)
		split := strings.Split(fixedLine, ",")
		num, err := strconv.Atoi(split[len(split)/2])
		if err != nil {
			panic(err)
		}
		total += num
	}
	fmt.Println(total)
}

func fixLine(baseLine string, pairs map[string][]string) string {
	line := baseLine
	for !isGoodLine(line, pairs) {
		nums := strings.Split(line, ",")
		toCheck := []string{nums[0]}
		shouldBreak := false
		for i := 1; i < len(nums); i++ {
			if shouldBreak {
				break
			}

			num := nums[i]
			for j, check := range toCheck {
				if shouldBreak {
					break
				}

				for _, toComp := range pairs[num] {
					if shouldBreak {
						break
					}

					if toComp == check {
						nums[j] = nums[i]
						nums[i] = toComp
						shouldBreak = true
					}
				}
			}

			toCheck = append(toCheck, nums[i])
		}
		line = strings.Join(nums, ",")
	}

	return line
}

func isGoodLine(line string, pairs map[string][]string) bool {
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

	return allowedLine
}
