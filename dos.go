package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"honey-do/utils"
)

var doneCheckboxStr string = "- [x]"
var undoneCheckboxStr string = "- [ ]"

type Metadata struct {
	AddedDate  time.Time
	PulledDate time.Time
}

type Do struct {
	Done     bool
	Task     string
	Metadata Metadata
}

// todo next time: write tests for this
func CountUndone(dos []Do) int {
	undoneCount := 0
	for _, do := range dos {
		if !do.Done {
			undoneCount++
		}
	}
	return undoneCount
}

func readDos(file string) []Do {
	lines := utils.ReadFile(file)
	var dos []Do
	for _, line := range lines {
		var do Do
		if strings.HasPrefix(line, undoneCheckboxStr) {
			do.Done = false
		} else if strings.HasPrefix(strings.ToLower(line), doneCheckboxStr) { // match [x] or [X]
			do.Done = true
		} else {
			// quietly ignore non-do lines
			break
		}

		taskStartIndex := len(doneCheckboxStr) + 1 // add one for space after checkbox
		metaStartIndex := strings.Index(line, " {")
		if metaStartIndex > taskStartIndex {
			var m Metadata
			do.Task = line[taskStartIndex:metaStartIndex]
			err := json.Unmarshal([]byte(line[metaStartIndex:]), &m)
			if err != nil {
				log.Println(err)
			} else {
				do.Metadata = m
			}
		} else {
			do.Task = line[taskStartIndex:]
		}

		dos = append(dos, do)
	}
	return dos
}

func writeDos(file string, dos []Do) {
	// make two slices so that undone items are all first
	var undoneLines []string
	var doneLines []string
	for _, do := range dos {
		metaBytes, _ := json.Marshal(do.Metadata)
		if do.Done {
			doneLines = append(doneLines, fmt.Sprintf("%s %s %s", doneCheckboxStr, do.Task, metaBytes))
		} else {
			undoneLines = append(undoneLines, fmt.Sprintf("%s %s %s", undoneCheckboxStr, do.Task, metaBytes))
		}
	}
	utils.WriteFile(file, append(undoneLines, doneLines...))
}
