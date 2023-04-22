package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Do struct {
	Done bool
	Task string
}

func readFile(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	var lines []string
	input := bufio.NewScanner(f)
	for input.Scan() {
		lines = append(lines, input.Text())
	}
	f.Close()

	return lines
}

func readDos(file string) []Do {
	doneStr := "- [x] "
	undoneStr := "- [ ] "

	lines := readFile(file)
	var dos []Do
	for _, line := range lines {
		var do Do
		if strings.HasPrefix(line, undoneStr) {
			do.Done = false
		} else if strings.HasPrefix(line, doneStr) {
			do.Done = true
		} else {
			// quietly ignore non-do lines
			break
		}
		do.Task = line[len(doneStr):]
		dos = append(dos, do)
	}
	return dos
}

func main() {
	dos := readDos(os.Args[1])
	for i, do := range dos {
		fmt.Printf("%d. %s (done? %v)\n", i+1, do.Task, do.Done)
	}
}
