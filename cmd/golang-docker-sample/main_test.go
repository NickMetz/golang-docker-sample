package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RandomNumber(t *testing.T) {
	n := RandomNumber()
	assert.IsType(t, 0, n)
}
