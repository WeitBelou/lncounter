package lncounter

import (
	"testing"
	"os"
)

func Test7Lines(t *testing.T) {
	file, _ := os.Open("./7_lines.go")

	nLines, err := CountLines(file)

	if err != nil {
		t.Error(err)
	}

	if nLines != 7 {
		t.Error()
	}
}
