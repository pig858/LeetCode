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

func TestQ2121GetDistances(t *testing.T) {
	type params struct {
		Arg      []int
		Expected []int64
	}

	datas := []params{
		{
			Arg:      StringToSliceType[int](`[2,1,3,1,2,3,3]`),
			Expected: StringToSliceType[int64](`[4,2,7,2,4,4,5]`),
		},
		{
			Arg:      StringToSliceType[int](`[10,5,10,10]`),
			Expected: StringToSliceType[int64](`[5,0,3,4]`),
		},
		{
			Arg:      StringToSliceType[int](`[11266]`),
			Expected: StringToSliceType[int64](`[0]`),
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q2121GetDistances(d.Arg), d)
	}

}

func TestQ1009BitwiseComplement(t *testing.T) {
	type params struct {
		Arg      int
		Expected int
	}

	datas := []params{
		{
			Arg:      5,
			Expected: 2,
		},
		{
			Arg:      7,
			Expected: 0,
		},
		{
			Arg:      10,
			Expected: 5,
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q1009BitwiseComplement(d.Arg), d)
	}

}

func TestQ2206DivideArray(t *testing.T) {
	type params struct {
		Arg      []int
		Expected bool
	}

	datas := []params{
		{
			Arg:      StringToSliceType[int](`[3,2,3,2,2,2]`),
			Expected: true,
		},
		{
			Arg:      StringToSliceType[int](`[1,2,3,4]`),
			Expected: false,
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q2206DivideArray(d.Arg), d)
	}

}

func TestQ150EvalRPN(t *testing.T) {
	type params struct {
		Arg      []string
		Expected int
	}

	datas := []params{
		{
			Arg:      StringToSliceType[string](`["2","1","+","3","*"]`),
			Expected: 9,
		},
		{
			Arg:      StringToSliceType[string](`["4","13","5","/","+"]`),
			Expected: 6,
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q150EvalRPN(d.Arg), d)
	}

}

func TestQ59GenerateMatrix(t *testing.T) {
	type params struct {
		Arg      int
		Expected [][]int
	}

	datas := []params{
		// {
		// 	Arg:      3,
		// 	Expected: StringTo2DSliceType[int](`[[1,2,3],[8,9,4],[7,6,5]]`),
		// },
		// {
		// 	Arg:      1,
		// 	Expected: StringTo2DSliceType[int](`[[1]]`),
		// },
		{
			Arg:      4,
			Expected: StringTo2DSliceType[int](`[[1,2,3,4],[12,13,14,5],[11,16,15,6],[10,9,8,7]]`),
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q59GenerateMatrix(d.Arg), d)
	}

}

func TestQ20IsValid(t *testing.T) {

	type params struct {
		Arg      string
		Expected bool
	}

	datas := []params{
		{
			Arg:      "()",
			Expected: true,
		},
		{
			Arg:      "()[]{}",
			Expected: true,
		},
		{
			Arg:      "(]",
			Expected: false,
		},
		{
			Arg:      "([)]",
			Expected: false,
		},
		{
			Arg:      "{[]}",
			Expected: true,
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q20IsValid(d.Arg), d)
	}
}

func TestQ1290GetDecimalValue(t *testing.T) {

	type params struct {
		Arg      *ListNode
		Expected int
	}

	datas := []params{
		{
			Arg:      StringToListNode(`[1,0,1]`),
			Expected: 5,
		},
		{
			Arg:      StringToListNode(`[0]`),
			Expected: 0,
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q1290GetDecimalValue(d.Arg), d)
	}
}

func TestQ1832CheckIfPangram(t *testing.T) {

	type params struct {
		Arg      string
		Expected bool
	}

	datas := []params{
		{
			Arg:      "thequickbrownfoxjumpsoverthelazydog",
			Expected: true,
		},
		{
			Arg:      "leetcode",
			Expected: false,
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q1832CheckIfPangram(d.Arg), d)
	}
}

func TestQ21MergeTwoLists(t *testing.T) {

	type params struct {
		Arg      *ListNode
		Arg1     *ListNode
		Expected *ListNode
	}

	datas := []params{
		{
			Arg:      StringToListNode(`[1,2,4]`),
			Arg1:     StringToListNode(`[1,3,4]`),
			Expected: StringToListNode(`[1,1,2,3,4,4]`),
		},
		{
			Arg:      StringToListNode(`[]`),
			Arg1:     StringToListNode(`[]`),
			Expected: StringToListNode(`[]`),
		},
		{
			Arg:      StringToListNode(`[]`),
			Arg1:     StringToListNode(`[0]`),
			Expected: StringToListNode(`[0]`),
		},
		{
			Arg:      StringToListNode(`[2]`),
			Arg1:     StringToListNode(`[1]`),
			Expected: StringToListNode(`[1,2]`),
		},
		{
			Arg:      StringToListNode(`[5]`),
			Arg1:     StringToListNode(`[1,2,4]`),
			Expected: StringToListNode(`[1,2,4,5]`),
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q21MergeTwoLists(d.Arg, d.Arg1), d)
	}
}

func TestQ38CountAndSay(t *testing.T) {

	type params struct {
		Arg      int
		Expected string
	}

	datas := []params{
		{
			Arg:      1,
			Expected: "1",
		},
		{
			Arg:      2,
			Expected: "11",
		},
		{
			Arg:      3,
			Expected: "21",
		},
		{
			Arg:      4,
			Expected: "1211",
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q38CountAndSay(d.Arg), d)
	}
}

func TestQ2AddTwoNumbers(t *testing.T) {

	type params struct {
		Arg      *ListNode
		Arg1     *ListNode
		Expected *ListNode
	}

	datas := []params{
		// {
		// 	Arg:      StringToListNode(`[2,4,3]`),
		// 	Arg1:     StringToListNode(`[5,6,4]`),
		// 	Expected: StringToListNode(`[7,0,8]`),
		// },
		// {
		// 	Arg:      StringToListNode(`[0]`),
		// 	Arg1:     StringToListNode(`[0]`),
		// 	Expected: StringToListNode(`[0]`),
		// },
		{
			Arg:      StringToListNode(`[9,9,9,9,9,9,9]`),
			Arg1:     StringToListNode(`[9,9,9,9]`),
			Expected: StringToListNode(`[8,9,9,9,0,0,0,1]`),
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q2AddTwoNumbers(d.Arg, d.Arg1), d)
	}
}

func TestQ692TopKFrequent(t *testing.T) {

	type params struct {
		Arg      []string
		Arg1     int
		Expected []string
	}

	datas := []params{
		{
			Arg:      StringToSliceType[string](`["i","love","leetcode","i","love","coding"]`),
			Arg1:     2,
			Expected: StringToSliceType[string](`["i", "love"]`),
		},
		{
			Arg:      StringToSliceType[string](`["the","day","is","sunny","the","the","the","sunny","is","is"]`),
			Arg1:     4,
			Expected: StringToSliceType[string](`["the","is","sunny","day"]`),
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q692TopKFrequent(d.Arg, d.Arg1), d)
	}
}

func TestQ12IntToRoman(t *testing.T) {
	type params struct {
		Arg      int
		Expected string
	}

	datas := []params{
		{
			Arg:      3,
			Expected: "III",
		},
		{
			Arg:      58,
			Expected: "LVIII",
		},
		{
			Arg:      1994,
			Expected: "MCMXCIV",
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q12IntToRoman(d.Arg), d)
	}
}

func TestQ219ContainsNearbyBuplicate(t *testing.T) {
	type params struct {
		Arg      []int
		Arg1     int
		Expected bool
	}

	datas := []params{
		{
			Arg:      StringToSliceType[int](`[1,2,3,1]`),
			Arg1:     3,
			Expected: true,
		},
		{
			Arg:      StringToSliceType[int](`[1,0,1,1]`),
			Arg1:     1,
			Expected: true,
		},
		{
			Arg:      StringToSliceType[int](`[1,2,3,1,2,3]`),
			Arg1:     2,
			Expected: false,
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q219ContainsNearbyDuplicate(d.Arg, d.Arg1), d)
	}
}

func TestQ76MinWindow(t *testing.T) {
	type params struct {
		Arg      string
		Arg1     string
		Expected string
	}

	datas := []params{
		{
			Arg:      "ADOBECODEBANC",
			Arg1:     "ABC",
			Expected: "BANC",
		},
		{
			Arg:      "a",
			Arg1:     "a",
			Expected: "a",
		},
		{
			Arg:      "a",
			Arg1:     "aa",
			Expected: "",
		},
		{
			Arg:      "ab",
			Arg1:     "a",
			Expected: "a",
		},
		{
			Arg:      "ab",
			Arg1:     "b",
			Expected: "b",
		},
		{
			Arg:      "abc",
			Arg1:     "b",
			Expected: "b",
		},
		{
			Arg:      "bba",
			Arg1:     "ab",
			Expected: "ba",
		},
		{
			Arg:      "bbaa",
			Arg1:     "aba",
			Expected: "baa",
		},
		{
			Arg:      "abcabdebac",
			Arg1:     "cea",
			Expected: "ebac",
		},
		{
			Arg:      "aBbaBBBBA",
			Arg1:     "BBA",
			Expected: "BBA",
		},
	}

	for _, d := range datas {
		assert.Equal(t, d.Expected, Q76MinWindow(d.Arg, d.Arg1), d)
	}
}
