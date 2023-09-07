package main

import "testing"

func TestReadDosSuccess(t *testing.T) {
	dos := readDos("./fixtures/fixture-test.md")
	if len(dos) != 3 {
		t.Errorf("there should be 3 dos (%d)", len(dos))
	}
}

func TestReadDosEmptyFile(t *testing.T) {
	dos := readDos("./fixtures/fixture-empty.md")
	if len(dos) > 0 {
		t.Errorf("there should be 0 dos (%d)", len(dos))
	}
}
