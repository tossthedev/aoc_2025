package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadInput_ReturnsFileContent(t *testing.T) {
	tempDir := t.TempDir()
	inputDir := filepath.Join(tempDir, "inputs")
	if err := os.Mkdir(inputDir, 0755); err != nil {
		t.Fatalf("failed creating temp inputs dir: %v", err)
	}

	want := "hello world"
	filePath := filepath.Join(inputDir, "day01.txt")
	if err := os.WriteFile(filePath, []byte(want+"\n"), 0644); err != nil {
		t.Fatalf("failed writing temp file: %v", err)
	}

	origWD, _ := os.Getwd()
	os.Chdir(tempDir)
	t.Cleanup(func() { os.Chdir(origWD) })

	got := ReadInput("day01")
	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}

func TestReadInput_PanicsIfFileMissing(t *testing.T) {
	tempDir := t.TempDir()
	origWD, _ := os.Getwd()
	os.Chdir(tempDir)
	t.Cleanup(func() { os.Chdir(origWD) })

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic but got none")
		}
	}()

	ReadInput("nope")
}

func TestLines_SplitsIntoLines(t *testing.T) {
	input := "first\nsecond\nthird"
	want := []string{"first", "second", "third"}

	got := Lines(input)

	if len(got) != len(want) {
		t.Fatalf("expected %d lines, got %d", len(want), len(got))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("line %d: expected %q, got %q", i, want[i], got[i])
		}
	}
}

func TestLines_HandlesEmptyString(t *testing.T) {
	got := Lines("")
	if len(got) != 0 {
		t.Fatalf("expected empty slice, got %#v", got)
	}
}

func TestLines_HandlesTrailingNewline(t *testing.T) {
	input := "a\nb\nc\n"
	want := []string{"a", "b", "c"}

	got := Lines(input)
	if len(got) != len(want) {
		t.Fatalf("expected %d lines, got %d", len(want), len(got))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("line %d: expected %q, got %q", i, want[i], got[i])
		}
	}
}
