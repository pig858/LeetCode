package main

import (
	"fmt"
	"math"
	"sort"
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
