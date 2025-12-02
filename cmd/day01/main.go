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
	input_lines := utils.Lines(input)
	instructions := parseInstructions(input_lines)

	fmt.Println("Part1 Code - ", part1(instructions))
	fmt.Println("Part2 Code - ", part2(instructions))
}

func parseInstructions(lines []string) []instruction {
	instructions := make([]instruction, 0, len(lines))

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		dir := line[0]
		steps, err := strconv.Atoi(line[1:])

		if err != nil {
			log.Fatalf("Invalid step value in %q: %v", line, err)
		}

		instructions = append(instructions, instruction{
			dir:   dir,
			steps: steps,
		})
	}

	return instructions
}

func part1(instructions []instruction) int {
	currentValue := startValue
	code := 0

	for _, inst := range instructions {
		switch inst.dir {
		case 'R':
			currentValue = (currentValue + inst.steps) % maxPoint
		case 'L':
			currentValue = (currentValue - inst.steps) % maxPoint
		default:
			log.Fatalf("Unknown direction: %q", inst.dir)
		}

		if currentValue == 0 {
			code++
		}
	}

	return code
}

func part2(instructions []instruction) int {
	currentValue := startValue
	code := 0

	for _, inst := range instructions {
		for i := 0; i < inst.steps; i++ {
			currentValue = stepOnce(currentValue, inst.dir, maxPoint)

			if currentValue == 0 {
				code++
			}
		}
	}

	return code
}

func stepOnce(current int, dir byte, maxValue int) int {
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
		log.Fatalf("Unknown direction: %q", dir)
	}

	return current
}
