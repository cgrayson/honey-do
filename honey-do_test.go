package main

import (
	"os"
	"testing"
)

// 0th arg is the executable name
var baseArgs = []string{"honey-do"}

const filename1 = "/tmp/honey-do.md"
const filename2 = "this-here-honey-do.md"

func TestParseCommandLine(t *testing.T) {

	var tests = []struct {
		setEnvFilename   bool
		params           []string
		expectedAction   string
		expectedFilename string
		expectedTask     string
	}{
		{false, nil, "pull", "", ""},
		{true, nil, "pull", filename1, ""},

		{false, []string{"pull"}, "pull", "", ""},
		{true, []string{"pull"}, "pull", filename1, ""},

		{false, []string{filename2}, "pull", filename2, ""},
		{true, []string{filename2}, "pull", filename2, ""},

		{false, []string{"pull", filename2}, "pull", filename2, ""},
		{true, []string{"pull", filename2}, "pull", filename2, ""},

		{false, []string{"add", "Do this task"}, "add", "", "Do this task"},
		{true, []string{"add", "Do this task"}, "add", filename1, "Do this task"},

		{false, []string{"add", ""}, "add", "", ""},
		{false, []string{"add"}, "add", "", ""},

		{false, []string{"add", "Do this task", filename2}, "add", filename2, "Do this task"},
		{true, []string{"add", "Do this task", filename2}, "add", filename2, "Do this task"},

		{false, []string{"foo"}, "pull", "foo", ""},
	}

	for i, test := range tests {
		if test.setEnvFilename {
			_ = os.Setenv("HONEY_DO_FILE", filename1)
		} else {
			_ = os.Unsetenv("HONEY_DO_FILE")
		}

		action, filename, taskText := parseCommandLine(append(baseArgs, test.params...))

		if action != test.expectedAction {
			t.Errorf("%d: action should be %q, not %q", i, test.expectedAction, action)
		}
		if filename != test.expectedFilename {
			t.Errorf("%d: filename should be %q, not %q", i, test.expectedFilename, filename)
		}
		if taskText != test.expectedTask {
			t.Errorf("%d: task text should be %q, not %s", i, test.expectedTask, taskText)
		}
	}
}

func TestHelpFlags(t *testing.T) {
	helpFlags := []string{"--help", "-?", "help", "?", "-h", "h", "--HELP", "HELP", "-H", "H"}
	for _, flag := range helpFlags {
		action, _, _ := parseCommandLine(append(baseArgs, flag))
		if action != "help" {
			t.Errorf("action should be help for flag %q (not %s)", flag, action)
		}
	}
}

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
