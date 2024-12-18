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

func ReadLine(path string) string {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return string(data)
}

func SplitLineOnSpaces(line string) []string {
	return strings.Fields(line)
}
