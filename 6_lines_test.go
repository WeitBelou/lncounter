package lncounter

import (
	"os"
	"testing"
)

func Test7Lines(t *testing.T) {
	file, _ := os.Open("./6_lines.txt")

	nLines, err := CountLines(file)

	if err != nil {
		t.Error("Someting happens:", err)
	}

	const NLines = 6
	if nLines != NLines {
		t.Error("Expected", NLines, "lines but", nLines, "given")
	}

	file.Close()
}
