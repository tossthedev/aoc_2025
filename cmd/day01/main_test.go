package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	instructions := []instruction{
		{steps: 68, dir: 'L'},
		{steps: 30, dir: 'L'},
		{steps: 48, dir: 'R'},
		{steps: 5, dir: 'L'},
		{steps: 60, dir: 'R'},
		{steps: 55, dir: 'L'},
		{steps: 1, dir: 'L'},
		{steps: 99, dir: 'L'},
		{steps: 14, dir: 'R'},
		{steps: 82, dir: 'L'},
	}

	got, err := part1(instructions)

	if err != nil {
		t.Errorf("part1 Failed: %v", err)
	}

	want := 3

	if got != want {
		t.Errorf("part1() = %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	instructions := []instruction{
		{steps: 68, dir: 'L'},
		{steps: 30, dir: 'L'},
		{steps: 48, dir: 'R'},
		{steps: 5, dir: 'L'},
		{steps: 60, dir: 'R'},
		{steps: 55, dir: 'L'},
		{steps: 1, dir: 'L'},
		{steps: 99, dir: 'L'},
		{steps: 14, dir: 'R'},
		{steps: 82, dir: 'L'},
	}

	got, err := part2(instructions)

	if err != nil {
		t.Errorf("part2 failed: %v", err)
	}

	want := 6

	if got != want {
		t.Errorf("part2() = %d; want %d", got, want)
	}
}

func TestInvalidDirectionPart1(t *testing.T) {
	instructions := []instruction{
		{steps: 20, dir: 'K'},
	}

	_, err := part1(instructions)

	if err == nil {
		t.Fatalf("expected error for unknown direction, got nil")
	}

	want := "unknown direction: 'K'"

	if err.Error() != want {
		t.Fatalf("unexpected error message\n got: %q\nwant: %q", err.Error(), want)
	}
}

func TestStepOnceUnknownDirection(t *testing.T) {
	current, dir, maxValue := 50, byte('K'), 100

	_, err := stepOnce(current, dir, maxValue)

	if err == nil {
		t.Fatalf("expected error for unknown direction, got nil")
	}

	want := "unknown direction: 'K'"

	if err.Error() != want {
		t.Fatalf("unexpected error message\n got: %q\nwant: %q", err.Error(), want)
	}
}

func TestStepOnceFails(t *testing.T) {
	instructions := []instruction{
		{steps: 20, dir: 'K'},
	}

	_, err := part2(instructions)

	if err == nil {
		t.Fatalf("expected error for unknown direction, got nil")
	}

	want := "stepOnce failed: unknown direction: 'K'"

	if err.Error() != want {
		t.Fatalf("unexpected error message\n got: %q\nwant: %q", err.Error(), want)
	}
}

func TestParseInstructions(t *testing.T) {
	input := "L40\nR30\nL32\nR60"

	want := []instruction{
		{steps: 40, dir: 'L'},
		{steps: 30, dir: 'R'},
		{steps: 32, dir: 'L'},
		{steps: 60, dir: 'R'},
	}

	got, err := parseInstructions(strings.Split(input, "\n"))

	if err != nil {
		t.Fatalf("parseInstructions failed: %v", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("mismatch: want=%v, got=%v", want, got)
	}
}

func TestStepOnceWrapRight(t *testing.T) {
	current, err := stepOnce(99, 'R', 100)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if current != 0 {
		t.Fatalf("expected 0, got %d", current)
	}
}

func TestStepOnceWrapLeft(t *testing.T) {
	current, err := stepOnce(0, 'L', 100)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if current != 99 {
		t.Fatalf("expected 99, got %d", current)
	}
}

func TestPart1Empty(t *testing.T) {
	got, err := part1(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 0 {
		t.Fatalf("expected 0, got %d", got)
	}
}

func TestPart2Empty(t *testing.T) {
	got, err := part2(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 0 {
		t.Fatalf("expected 0, got %d", got)
	}
}

func TestParseInstructionsInvalidStepValue(t *testing.T) {
	input := []string{"Lxx"}
	_, err := parseInstructions(input)
	if err == nil {
		t.Fatalf("expected error for invalid step value, got nil")
	}

	if !strings.Contains(err.Error(), "invalid step value") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestParseInstructionsSkipsEmptyLines(t *testing.T) {
	input := []string{"L10", "", "R20"}
	got, err := parseInstructions(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []instruction{
		{steps: 10, dir: 'L'},
		{steps: 20, dir: 'R'},
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("mismatch: want=%v, got=%v", want, got)
	}
}
