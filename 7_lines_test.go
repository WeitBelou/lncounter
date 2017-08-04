package lncounter

import (
	"os"
	"testing"
)

func Test7Lines(t *testing.T) {
	file, _ := os.Open("./7_lines.txt")

	nLines, err := CountLines(file)

	if err != nil {
		t.Error("Someting happens:", err)
	}

	if nLines != 7 {
		t.Error("Expected 7 lines but", nLines, "given")
	}

	file.Close()
}
