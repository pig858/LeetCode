package main

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
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

func Q20IsValid(s string) bool {

	l := len(s)

	if l == 1 {
		return false
	}

	ans := true
	b := []byte(s)
	m := make(map[string]int)
	// pb => parentheses brackets => (
	// sb => square brackets => [
	// cb => curly brackets => {
	m["pb"] = 0
	m["sb"] = 0
	m["cb"] = 0
	// f is a queue for which open brackets it is
	f := []string{}
	openBracket := 0

	for _, v := range b {

		switch string(v) {
		case "(":
			m["pb"]++
			f = append(f, "pb")
			openBracket++
		case ")":
			if openBracket == 0 || f[len(f)-1] != "pb" || m["pb"] == 0 {
				return false
			}
			m["pb"]--
			f = f[:len(f)-1]
			openBracket--
		case "[":
			m["sb"]++
			f = append(f, "sb")
			openBracket++
		case "]":
			if openBracket == 0 || f[len(f)-1] != "sb" || m["sb"] == 0 {
				return false
			}
			m["sb"]--
			f = f[:len(f)-1]
			openBracket--
		case "{":
			m["cb"]++
			f = append(f, "cb")
			openBracket++
		case "}":
			if openBracket == 0 || f[len(f)-1] != "cb" || m["cb"] == 0 {
				return false
			}
			m["cb"]--
			f = f[:len(f)-1]
			openBracket--
		}
	}

	if len(f) != 0 {
		return false
	}

	return ans
}

func Q1290GetDecimalValue(head *ListNode) int {

	ans := 0

	for head != nil {
		ans = ans*2 + head.Val
		head = head.Next
	}

	return ans
}

func Q1832CheckIfPangram(sentence string) bool {

	m := make(map[string]int)
	b := []byte(sentence)

	for _, v := range b {
		s := string(v)
		m[s]++
	}

	return len(m) == 26
}

func Q21MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	t := &ListNode{}
	ans := t

	for list1 != nil {
		if list2 == nil {
			t.Next = list1
			list1 = list1.Next
			t = t.Next
			continue
		}

		if list1.Val > list2.Val {
			t.Next = list2
			list2 = list2.Next
			t = t.Next
		} else {
			t.Next = list1
			list1 = list1.Next
			t = t.Next
		}

	}

	for list2 != nil {
		t.Next = list2
		list2 = list2.Next
		t = t.Next

	}

	return ans.Next
}

func Q38CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	s := Q38CountAndSay(n - 1)

	m := []string{}
	c := 0
	pre := ""
	b := []byte(s)

	for _, v := range b {
		if pre == "" {
			pre = string(v)
			c++
			continue
		}

		if pre != string(v) {
			var tmp bytes.Buffer
			tmp.WriteString(strconv.FormatInt(int64(c), 10))
			tmp.WriteString(pre)
			m = append(m, tmp.String())
			c = 1
			pre = string(v)
			continue
		}

		c++

	}

	var tmp bytes.Buffer
	tmp.WriteString(strconv.FormatInt(int64(c), 10))
	tmp.WriteString(pre)
	m = append(m, tmp.String())

	return strings.Join(m, "")
}

func Q2AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	t := &ListNode{}
	ans := t
	// c for carry res >= 10 will be 1
	c := 0
	res := 0

	for l1 != nil {
		if l2 == nil {
			res = l1.Val + c
			c = 0
			if res >= 10 {
				res -= 10
				c = 1
			}
			node := &ListNode{
				Val:  res,
				Next: nil,
			}

			t.Next = node
			t = t.Next
			l1 = l1.Next

			continue
		}

		res = l1.Val + l2.Val + c
		c = 0
		if res >= 10 {
			res -= 10
			c = 1
		}

		node := &ListNode{
			Val:  res,
			Next: nil,
		}

		t.Next = node
		t = t.Next
		l1 = l1.Next
		l2 = l2.Next

	}

	for l2 != nil {
		res = l2.Val + c
		c = 0
		if res >= 10 {
			res -= 10
			c = 1
		}

		node := &ListNode{
			Val:  res,
			Next: nil,
		}

		t.Next = node
		t = t.Next
		l2 = l2.Next
	}

	if c == 1 {
		node := &ListNode{
			Val:  1,
			Next: nil,
		}
		t.Next = node
	}

	return ans.Next
}

func Q692TopKFrequent(words []string, k int) []string {

	if len(words) == 1 {
		return words
	}

	// w for the word
	// t for the times
	type data struct {
		w string
		t int
	}

	m := make(map[string]int)

	for _, v := range words {
		m[v]++
	}

	var d []data

	for s, t := range m {
		d = append(d, data{
			w: s,
			t: t,
		})
	}

	sort.Slice(d, func(i int, j int) bool {
		// if two word appear same times
		// return the word by lexicographical order
		if d[i].t == d[j].t {
			return d[i].w < d[j].w
		}

		// return the higher frequency
		return d[i].t > d[j].t
	})

	ans := make([]string, k)
	i := 0
	for i < k {
		ans[i] = d[i].w
		i++
	}

	return ans
}

func Q12IntToRoman(num int) string {
	var ans bytes.Buffer

	n := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	r := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	i := 0
	for i < 13 {
		for num >= n[i] {
			ans.WriteString(r[i])
			num -= n[i]
		}
		i++
	}

	return ans.String()
}

func Q219ContainsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)

	abs := func(i int, j int) int {
		if (i - j) < 0 {
			return -(i - j)
		}

		return (i - j)
	}

	for i := range nums {
		j, ok := m[nums[i]]
		if ok && abs(j, i) <= k {
			return true
		}

		m[nums[i]] = i
	}

	return false
}

func Q76MinWindow(s string, t string) string {
	left, right := 0, 0
	start := 0
	minLen := len(s) + 1

	window := make(map[byte]int)
	need := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}

	match := 0

	for right < len(s) {
		_, ok := need[s[right]]
		if ok {
			window[s[right]]++
			if window[s[right]] == need[s[right]] {
				match++
			}
		}
		right++

		for {
			if match != len(need) {
				break
			}

			if right-left < minLen {
				start = left
				minLen = right - left
			}

			_, ok := need[s[left]]
			if ok {
				window[s[left]]--
				if window[s[left]] < need[s[left]] {
					match--
				}
			}

			left++
		}
	}

	if minLen > len(s) {
		return ""
	} else {
		res := []byte(s)[start : minLen+start]
		return string(res)
	}
}

// 正整數 輸出 字串
// 從1開始到n 輸出
// 3或3的倍數 改輸出fizz
// 5或5的倍數 改輸出buzz
// 15的話 fizzbuzz
// n = 10 => 1 2 3 4 5 6 7 8 9 10
// 數字間用空白隔開
func CountInt(n int) string {
	var ans bytes.Buffer

	for i := 1; i <= n; i++ {

		if i%3 == 0 || i%5 == 0 {

			if i%3 == 0 {
				// 3或3的倍數指定的字串
				ans.WriteString("fizz")
			}

			if i%5 == 0 {
				// 5或5的倍數指定的字串
				ans.WriteString("buzz")
			}

		} else {
			// 將數字轉成字串並寫入
			ans.WriteString(strconv.Itoa(i))
		}

		// 因數字間要用空白隔開 但跟n相同時不需要最後的那個空白
		if i < n {
			ans.WriteString(" ")
		}
	}

	//輸出字串
	return ans.String()
}

func Helloworld() string {
	return "hello world"
}

func Sort(input []int) string {

	sort.Ints(input)

	ans := strings.Trim(fmt.Sprint(input), "[]")

	return ans
}

func Card(input []int) int {

	// 裝牌的
	nums := make(map[int]int, len(input))
	// 裝花色的
	colors := make(map[int]int, len(input))

	for _, v := range input {
		n := (0x0FF & v)
		nums[n]++

		c := (0xF00 & v)
		colors[c]++
	}

	num := len(nums)
	color := len(colors)

	if color > 1 {
		switch num {

		// 葫蘆 (4)或 金剛(3)
		case 2:
			for _, v := range nums {
				if v == 4 {
					return 3
				}
			}
			return 4

		// 3條(7) 或 兩對(8)
		case 3:
			for _, v := range nums {
				if v == 3 {
					return 7
				}
			}
			return 8

		// 一對(9)
		case 4:
			return 9
		// 順子(6) 或 高牌(10)
		case 5:
			if isStraight(input) {
				return 6
			}
			return 10
		}

	} else {
		// 是順子就是同花順(2)/皇家同花順(1) 不然就是同花(5)
		if !isStraight(input) {
			return 5
		}

		return isRoyal(input)
	}

	// 皆非以上
	return -1
}

func isStraight(card []int) bool {

	nums := make([]int, 0, len(card))

	for _, v := range card {
		n := (0x0FF & v)
		nums = append(nums, n)
	}

	straight := false

	max := 0x002
	min := 0x00e
	sum := 0

	for _, v := range nums {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}

		sum += v
	}

	if max-min == 4 {
		straight = true
	}

	// A 2 3 4 5
	if max-min == 12 && sum == 28 {
		straight = true
	}

	return straight
}

func isRoyal(card []int) int {
	nums := make([]int, 0, len(card))

	for _, v := range card {
		n := (0x0FF & v)
		nums = append(nums, n)
	}

	max := 0x002
	min := 0x00e

	for _, v := range nums {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	if max == 0x00e {
		return 1
	}

	return 2
}

func JokerCard(input []int) int {

	// 裝牌的
	nums := make(map[int]int, len(input))
	// 裝花色的
	colors := make(map[int]int, len(input))
	// 鬼牌
	//0x50f代表小王
	//0x610代表大王
	joker := 0

	for _, v := range input {

		if v == 0x50f || v == 0x610 {
			joker++
			continue
		}

		n := (0x0FF & v)
		nums[n]++

		c := (0xF00 & v)
		colors[c]++
	}

	num := len(nums)
	color := len(colors)

	if color > 1 {
		switch num {
		// 順子(6) 或 高牌(10)
		case 5:
			if isJokerStraight(input) {
				return 6
			}
			return 10
		// 有可能是順子(4 + joker) 或 一對
		case 4:
			if isJokerStraight(input) {
				return 6
			}
			return 9
		// 0 鬼牌 2 2 3 3 3 三條
		// 1 鬼牌 2 2 3 4 x 三條
		// 2 鬼牌 2 3 4 x x 順子
		// 3 張不一樣數字 + 最大2張鬼牌
		// 3條 順子
		case 3:
			if isJokerStraight(input) {
				return 6
			}
			return 7

		// 0 鬼牌 2 2 3 3 3 葫蘆
		// 1 鬼牌 2 2 3 3 x 葫蘆 / 2 2 2 3 x 金剛
		// 2 鬼牌 2 3 4 x x 順子 / 2 2 3 x x 金剛 / 2 2 2 x x 金剛
		case 2:
			if isJokerStraight(input) {
				return 6
			}

			for _, v := range nums {
				if v+joker == 4 {
					return 3
				}
			}

			return 4
		}
	} else {
		// 是順子就是同花順(2)/皇家同花順(1) 不然就是同花(5)
		if !isJokerStraight(input) {
			return 5
		}

		return isJokerRoyal(input)
	}

	// 皆非以上
	return -1
}

func isJokerStraight(card []int) bool {

	nums := make([]int, 0, len(card))
	joker := 0
	haveA := false

	for _, v := range card {

		if v == 0x50f || v == 0x610 {
			joker++
			continue
		}

		n := (0x0FF & v)

		if n == 0x00e {
			haveA = true
		}

		nums = append(nums, n)
	}

	straight := false

	max := 0x002
	min := 0x00e
	noAmax := 0x002
	noAmin := 0x00e
	sum := 0

	for _, v := range nums {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}

		if haveA {
			if v == 0x00e {
				continue
			}
			if v > noAmax {
				noAmax = v
			}
			if v < noAmin {
				noAmin = v
			}
		}

		sum += v
	}

	// 沒有鬼牌
	if joker == 0 {
		if max-min == 4 {
			straight = true
		}

		// A 2 3 4 5
		if max-min == 12 && sum == 28 {
			straight = true
		}

		//有鬼牌
	} else {

		// 2 2 3 x x

		// A 3 x 5 6
		if haveA {
			if noAmax-noAmin <= 3 {
				straight = true
			}

		} else {

			if max-min <= 4 {
				straight = true
			}
		}
	}

	return straight
}

func isJokerRoyal(card []int) int {
	nums := make([]int, 0, len(card))

	for _, v := range card {
		n := (0x0FF & v)
		nums = append(nums, n)
	}

	max := 0x002
	min := 0x00e

	for _, v := range nums {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	if max == 0x00e {
		return 1
	}

	return 2
}
