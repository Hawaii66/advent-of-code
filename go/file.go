package utils

import (
	"os"
	"strings"
)

func ReadLines(path string) []string {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	content := string(data)

	return strings.Split(content, "\n")
}

func SplitLineOnSpaces(line string) []string {
	return strings.Fields(line)
}
