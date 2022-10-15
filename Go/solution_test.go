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

func TestQ14LongestCommonPrefix(t *testing.T) {
	type params struct {
		Arg      []string
		Expected string
	}

	datas := []params{
		{
			Arg:      StringToSliceType[string](`["flower", "flow", "flight"]`),
			Expected: "fl",
		},
		{
			Arg:      StringToSliceType[string](`["dog", "racecar", "car"]`),
			Expected: "",
		},
		{
			Arg:      StringToSliceType[string](`["ab", "a"]`),
			Expected: "a",
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q14LongestCommonPrefix(d.Arg), d)
	}

}

func TestQ1108DefangingAnIPAddress(t *testing.T) {
	type params struct {
		Arg      string
		Expected string
	}

	datas := []params{
		{
			Arg:      "1.1.1.1",
			Expected: "1[.]1[.]1[.]1",
		},
		{
			Arg:      "255.100.50.0",
			Expected: "255[.]100[.]50[.]0",
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q1108DefangingAnIPAddress(d.Arg), d)
	}

}

func TestQ771JewelsAndStones(t *testing.T) {
	type params struct {
		Arg      string
		Arg2     string
		Expected int
	}

	datas := []params{
		{
			Arg:      "aA",
			Arg2:     "aAAbbbb",
			Expected: 3,
		},
		{
			Arg:      "z",
			Arg2:     "ZZ",
			Expected: 0,
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q771JewelsAndStones(d.Arg, d.Arg2), d)
	}

}

func TestQ938RangeSumOfBST(t *testing.T) {
	type params struct {
		Arg      *TreeNode
		Arg2     int
		Arg3     int
		Expected int
	}

	datas := []params{
		{
			Arg:      StringToTreeNode(`[10,5,15,3,7,null,18]`),
			Arg2:     7,
			Arg3:     15,
			Expected: 32,
		},
		{
			Arg:      StringToTreeNode(`[10,5,15,3,7,13,18,1,null,6]`),
			Arg2:     6,
			Arg3:     10,
			Expected: 23,
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q938RangeSumOfBST(d.Arg, d.Arg2, d.Arg3), d)
	}

}
