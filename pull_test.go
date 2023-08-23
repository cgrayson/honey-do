package main

import "testing"

func TestPullEmptyList(t *testing.T) {
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

func TestPullLast(t *testing.T) {
	dos := readDos("./fixtures/fixture-test.md")

	aDo := pullDo(dos)
	if aDo.Task != "undone do" {
		t.Errorf("pulled task should have been 'undone do', not '%s'", aDo.Task)
	}

	aDo = pullDo(dos)
	if aDo.Task != "" {
		t.Errorf("pulled task should have been '' (empty string), not '%s'", aDo.Task)
	}
}
