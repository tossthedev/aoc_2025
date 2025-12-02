package utils

import (
	"bufio"
	"os"
	"strings"
)

func ReadInput(day string) string {
	data, err := os.ReadFile("inputs/" + day + ".txt")
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(data))
}

func Lines(s string) []string {
	lines := []string{}
	scanner := bufio.NewScanner(strings.NewReader(s))

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

