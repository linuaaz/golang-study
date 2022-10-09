package ch1

import "testing"

func TestGreet(t *testing.T) {
	s := greet("test")

	if s != "hello test" {
		t.Errorf("want 'hello test' but got %s", s)
	}
}

// Does not necessarily need to be capitalized
func Test_great(t *testing.T) {
	s := greet("test")

	if s != "hello test" {
		t.Errorf("want 'hello test' but got %s", s)
	}
}
