package main

import (
	"fmt"
	"testing"
	"time"
)

func verifyReplacedDo(replacedDo Do, taskText string) string {
	// test the return value
	if replacedDo.Task != taskText {
		return fmt.Sprintf("replaced task should be '%s', not '%s'", taskText, replacedDo.Task)
	}
	if replacedDo.Done != false {
		return fmt.Sprintf("replaced task should be marked not Done")
	}
	nilTime := time.Time{}
	if replacedDo.Metadata.PulledDate != nilTime {
		return fmt.Sprintf("replaced task's pulled date should be nil value, not '%v'", replacedDo.Metadata.PulledDate)
	}
	return ""
}

func TestReplace(t *testing.T) {
	dos := readDos("./fixtures/fixture-test.md")
	replacedDo := replaceLastPulled(dos)

	// test all aspects of the return value
	if err := verifyReplacedDo(replacedDo, "DONE DO"); err != "" {
		t.Errorf(err)
	}

	var sliceDo Do
	// test the entry in the slice
	for _, do := range dos {
		if do.Task == "DONE DO" {
			sliceDo = do
			break
		}
	}

	if err := verifyReplacedDo(sliceDo, "DONE DO"); err != "" {
		t.Errorf(err)
	}

	// todo next time: more tests:
	// what if no replaced found?
}
