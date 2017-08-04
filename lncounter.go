package lncounter

import (
	"io"
	"strings"
)

func CountLines(r io.Reader) (int, error) {
	const BufferSize = 1024
	buffer := make([]byte, BufferSize)

	const LineSeparator = "\n"

	nLines := 0

	for {
		nBytes, err := r.Read(buffer)
		nLines += strings.Count(string(buffer[:nBytes]), LineSeparator)

		switch {
		case err == io.EOF:
			return nLines, nil

		case err != nil:
			return nLines, err
		}
	}
}
