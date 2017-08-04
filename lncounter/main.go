package main

import (
	"github.com/weitbelou/lncounter"
	"log"
	"os"
)

func processFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Print(err)
		return
	}
	defer file.Close()

	nLines, err := lncounter.CountLines(file)

	if err != nil {
		log.Print(err)
	}
	println(nLines, filename)
}

func processFiles(files []string) {
	for _, filename := range files {
		processFile(filename)
	}
}

func processStdIn() {
	file := os.Stdin
	nLines, err := lncounter.CountLines(file)

	if err != nil {
		log.Print(err)
	}
	println(nLines)
}

func main() {
	nArgs := len(os.Args)
	nArgsWithoutProgramName := nArgs - 1

	switch nArgsWithoutProgramName {
	case 0:
		processStdIn()
	case 1:
		processFile(os.Args[1])
	default:
		processFiles(os.Args[1:])
	}
}
