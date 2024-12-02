package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func promptUser(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("Input cannot be empty. Please try again.")
			continue
		}
		return input
	}
}

func updateEnvFile(envVars map[string]string, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating .env file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for key, value := range envVars {
		_, err := writer.WriteString(fmt.Sprintf("%s=%s\n", key, value))
		if err != nil {
			return fmt.Errorf("error writing to .env file: %w", err)
		}
	}
	return writer.Flush()
}

func checkAndPromptEnv() {
	envVars, err := godotenv.Read(".env")
	if err != nil {
		envVars = make(map[string]string)
	}

	if _, ok := envVars["COOKIE"]; !ok {
		envVars["COOKIE"] = promptUser("Enter COOKIE: ")
	}
	if _, ok := envVars["YEAR"]; !ok {
		envVars["YEAR"] = promptUser("Enter YEAR: ")
	}

	err = updateEnvFile(envVars, ".env")
	if err != nil {
		fmt.Printf("Error updating .env file: %v\n", err)
	}
}

func fetchAdventOfCodeInput(sessionCookie, year, day string) ([]byte, error) {
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, day)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: sessionCookie})

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response code: %d", resp.StatusCode)
	}

	text, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return text, nil
}

func removeEmptyLines(input string) string {
	lines := strings.Split(input, "\n")

	var filteredLines []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			filteredLines = append(filteredLines, line)
		}
	}

	return strings.Join(filteredLines, "\n")
}

func emptyFile(path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString("")
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

func main() {
	checkAndPromptEnv()

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		return
	}

	cookie := os.Getenv("COOKIE")
	year := os.Getenv("YEAR")
	day := promptUser("Date: ")

	fmt.Println("Getting AOC", year, day)
	body, err := fetchAdventOfCodeInput(cookie, year, day)
	if err != nil {
		emptyFile(".env")
		fmt.Println("Failed to get AOC, try again but set a new cookie")
		panic(err)
	}

	dirPath := filepath.Join(year, day)
	filePath := filepath.Join(dirPath, "input.txt")
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	str := string(body)
	_, err = file.WriteString(removeEmptyLines(str))
	if err != nil {
		panic(err)
	}
}
