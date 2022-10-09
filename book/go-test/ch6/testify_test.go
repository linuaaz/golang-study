package ch6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Object struct {
	s string
}

func TestSomething(t *testing.T) {
	assert.Equal(t, 123, 123, "they should be equal")
	assert.NotEqual(t, 123, 456, "they should not be equal")
	assert.Nil(t, nil)
	o := Object{s: "something"}
	if assert.NotNil(t, o) {
		assert.Equal(t, "something", o.s)
	}
}
