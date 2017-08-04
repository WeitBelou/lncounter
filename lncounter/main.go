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

func processSeveralFiles(files []string) {
	nFiles := len(files)

	for i := 0; i < nFiles; i++ {
		processFile(files[i])
	}
}

func main() {
	nArgs := len(os.Args)
	if nArgs == 1 {
		log.Print("Unexpected number of arguments", nArgs-1)
		os.Exit(1)
	}

	filenames := os.Args[1:]
	nFiles := len(filenames)

	if nFiles == 1 {
		processFile(filenames[0])
	} else {
		processSeveralFiles(filenames)
	}
}
