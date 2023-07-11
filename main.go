package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golang.org/x/exp/slices"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
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

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func readFile(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFile(f)
	var lines []string
	input := bufio.NewScanner(f)
	for input.Scan() {
		lines = append(lines, input.Text())
	}

	return lines
}

func writeFile(file string, lines []string) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFile(f)
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
	writeFile(file, append(undoneLines, doneLines...))
}

func replaceLastPulled(dos []Do) Do {
	replacedIndex := -1
	latest := Do{} // initialize to empty Do
	if len(dos) > 0 {
		// find the latest pulled-date
		for i, do := range dos {
			if do.Metadata.PulledDate.After(latest.Metadata.PulledDate) {
				replacedIndex = i
				latest = dos[replacedIndex]
			}
		}
		if replacedIndex >= 0 {
			fmt.Printf("replaced '%s'\n", latest.Task)

			// have to make updates directly to the slice element (not the reference, 'latest')
			dos[replacedIndex].Done = false
			dos[replacedIndex].Metadata.PulledDate = time.Time{}

			return dos[replacedIndex]
		} else {
			fmt.Println("no done tasks found")
		}
	}
	return latest
}

func pullDo(dos []Do) (aDo Do) {
	undoneCount := 0
	for _, do := range dos {
		if !do.Done {
			undoneCount++
		}
	}

	if undoneCount > 0 {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(undoneCount)

		for i, _ := range dos {
			if i == r {
				fmt.Println(dos[i].Task)
				dos[i].Done = true
				dos[i].Metadata.PulledDate = time.Now()
				aDo = dos[i]
				break
			}
		}
	} else {
		fmt.Println("[no undone tasks found!]")
	}
	return
}

// possibilities:
// honey-do
// honey-do pull
// honey-do filename
// honey-do pull filename
// honey-do add "task"
// honey-do add "task" filename
// etc.
func parseCommandLine(args []string) (action string, filename string, newTask string) {
	validActions := []string{"pull", "add", "unpull", "swap"}

	for i := 1; i < len(args); i++ {
		currentArg := args[i]
		if slices.Contains(validActions, currentArg) {
			action = currentArg

			if action == "add" {
				i++
				newTask = args[i]
			}
		} else {
			// must be a filename
			filename = currentArg
		}
	}

	// apply defaults if needed
	if action == "" {
		action = "pull"
	}
	if filename == "" {
		filename = os.Getenv("HONEY_DO_FILE")
	}
	return
}

func act(action string, dos []Do, task string) (updatedDos []Do) {
	switch action {
	case "pull":
		aDo := pullDo(dos)
		fmt.Println("[ado is: " + aDo.Task + "]")
	case "add":
		newDo := Do{Done: false, Task: task, Metadata: Metadata{AddedDate: time.Now()}}
		dos = append(dos, newDo)
	case "unpull":
		aDo := replaceLastPulled(dos)
		fmt.Println("[ado is: " + aDo.Task + "]")
	case "swap":
		aDo := replaceLastPulled(dos)
		if aDo.Task != "" {
			for {
				newDo := pullDo(dos)
				if newDo.Task != aDo.Task {
					break
				}
			}
		}
	default:
		fmt.Printf("unrecognized action '%s'\n", action)
		os.Exit(2)
	}

	return dos
}

func main() {
	action, filename, task := parseCommandLine(os.Args)

	// todo: bonk on empty filename
	dos := readDos(filename)

	// todo: put out message & skip the rest here on no dos
	updatedDos := act(action, dos, task)

	if len(updatedDos) > 0 {
		writeDos(filename, dos)
	}
}
