package main

import "testing"

func TestPullDoEmptyFile(t *testing.T) {
	pulled := pullDo([]Do{})
	if pulled.Task != "" {
		t.Errorf("pulled should be empty (%s)", pulled.Task)
	}
}

func TestPullOne(t *testing.T) {
	dos := readDos("./fixtures/fixture-test.md")

	aDo := pullDo(dos)
	if aDo.Task != "undone do" {
		t.Errorf("pulled task should have been 'undone do', not '%s'", aDo.Task)
	}
}

// todo: add some more tests here
