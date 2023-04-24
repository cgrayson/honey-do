package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var doneCheckboxStr string = "- [x]"
var undoneCheckboxStr string = "- [ ]"

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

func writeFile(file string, lines []string) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(f)
	for _, line := range lines {
		_, err := fmt.Fprintf(w, "%s\n", line)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = w.Flush()
	if err != nil {
		log.Println(err)
		return
	}
}

func readDos(file string) []Do {
	lines := readFile(file)
	var dos []Do
	for _, line := range lines {
		var do Do
		if strings.HasPrefix(line, undoneCheckboxStr) {
			do.Done = false
		} else if strings.HasPrefix(line, doneCheckboxStr) {
			do.Done = true
		} else {
			// quietly ignore non-do lines
			break
		}
		do.Task = line[len(doneCheckboxStr)+1:] // add one for space after checkbox
		dos = append(dos, do)
	}
	return dos
}

func writeDos(file string, dos []Do) {
	// make two slices so that undone items are all first
	var undoneLines []string
	var doneLines []string
	for _, do := range dos {
		if do.Done {
			doneLines = append(doneLines, fmt.Sprintf("%s %s", doneCheckboxStr, do.Task))
		} else {
			undoneLines = append(undoneLines, fmt.Sprintf("%s %s", undoneCheckboxStr, do.Task))
		}
	}
	writeFile(file, append(undoneLines, doneLines...))
}

func pullDo(dos []Do) {
	undoneCount := 0
	for _, do := range dos {
		if !do.Done {
			undoneCount++
		}
	}
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(undoneCount)

	undoneCount = 0
	for i, do := range dos {
		if !do.Done {
			undoneCount++
		}
		if i == r {
			fmt.Println(dos[i].Task)
			dos[i].Done = true
			break
		}
	}
}

func parseCommandLine(args []string) (string, string, string) {
	filename := args[1]
	action := "pull"
	var task string

	if len(args) > 2 {
		action = args[2]
	}

	if action == "add" && len(args) > 2 {
		task = args[3]
	}

	return filename, action, task
}

func main() {
	filename, action, task := parseCommandLine(os.Args)
	dos := readDos(filename)

	switch action {
	case "pull":
		pullDo(dos)
	case "add":
		newDo := Do{Done: false, Task: task}
		dos = append(dos, newDo)
	}
	writeDos(filename, dos)
}
