package utils

import (
	"sort"
)

func SortSlice[T comparable](slice []T, compare func(a, b T) bool) {
	sort.Slice(slice, func(i, j int) bool {
		return compare(slice[i], slice[j])
	})
}
