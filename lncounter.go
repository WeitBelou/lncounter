package lncounter

import (
	"io"
	"strings"
)

func CountLines(r io.Reader) (int, error) {
	const BufferSize = 1024
	buffer := make([]byte, BufferSize)

	const LineSeparator = "\n"

	nLines := 1 // Empty file has one line

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

func main() {

}
