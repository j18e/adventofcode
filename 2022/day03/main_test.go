package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_letterScore(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1, letterScore('a'))
	assert.Equal(26, letterScore('z'))
	assert.Equal(27, letterScore('A'))
	assert.Equal(52, letterScore('Z'))
}
