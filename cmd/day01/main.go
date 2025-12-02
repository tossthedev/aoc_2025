package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/tossthedev/aoc_2025/internal/utils"
)

const (
	startValue = 50
	maxPoint   = 100
)

type instruction struct {
	dir   byte
	steps int
}

func main() {
	input := utils.ReadInput("day01")
	inputLines := utils.Lines(input)
	instructions, err := parseInstructions(inputLines)

	if err != nil {
		log.Fatalf("parseInstructions failed: %v", err)
	}

	part1Code, err := part1(instructions)

	if err != nil {
		log.Fatalf("part1 failed: %v", err)
	}

	part2Code, err := part2(instructions)

	if err != nil {
		log.Fatalf("part2 failed: %v", err)
	}

	fmt.Println("part1 code - ", part1Code)
	fmt.Println("part2 code - ", part2Code)
}

func parseInstructions(lines []string) ([]instruction, error) {
	instructions := make([]instruction, 0, len(lines))

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		dir := line[0]
		steps, err := strconv.Atoi(line[1:])

		if err != nil {
			return nil, fmt.Errorf("invalid step value in %q: %v", line, err)
		}

		instructions = append(instructions, instruction{
			dir:   dir,
			steps: steps,
		})
	}

	return instructions, nil
}

func part1(instructions []instruction) (int, error) {
	currentValue := startValue
	code := 0

	for _, inst := range instructions {
		switch inst.dir {
		case 'R':
			currentValue = (currentValue + inst.steps) % maxPoint
		case 'L':
			currentValue = (currentValue - inst.steps) % maxPoint
		default:
			return 0, fmt.Errorf("unknown direction: %q", inst.dir)
		}

		if currentValue == 0 {
			code++
		}
	}

	return code, nil
}

func part2(instructions []instruction) (int, error) {
	currentValue := startValue
	code := 0
	var err error

	for _, inst := range instructions {
		for i := 0; i < inst.steps; i++ {
			currentValue, err = stepOnce(currentValue, inst.dir, maxPoint)

			if err != nil {
				return 0, fmt.Errorf("stepOnce failed: %v", err)
			}

			if currentValue == 0 {
				code++
			}
		}
	}

	return code, nil
}

func stepOnce(current int, dir byte, maxValue int) (int, error) {
	switch dir {
	case 'R':
		current++
		if current == maxValue {
			current = 0
		}
	case 'L':
		current--

		if current < 0 {
			current = maxValue - 1
		}
	default:
		return 0, fmt.Errorf("unknown direction: %q", dir)
	}

	return current, nil
}
