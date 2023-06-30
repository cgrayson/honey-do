package main

import (
	"testing"
	"time"
)

func TestReplace(t *testing.T) {
	dos := readDos("./fixtures/fixture-test.md")
	replacedDo := replaceLastPulled(dos)

	// test the return value
	if replacedDo.Task != "DONE DO" {
		t.Errorf("replaced task should be 'DONE DO', not '%s'", replacedDo.Task)
	}
	if replacedDo.Done != false {
		t.Errorf("replaced task should be marked not Done")
	}
	nilTime := time.Time{}
	if replacedDo.Metadata.PulledDate != nilTime {
		t.Errorf("replaced task's pulled date should be nil value, not '%v'", replacedDo.Metadata.PulledDate)
	}

	var sliceDo Do
	// test the entry in the slice
	for _, do := range dos {
		if do.Task == "DONE DO" {
			sliceDo = do
			break
		}
	}
	if sliceDo.Done != false {
		t.Errorf("replaced task should be marked not Done")
	}
	if sliceDo.Metadata.PulledDate != nilTime {
		t.Errorf("replaced task's pulled date should be nil value, not '%v'", sliceDo.Metadata.PulledDate)
	}

	// todo next time: more tests:
	// what if no replaced found?
}
