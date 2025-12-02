package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/tossthedev/aoc_2025/internal/utils"
)

type IDRange struct {
	Start int
	End   int
}

func main() {
	input := utils.ReadInput("day02")
	ids, err := readIds(input)

	if err != nil {
		log.Fatalf("error while reading IDs: %v", err)
	}

	fmt.Println("Part1: ", part1(ids))
	fmt.Println("Part2: ", part2(ids))
}

func part1(ids []IDRange) int {
	return sumInvalid(ids, part1InvalidID)
}

func part2(ids []IDRange) int {
	return sumInvalid(ids, part2InvalidID)
}

func sumInvalid(ids []IDRange, invalid func(int) bool) int {
	sum := 0

	for _, id := range ids {
		for i := id.Start; i <= id.End; i++ {
			if invalid(i) {
				sum += i
			}
		}
	}

	return sum
}

func readIds(input string) ([]IDRange, error) {
	parts := strings.Split(input, ",")
	var ids []IDRange

	for _, part := range parts {
		points := strings.Split(part, "-")
		start, err := strconv.Atoi(points[0])
		if err != nil {
			return nil, fmt.Errorf("unexpected start: %q, %v", points[0], err)
		}

		end, err := strconv.Atoi(points[1])

		if err != nil {
			return nil, fmt.Errorf("unexpected end: %q, %v", points[1], err)
		}

		ids = append(ids, IDRange{
			Start: start,
			End:   end,
		})
	}

	return ids, nil
}

func part1InvalidID(id int) bool {
	s := strconv.Itoa(id)

	if len(s)%2 != 0 {
		return false
	}

	half := len(s) / 2
	return s[:half] == s[half:]
}

func part2InvalidID(id int) bool {
	s := strconv.Itoa(id)
	n := len(s)

	for size := 1; size <= n/2; size++ {
		if n%size != 0 {
			continue
		}

		pattern := s[:size]

		ok := true
		for i := size; i < n; i += size {
			if s[i:i+size] != pattern {
				ok = false
				break
			}
		}

		if ok {
			return true
		}
	}

	return false
}
