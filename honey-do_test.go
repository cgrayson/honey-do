package main

import (
	"os"
	"testing"
)

// 0th arg is the executable name
var baseArgs = []string{"honey-do"}

// honey-do # no env
func TestNoParametersNoEnv(t *testing.T) {
	_ = os.Unsetenv("HONEY_DO_FILE")
	action, filename, taskText := parseCommandLine(baseArgs)
	if action != "pull" {
		t.Errorf("action should be pull (%s)", action)
	}
	if filename != "" {
		t.Errorf("filename should be empty (%s)", filename)
	}
	if taskText != "" {
		t.Errorf("task text should be empty (%s)", taskText)
	}
}

// honey-do # with env
func TestNoParametersWithEnv(t *testing.T) {
	_ = os.Setenv("HONEY_DO_FILE", "/tmp/honey-do.md")
	action, filename, taskText := parseCommandLine(baseArgs)
	if action != "pull" {
		t.Errorf("action should be pull (%s)", action)
	}
	if filename != "/tmp/honey-do.md" {
		t.Errorf("filename should be set (%s)", filename)
	}
	if taskText != "" {
		t.Errorf("task text should be empty (%s)", taskText)
	}
}

// honey-do pull # no env
func TestActionOnlyNoEnv(t *testing.T) {
	_ = os.Unsetenv("HONEY_DO_FILE")
	action, filename, taskText := parseCommandLine(append(baseArgs, "pull"))
	if action != "pull" {
		t.Errorf("action should be pull (%s)", action)
	}
	if filename != "" {
		t.Errorf("filename should be empty (%s)", filename)
	}
	if taskText != "" {
		t.Errorf("task text should be empty (%s)", taskText)
	}
}

// honey-do pull # with env
func TestActionOnlyWithEnv(t *testing.T) {
	_ = os.Setenv("HONEY_DO_FILE", "/tmp/honey-do.md")
	action, filename, taskText := parseCommandLine(append(baseArgs, "pull"))
	if action != "pull" {
		t.Errorf("action should be pull (%s)", action)
	}
	if filename != "/tmp/honey-do.md" {
		t.Errorf("filename should be set (%s)", filename)
	}
	if taskText != "" {
		t.Errorf("task text should be empty (%s)", taskText)
	}
}

// honey-do filename # no env
func TestFilenameOnlyNoEnv(t *testing.T) {
	_ = os.Unsetenv("HONEY_DO_FILE")
	action, filename, taskText := parseCommandLine(append(baseArgs, "pull", "this-here-honey-do.md"))
	if action != "pull" {
		t.Errorf("action should be pull (%s)", action)
	}
	if filename != "this-here-honey-do.md" {
		t.Errorf("filename should be set (%s)", filename)
	}
	if taskText != "" {
		t.Errorf("task text should be empty (%s)", taskText)
	}
}

// honey-do filename # with env
func TestFilenameOnlyWithEnv(t *testing.T) {
	_ = os.Setenv("HONEY_DO_FILE", "/tmp/honey-do.md")
	action, filename, taskText := parseCommandLine(append(baseArgs, "this-here-honey-do.md"))
	if action != "pull" {
		t.Errorf("action should be pull (%s)", action)
	}
	if filename != "this-here-honey-do.md" {
		t.Errorf("filename should be set (%s)", filename)
	}
	if taskText != "" {
		t.Errorf("task text should be empty (%s)", taskText)
	}

}

// honey-do pull filename # no env
func TestActionAndFilenameNoEnv(t *testing.T) {
	_ = os.Unsetenv("HONEY_DO_FILE")
	action, filename, taskText := parseCommandLine(append(baseArgs, "pull", "this-here-honey-do.md"))
	if action != "pull" {
		t.Errorf("action should be pull (%s)", action)
	}
	if filename != "this-here-honey-do.md" {
		t.Errorf("filename should be set (%s)", filename)
	}
	if taskText != "" {
		t.Errorf("task text should be empty (%s)", taskText)
	}

}

// honey-do pull filename # with env
func TestActionAndFilenameWithEnv(t *testing.T) {
	_ = os.Setenv("HONEY_DO_FILE", "/tmp/honey-do.md")
	action, filename, taskText := parseCommandLine(append(baseArgs, "pull", "this-here-honey-do.md"))
	if action != "pull" {
		t.Errorf("action should be pull (%s)", action)
	}
	if filename != "this-here-honey-do.md" {
		t.Errorf("filename should be set (%s)", filename)
	}
	if taskText != "" {
		t.Errorf("task text should be empty (%s)", taskText)
	}

}

// honey-do add "task" # no env
func TestAddOnlyNoEnv(t *testing.T) {
	_ = os.Unsetenv("HONEY_DO_FILE")
	action, filename, taskText := parseCommandLine(append(baseArgs, "add", "Do this task"))
	if action != "add" {
		t.Errorf("action should be add (%s)", action)
	}
	if filename != "" {
		t.Errorf("filename should be empty (%s)", filename)
	}
	if taskText != "Do this task" {
		t.Errorf("task text should be set (%s)", taskText)
	}

}

func TestAddWithEmptyTask(t *testing.T) {
	_ = os.Unsetenv("HONEY_DO_FILE")
	action, filename, taskText := parseCommandLine(append(baseArgs, "add", ""))
	if action != "add" {
		t.Errorf("action should be add (%s)", action)
	}
	if filename != "" {
		t.Errorf("filename should be empty (%s)", filename)
	}
	if taskText != "" {
		t.Errorf("task text should be empty (%s)", taskText)
	}
}

func TestAddWithNoTask(t *testing.T) {
	_ = os.Unsetenv("HONEY_DO_FILE")
	action, filename, taskText := parseCommandLine(append(baseArgs, "add"))
	if action != "add" {
		t.Errorf("action should be add (%s)", action)
	}
	if filename != "" {
		t.Errorf("filename should be empty (%s)", filename)
	}
	if taskText != "" {
		t.Errorf("task text should be empty (%s)", taskText)
	}
}

// honey-do add "task" # with env
func TestAddOnlyWithEnv(t *testing.T) {
	_ = os.Setenv("HONEY_DO_FILE", "/tmp/honey-do.md")
	action, filename, taskText := parseCommandLine(append(baseArgs, "add", "Do this task"))
	if action != "add" {
		t.Errorf("action should be add (%s)", action)
	}
	if filename != "/tmp/honey-do.md" {
		t.Errorf("filename should be set (%s)", filename)
	}
	if taskText != "Do this task" {
		t.Errorf("task text should be set (%s)", taskText)
	}
}

// honey-do add "task" filename # no env
func TestAddAndFilenameNoEnv(t *testing.T) {
	_ = os.Unsetenv("HONEY_DO_FILE")
	action, filename, taskText := parseCommandLine(append(baseArgs, "add", "Do this task", "this-here-honey-do.md"))
	if action != "add" {
		t.Errorf("action should be add (%s)", action)
	}
	if filename != "this-here-honey-do.md" {
		t.Errorf("filename should be set (%s)", filename)
	}
	if taskText != "Do this task" {
		t.Errorf("task text should be set (%s)", taskText)
	}
}

// honey-do add "task" filename # with env
func TestAddAndFilenameWithEnv(t *testing.T) {
	_ = os.Setenv("HONEY_DO_FILE", "/tmp/honey-do.md")
	action, filename, taskText := parseCommandLine(append(baseArgs, "add", "Do this task", "this-here-honey-do.md"))
	if action != "add" {
		t.Errorf("action should be add (%s)", action)
	}
	if filename != "this-here-honey-do.md" {
		t.Errorf("filename should be set (%s)", filename)
	}
	if taskText != "Do this task" {
		t.Errorf("task text should be set (%s)", taskText)
	}
}

// honey-do foo
func TestInvalidActionOnly(t *testing.T) {
	_ = os.Unsetenv("HONEY_DO_FILE")
	action, filename, taskText := parseCommandLine(append(baseArgs, "foo"))
	if action != "pull" {
		t.Errorf("action should be pull (%s)", action)
	}
	if filename != "foo" {
		t.Errorf("filename should be foo (%s)", filename)
	}
	if taskText != "" {
		t.Errorf("task text should be empty (%s)", taskText)
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
