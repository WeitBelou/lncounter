package main

import (
	"github.com/weitbelou/lncounter"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	nLines, err := lncounter.CountLines(file)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	println(nLines, filename)

	file.Close()
}
