package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("hello, world")
}

func FirstTest(i int) int {
	return i
}

func Q14LongestCommonPrefix(strs []string) string {
	// sort strs ensure use the shortest string to compare
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	s := (strs[0])

	l := len(s)
	newStrs := strs[1:]
	ans := ""

	for i := 0; i < l; i++ {
		pre := s[i]
		for _, v := range newStrs {
			if v[i] != pre {
				return ans
			}
		}
		ans += string(pre)
	}

	return ans
}

func Q1108DefangingAnIPAddress(address string) string {
	ans := ""
	arr := []byte(address)

	for _, v := range arr {
		if string(v) == "." {
			ans += "[.]"
			continue
		}

		ans += string(v)
	}

	return ans
}

func Q771JewelsAndStones(jewels string, stones string) int {
	ans := 0
	jl := len(jewels)
	sl := len(stones)
	for i := 0; i < jl; i++ {
		for k := 0; k < sl; k++ {
			if jewels[i] == stones[k] {
				ans++
			}
		}
	}
	return ans
}

func Q938RangeSumOfBST(root *TreeNode, low int, high int) int {
	ans := 0

	if root.Val >= low && root.Val <= high {
		ans += root.Val
	}

	if root.Left != nil {
		ans += Q938RangeSumOfBST(root.Left, low, high)
	}

	if root.Right != nil {
		ans += Q938RangeSumOfBST(root.Right, low, high)
	}

	return ans
}

func Q1281SubtractTheProductAndSumOfDigitsOfAnInteger(n int) int {
	arr := []int{}

	for n > 0 {
		arr = append(arr, n%10)
		n /= 10
	}

	pod := 1
	sod := 0

	for _, v := range arr {
		pod *= v
		sod += v
	}

	return pod - sod
}

func Q1295FindNumbers(nums []int) int {
	ans := 0

	for _, v := range nums {
		d := 0
		for v > 0 {
			v /= 10
			d++
		}

		if d%2 == 0 {
			ans++
		}
	}

	return ans
}

func Q643FindMaxAverage(nums []int, k int) float64 {
	var ans, sum int

	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	ans = sum

	for j := k; j < len(nums); j++ {
		sum += nums[j] - nums[j-k]
		if sum > ans {
			ans = sum
		}
	}

	return float64(ans) / float64(k)
}

func Q697FindShortestSubArray(nums []int) int {
	ans := math.MaxInt
	m := make(map[int][]int)
	maxDegree := 0

	// put duplicated elements and their index (elements as key and all the index as the value, as slice)
	for i, v := range nums {
		m[v] = append(m[v], i)
	}

	// count most one and calc the lens
	for _, v := range nums {
		tmpDegree := len(m[v])
		l := len(nums[m[v][0] : (m[v][tmpDegree-1])+1])
		if tmpDegree > maxDegree {
			ans = l
			maxDegree = tmpDegree
		}

		if tmpDegree == maxDegree {
			if ans > l {
				ans = l
			}
		}
	}

	return ans
}

func Q2121GetDistances(arr []int) []int64 {
	l := len(arr)

	if l == 1 {
		return []int64{0}
	}

	ans := make([]int64, l)
	d := make(map[int][]int)
	sum := make(map[int]int64)

	for i, v := range arr {
		tmp, ok := d[v]
		if !ok {
			tmp = []int{}
		}
		tmp = append(tmp, i)
		d[v] = tmp
		sum[v] += int64(i)
	}

	// for intervals
	// [2,1,3,1,2,3,3] 3 => 2,5,6 => sum = 13
	// index 2 => left = 0 right = sum - left - self => 13 - 0 - 2 = 11
	// interval => right - left - (len(right) - len(left))*self => 11 - 0 - (2 - 0)*2 => 11 - 4 = 7
	// index 5 => left = 2 right = 13 - 2 - 5 = 6
	// interval => 6 - 2 - (1 - 1)*5 = 4
	// index 6 => left = 7 right = 13 - 7 - 6 = 0
	// interval => 0 - 7 - (0 - 2)*6 = 5
	for n, indexs := range d {
		var left int64 = 0
		var right int64 = 0
		for pos, index := range indexs {
			right = sum[n] - left - int64(index)
			ans[index] = right - left - int64((len(indexs)-pos-1)-(pos))*int64(index)
			left += int64(index)
		}
	}

	return ans

}

func Q1009BitwiseComplement(n int) int {
	ans := 0

	s := fmt.Sprintf("%b", n)
	b := []byte(s)
	tmp := []int{}

	// reverse slice to let the binary number start at right instead of left
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	for _, v := range b {
		if v == byte('1') {
			tmp = append(tmp, 0)
			continue
		}

		tmp = append(tmp, 1)
	}

	for i, v := range tmp {
		ans += v << i
	}

	return ans
}

func Q2206DivideArray(nums []int) bool {
	l := len(nums)
	pairs := l / 2

	sum := make(map[int]int)

	for _, v := range nums {
		sum[v]++
	}

	for _, n := range sum {
		// elements only appear odd times
		if n%2 == 1 {
			return false
		}

		// if pairs < 0 means not enough pair can be divided
		pairs -= n / 2
		if pairs < 0 {
			return false
		}
	}

	return true
}

func Q150EvalRPN(tokens []string) int {
	calc := func(s string, n1 int, n2 int) int {
		if s == "+" {
			return n1 + n2
		} else if s == "-" {
			return n1 - n2
		} else if s == "*" {
			return n1 * n2
		} else {
			return int(n1 / n2)
		}
	}

	d := []int{}

	for _, v := range tokens {
		if v == "+" || v == "-" || v == "*" || v == "/" {
			res := calc(v, d[len(d)-2], d[len(d)-1])
			d = d[:len(d)-2]
			d = append(d, res)
		} else {
			v, _ := strconv.Atoi(v)
			d = append(d, v)
		}
	}

	return d[0]
}

func Q59GenerateMatrix(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	colStart := 0
	colEnd := n - 1

	rowStart := 0
	rowEnd := n - 1

	max := n * n

	v := 1
	for v <= max {
		// left to right
		for i := colStart; i <= max && i <= colEnd; i++ {
			ans[colStart][i] = v
			v++
		}
		rowStart++

		// top to bottom
		for i := rowStart; i <= max && i <= rowEnd; i++ {
			ans[i][colEnd] = v
			v++
		}
		colEnd--

		// right to left
		for i := colEnd; i <= max && i >= colStart; i-- {
			ans[rowEnd][i] = v
			v++
		}
		rowEnd--

		// bottom to top
		for i := rowEnd; i <= max && i >= rowStart; i-- {
			ans[i][colStart] = v
			v++
		}
		colStart++
	}

	return ans
}
