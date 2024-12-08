package main

import (
	"fmt"
	utils "hawaiidev/advent-of-code/go"
	"strconv"
	"strings"
)

type SumTest struct {
	total   int
	numbers []int
}

func main() {
	lines := utils.ReadLines("../input.txt")

	total := 0
	for _, line := range lines {
		split := strings.Split(line, ":")
		tot, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}

		nums := utils.StringToIntSlice(utils.SplitLineOnSpaces(split[1]))

		works := doesWork(SumTest{
			total:   tot,
			numbers: nums,
		})

		if works {
			total += tot
		}
	}

	fmt.Println(total)
}

func doesWork(test SumTest) bool {
	if len(test.numbers) < 2 {
		return false
	}

	num1 := test.numbers[0]
	num2 := test.numbers[1]

	if len(test.numbers) == 2 {
		if num1+num2 == test.total {
			return true
		}
		if num1*num2 == test.total {
			return true
		}

		con, err := strconv.Atoi(strconv.Itoa(num1) + strconv.Itoa(num2))
		if err != nil {
			panic(err)
		}
		if con == test.total {
			return true
		}
	}

	con, err := strconv.Atoi(strconv.Itoa(num1) + strconv.Itoa(num2))
	if err != nil {
		panic(err)
	}
	newNumbersCat := []int{con}
	newNumbersCat = append(newNumbersCat, test.numbers[2:]...)

	workCat := doesWork(SumTest{
		total:   test.total,
		numbers: newNumbersCat,
	})
	if workCat {
		return true
	}

	newNumbersAdd := []int{num1 + num2}
	newNumbersAdd = append(newNumbersAdd, test.numbers[2:]...)

	workAdd := doesWork(SumTest{
		total:   test.total,
		numbers: newNumbersAdd,
	})
	if workAdd {
		return true
	}

	newNumbersMult := []int{num1 * num2}
	newNumbersMult = append(newNumbersMult, test.numbers[2:]...)

	workMult := doesWork(SumTest{
		total:   test.total,
		numbers: newNumbersMult,
	})

	return workMult
}
