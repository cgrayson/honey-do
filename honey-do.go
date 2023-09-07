package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"math/rand"
	"os"
	"strings"
	"time"
)

func replaceLastPulled(dos []Do) Do {
	replacedIndex := -1
	latest := Do{} // initialize to empty Do
	if len(dos) > 0 {
		// find the latest pulled-date
		for i, do := range dos {
			// PulledDate shouldn't be set at all if Done isn't true, but just in case
			if do.Done && (do.Metadata.PulledDate.After(latest.Metadata.PulledDate) || do.Metadata.PulledDate.Equal(latest.Metadata.PulledDate)) {
				replacedIndex = i
				latest = dos[replacedIndex]
			}
		}
		if replacedIndex >= 0 {
			// have to make updates directly to the slice element (not the reference, 'latest')
			dos[replacedIndex].Done = false
			dos[replacedIndex].Metadata.PulledDate = time.Time{}

			return dos[replacedIndex]
		}
	}
	return latest
}

func pullDo(dos []Do) (aDo Do) {
	numUndone := CountUndone(dos)

	if numUndone > 0 {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(numUndone)
		undoneIndex := -1

		var i int
		for i, aDo = range dos {
			if !aDo.Done {
				undoneIndex++
			}
			// check the random # against how many undones we've gone through
			if undoneIndex == r {
				// use the true index to mark done & set date
				dos[i].Done = true
				dos[i].Metadata.PulledDate = time.Now()
				break
			}
		}
	}
	return aDo
}

func helpRequested(args []string) bool {
	if len(args) > 1 {
		switch strings.ToLower(args[1]) {
		case "--help", "-?", "help", "?", "-h", "h":
			return true
		}
	}
	return false
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
	// check for help first
	if helpRequested(args) {
		return "help", "", ""
	}
	validActions := []string{"pull", "add", "unpull", "swap"}

	for i := 1; i < len(args); i++ { // start at 1 to skip executable name
		currentArg := args[i]
		if slices.Contains(validActions, currentArg) {
			action = currentArg

			if action == "add" {
				i++
				if len(args) > i { // make sure there's another arg to get
					newTask = args[i]
				} else {
					newTask = "" // empty task will be rejected later
				}
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

func act(action string, dos []Do, task string) ([]Do, string) {
	var message string
	numUndone := CountUndone(dos)

	switch action {
	case "pull":
		if numUndone >= 1 {
			aDo := pullDo(dos)
			if aDo != (Do{}) {
				message = fmt.Sprintf("your task is: %s", aDo.Task)
			}
		} else {
			message = fmt.Sprintf("[no undone tasks found!]")
		}
	case "add":
		task := strings.TrimSpace(task)
		if task != "" {
			newDo := Do{Done: false, Task: task, Metadata: Metadata{AddedDate: time.Now()}}
			message = fmt.Sprintf("added task: %s", newDo.Task)
			dos = append(dos, newDo)
		} else {
			message = "[no task to add]"
		}
	case "unpull":
		aDo := replaceLastPulled(dos)
		if aDo != (Do{}) {
			message = fmt.Sprintf("returned task: %s", aDo.Task)
		} else {
			message = fmt.Sprintf("[no tasks to return]")
		}
	case "swap":
		if numUndone >= 1 { // have to have at least 1 undone to swap
			aDo := replaceLastPulled(dos)
			if aDo != (Do{}) {
				message = fmt.Sprintf("returned task: %s", aDo.Task)
				for {
					newDo := pullDo(dos)
					if newDo.Task == aDo.Task {
						// put it back again
						replaceLastPulled(dos)
					} else {
						message += fmt.Sprintf("\nyour new task is: %s", newDo.Task)
						break
					}
				}
			} else {
				message = fmt.Sprintf("[no tasks to return]")
			}
		} else {
			message = fmt.Sprintf("[no undone tasks to swap for!]")
		}
	default:
		message = fmt.Sprintf("[oops: unrecognized action '%s']", action)
	}

	return dos, message
}

func main() {
	action, filename, task := parseCommandLine(os.Args)

	if action == "help" {
		fmt.Println("honey-do [ pull | unpull | swap | add 'task' ] filename.md (or $HONEY_DO_FILE)")
		return
	}

	if filename == "" {
		fmt.Println("[oops: you have to specify a honey-do file to work with]")
	} else {
		dos := readDos(filename)
		if len(dos) == 0 && action != "add" {
			fmt.Printf("[oops: no to-dos found in the file '%s']\n", filename)
		} else {
			updatedDos, message := act(action, dos, task)
			fmt.Println(message)

			if len(updatedDos) > 0 {
				writeDos(filename, updatedDos)
			}
		}
	}
}
