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

	// first pull: replace "DONE DO"
	replacedDo := replaceLastPulled(dos)
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

	// second pull: replace "done do"
	replacedDo = replaceLastPulled(dos)
	if err := verifyReplacedDo(replacedDo, "done do"); err != "" {
		t.Errorf(err)
	}
	for _, do := range dos {
		if do.Task == "done do" {
			sliceDo = do
			break
		}
	}
	if err := verifyReplacedDo(sliceDo, "done do"); err != "" {
		t.Errorf(err)
	}

	// third pull: nothing left to replace
	replacedDo = replaceLastPulled(dos)
	if err := verifyReplacedDo(replacedDo, ""); err != "" {
		t.Errorf(err)
	}
}
