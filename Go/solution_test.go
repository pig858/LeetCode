package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstTest(t *testing.T) {
	type params struct {
		Arg      int
		Expected int
	}

	datas := []params{
		{
			Arg:      1,
			Expected: 1,
		},
		{
			Arg:      2,
			Expected: 3,
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, FirstTest(d.Arg), d)
	}

}

func TestSecondTest(t *testing.T) {
	type params struct {
		Arg      int
		Expected int
	}

	datas := []params{
		{
			Arg:      1,
			Expected: 2,
		},
		{
			Arg:      2,
			Expected: 2,
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, FirstTest(d.Arg), d)
	}

}
