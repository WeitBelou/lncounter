package main

import (
	"github.com/weitbelou/lncounter"
	"log"
	"os"
)

func main() {
	nArgs := len(os.Args)
	if nArgs == 1 {
		log.Print("Unexpected number of arguments", nArgs-1)
		os.Exit(1)
	}
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
