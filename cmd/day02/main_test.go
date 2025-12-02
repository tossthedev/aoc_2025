package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	ranges := []IDRange{
		{Start: 11, End: 22},
		{Start: 95, End: 115},
		{Start: 998, End: 1012},
		{Start: 1188511880, End: 1188511890},
		{Start: 222220, End: 222224},
		{Start: 1698522, End: 1698528},
		{Start: 446443, End: 446449},
		{Start: 38593856, End: 38593862},
	}

	got := part1(ranges)
	want := 1227775554

	if got != want {
		t.Errorf("part1() = %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	ranges := []IDRange{
		{Start: 11, End: 22},
		{Start: 95, End: 115},
		{Start: 998, End: 1012},
		{Start: 1188511880, End: 1188511890},
		{Start: 222220, End: 222224},
		{Start: 1698522, End: 1698528},
		{Start: 446443, End: 446449},
		{Start: 38593856, End: 38593862},
		{Start: 565653, End: 565659},
		{Start: 824824821, End: 824824827},
		{Start: 2121212118, End: 2121212124},
	}

	got := part2(ranges)
	want := 4174379265

	if got != want {
		t.Errorf("part2() = %d; want %d", got, want)
	}
}

func TestPart1InvalidID(t *testing.T) {
	tests := []struct {
		id   int
		want bool
	}{
		{11, true},
		{22, true},
		{1010, true},
		{1234, false},
		{123123, true},
	}

	for _, tt := range tests {
		got := part1InvalidID(tt.id)
		if got != tt.want {
			t.Errorf("part1InvalidID(%d) = %v, want %v",
				tt.id, got, tt.want)
		}
	}
}

func TestPart2InvalidID(t *testing.T) {
	tests := []struct {
		id   int
		want bool
	}{
		{11, true},
		{111, true},
		{123123, true},
		{565656, true},
		{38593859, true},
		{12345, false},
		{1698522, false},
	}

	for _, tt := range tests {
		got := part2InvalidID(tt.id)
		if got != tt.want {
			t.Errorf("part2InvalidID(%d) = %v, want %v",
				tt.id, got, tt.want)
		}
	}
}

func TestReadIDs(t *testing.T) {
	input := "11-22"
	got, err := readIds(input)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	want := IDRange{
		Start: 11,
		End:   22,
	}

	idRange := got[0]

	if want.Start != idRange.Start {
		t.Errorf("start want = %v, got = %v", want.Start, idRange.Start)
	}

	if want.End != idRange.End {
		t.Errorf("end want = %v, got = %v", want.Start, idRange.Start)
	}
}

func TestReadIDsIncorrectStart(t *testing.T) {
	input := "11.2-22"

	_, err := readIds(input)

	if err == nil {
		t.Fatalf("expected error for unknown direction, got nil")
	}

	want := "unexpected start: \"11.2\", strconv.Atoi: parsing \"11.2\": invalid syntax"

	if want != err.Error() {
		t.Fatalf("unexpected error message\n got: %q\nwant: %q", err.Error(), want)
	}
}

func TestReadIDsIncorrectEnd(t *testing.T) {
	input := "11-22.2"

	_, err := readIds(input)

	if err == nil {
		t.Fatalf("expected error for unknown direction, got nil")
	}

	want := "unexpected end: \"22.2\", strconv.Atoi: parsing \"22.2\": invalid syntax"

	if want != err.Error() {
		t.Fatalf("unexpected error message\n got: %q\nwant: %q", err.Error(), want)
	}
}
