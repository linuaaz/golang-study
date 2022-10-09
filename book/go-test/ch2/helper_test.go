package ch2

import "testing"

// if an error occurs in the helper function, we use the t.Helper()
// then the error code line message will be displayed directly to the location where the helper function was called,
// not the line of code within the helper function where the error occurred

func notHelper(t *testing.T, msg string) {
	t.Error(msg)
}

func helper(t *testing.T, msg string) {
	t.Helper()
	t.Errorf(msg)
}

func notHelperCallingHelper(t *testing.T, msg string) {
	helper(t, msg)
}

func helperCallingHelper(t *testing.T, msg string) {
	t.Helper()
	helper(t, msg)
}

func TestHelper(t *testing.T) {
	notHelper(t, "0")              // not use t.Helper(), error line: 10
	helper(t, "1")                 // use t.Helper(), error line: 29
	notHelperCallingHelper(t, "2") // use func(not use t.Helper()) to call another func(use t.Helper()), error line: 19
	helperCallingHelper(t, "3")    // use func(use t.Helper()) to call another func(use t.Helper()), error line: 31

	fn := func(msg string) {
		t.Helper()
		t.Errorf(msg)
	}
	fn("4") // use t.Helper(), error line: 37

	t.Helper()
	t.Error("5")

	t.Run("sub", func(t *testing.T) {
		helper(t, "6")                 // use t.Helper(), error line: 43
		notHelperCallingHelper(t, "7") // use func(not use t.Helper()) to call another func(use t.Helper()), error line: 19
		t.Helper()
		t.Errorf("8") // use t.Helper(), error line: 42
	})

}
