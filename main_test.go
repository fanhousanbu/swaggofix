package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReg(t *testing.T) {

	t1 := "a/:abc/def/"
	r1 := repl(&t1)
	assert.Equal(t, "a/{abc}/def/", *r1)

	t2 := "/:abc/def/"
	r2 := repl(&t2)
	assert.Equal(t, "/{abc}/def/", *r2)

	t3 := "/:abc"
	r3 := repl(&t3)
	assert.Equal(t, "/{abc}", *r3)

	t4 := ":abc"
	r4 := repl(&t4)
	assert.Equal(t, ":abc", *r4)
}
