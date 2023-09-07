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

func TestReplaceEmptyList(t *testing.T) {
	dos := []Do{}

	aDo := replaceLastPulled(dos)
	if aDo != (Do{}) {
		t.Errorf("unpulled task should be empty, not '%s'", aDo.Task)
	}
}

func TestReplace(t *testing.T) {
	dos := readDos("./fixtures/fixture-test.md")

	// first replace: replace "DONE DO"
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

	// second replace: replace "done do"
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

	// third replace: nothing left to replace
	replacedDo = replaceLastPulled(dos)
	if err := verifyReplacedDo(replacedDo, ""); err != "" {
		t.Errorf(err)
	}
}

func TestReplaceWithCorruptDoBogusTimestamp(t *testing.T) {
	dos := readDos("./fixtures/fixture-test.md")

	if dos[1].Done {
		t.Errorf("this test Do isn't supposed to be done yet")
	}
	dos[1].Metadata.PulledDate = time.Now() // add timestamp even though this do isn't Done

	// first replace: replace "DONE DO" (not "undone do", despite the latter having a newer timestamp we just set)
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
}

func TestReplaceWithCorruptDoMissingTimestamp(t *testing.T) {
	dos := readDos("./fixtures/fixture-test.md")

	if !dos[2].Done {
		t.Errorf("this test Do is supposed to be done")
	}
	dos[2].Metadata.PulledDate = time.Time{} // zero out timestamp even though this do is Done

	// first replace: replace "done do" (not "DONE DO", which has a zero'd timestamp)
	replacedDo := replaceLastPulled(dos)
	if err := verifyReplacedDo(replacedDo, "done do"); err != "" {
		t.Errorf(err)
	}

	// second replace: now we get "DONE DO", since there aren't any newer ones than time zero
	replacedDo = replaceLastPulled(dos)
	if err := verifyReplacedDo(replacedDo, "DONE DO"); err != "" {
		t.Errorf(err)
	}
}
