package main

import "testing"

func verifyMessage(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Errorf("message should be '%s' (not '%s')", expected, actual)
	}
}

func TestPullAction(t *testing.T) {
	dos := readDos("./fixtures/fixture-test.md")

	_, message := act("pull", dos, "")
	verifyMessage(t, message, "your task is: undone do")

	// pull again, there shouldn't be any left
	_, message = act("pull", dos, "")
	verifyMessage(t, message, "[no undone tasks found!]")
}

func TestPullActionEmptyList(t *testing.T) {
	var dos []Do

	_, message := act("pull", dos, "")
	verifyMessage(t, message, "[no undone tasks found!]")
}

func TestAddAction(t *testing.T) {
	dos := readDos("./fixtures/fixture-test.md")

	_, message := act("add", dos, "Do this other thing")
	verifyMessage(t, message, "added task: Do this other thing")
}

func TestAddActionEmptyTask(t *testing.T) {
	dos := readDos("./fixtures/fixture-test.md")

	_, message := act("add", dos, "")
	verifyMessage(t, message, "[no task to add]")
}

func TestUnpullAction(t *testing.T) {
	dos := readDos("./fixtures/fixture-test.md")

	_, message := act("unpull", dos, "")
	verifyMessage(t, message, "returned task: DONE DO")

	// unpull again
	_, message = act("unpull", dos, "")
	verifyMessage(t, message, "returned task: done do")

	// once more, shouldn't be any left
	_, message = act("unpull", dos, "")
	verifyMessage(t, message, "[no tasks to return]")
}

func TestUnpullActionEmptyList(t *testing.T) {
	var dos []Do

	_, message := act("unpull", dos, "")
	verifyMessage(t, message, "[no tasks to return]")
}

func TestSwapAction(t *testing.T) {
	dos := readDos("./fixtures/fixture-test.md")

	// first swap
	_, message := act("swap", dos, "")
	verifyMessage(t, message, "returned task: DONE DO\nyour new task is: undone do")

	// then pull to get the last one
	_, message = act("pull", dos, "")
	verifyMessage(t, message, "your task is: DONE DO")

	// then there are none left to swap
	_, message = act("swap", dos, "")
	verifyMessage(t, message, "[no undone tasks to swap for!]")
}

func TestSwapActionEmptyList(t *testing.T) {
	var dos []Do

	_, message := act("swap", dos, "")
	verifyMessage(t, message, "[no undone tasks to swap for!]")
}

func TestBadAction(t *testing.T) {
	dos := readDos("./fixtures/fixture-test.md")

	_, message := act("foobar", dos, "")
	verifyMessage(t, message, "[oops: unrecognized action 'foobar']")
}
