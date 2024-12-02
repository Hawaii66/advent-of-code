package utils

import (
	"sort"
	"strconv"
)

func SortSlice[T comparable](slice []T, compare func(a, b T) bool) {
	sort.Slice(slice, func(i, j int) bool {
		return compare(slice[i], slice[j])
	})
}

func StringToIntSlice(slice []string) []int {
	ints := make([]int, len(slice))
	for i := 0; i < len(slice); i++ {
		a, err := strconv.Atoi(slice[i])
		if err != nil {
			panic(err)
		}

		ints[i] = a
	}
	return ints
}
