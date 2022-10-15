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

func TestQ1281SubtractTheProductAndSumOfDigitsOfAnInteger(t *testing.T) {
	type params struct {
		Arg      int
		Expected int
	}

	datas := []params{
		{
			Arg:      234,
			Expected: 15,
		},
		{
			Arg:      4421,
			Expected: 21,
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q1281SubtractTheProductAndSumOfDigitsOfAnInteger(d.Arg), d)
	}

}

func TestQ1295FindNumbers(t *testing.T) {
	type params struct {
		Arg      []int
		Expected int
	}

	datas := []params{
		{
			Arg:      StringToSliceType[int](`[12,345,2,6,7896]`),
			Expected: 2,
		},
		{
			Arg:      StringToSliceType[int](`[555,901,482,1771]`),
			Expected: 1,
		},
		{
			Arg:      StringToSliceType[int](`[437,315,322,431,686,264,442]`),
			Expected: 0,
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q1295FindNumbers(d.Arg), d)
	}

}

func TestQ643FindMaxAverage(t *testing.T) {
	type params struct {
		Arg      []int
		Arg1     int
		Expected float64
	}

	datas := []params{
		{
			Arg:      StringToSliceType[int](`[1,12,-5,-6,50,3]`),
			Arg1:     4,
			Expected: 12.75000,
		},
		{
			Arg:      StringToSliceType[int](`[5]`),
			Arg1:     1,
			Expected: 5.00000,
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q643FindMaxAverage(d.Arg, d.Arg1), d)
	}

}
func TestQ697FindShortestSubArray(t *testing.T) {
	type params struct {
		Arg      []int
		Expected int
	}

	datas := []params{
		{
			Arg:      StringToSliceType[int](`[1,2,2,3,1]`),
			Expected: 2,
		},
		{
			Arg:      StringToSliceType[int](`[1,2,2,3,1,4,2]`),
			Expected: 6,
		},
		{
			Arg:      StringToSliceType[int](`[1, 5, 2, 1, 5, 4, 5, 1]`),
			Expected: 6,
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q697FindShortestSubArray(d.Arg), d)
	}

}
