package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	w  = bufio.NewWriter(os.Stdout)
)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func rotateClockwise(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return matrix
	}
	n := len(matrix)

	// 正方行列なので、新しい行列の列数もn
	rotatedMatrix := make([][]int, n)
	for i := range rotatedMatrix {
		rotatedMatrix[i] = make([]int, n) // 各行の要素数もn
	}

	for r := 0; r < n; r++ { // 元の行列の行
		for c := 0; c < n; c++ { // 元の行列の列
			rotatedMatrix[c][n-1-r] = matrix[r][c]
		}
	}
	return rotatedMatrix
}

func main() {
	defer flush()

	// Implement
	var s [][]int
	var t [][]int
	n := scanAsInt()
	for i := 0; i < n; i++ {
		ss := scanAsString()
		b := []byte(ss)
		var arr []int
		for j := 0; j < len(b); j++ {
			if b[j] == '#' {
				arr = append(arr, 1)
			} else {
				arr = append(arr, 0)
			}
		}
		s = append(s, arr)
	}
	for i := 0; i < n; i++ {
		ts := scanAsString()
		b := []byte(ts)
		var arr []int
		for j := 0; j < len(b); j++ {
			if b[j] == '#' {
				arr = append(arr, 1)
			} else {
				arr = append(arr, 0)
			}
		}
		t = append(t, arr)
	}
	// 0
	c0 := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			se := t[i][j]
			te := s[i][j]
			if se^te == 1 {
				c0++
			}
		}
	}
	// 90
	c1 := 0
	s90 := rotateClockwise(s)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			se := t[i][j]
			te := s90[i][j]
			if se^te == 1 {
				c1++
			}
		}
	}
	c1 = c1 + 1
	// 180
	c2 := 0
	s180 := rotateClockwise(s90)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			se := t[i][j]
			te := s180[i][j]
			if se^te == 1 {
				c2++
			}
		}
	}
	c2 = c2 + 2
	// 270
	c3 := 0
	s270 := rotateClockwise(s180)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			se := t[i][j]
			te := s270[i][j]
			if se^te == 1 {
				c3++
			}
		}
	}
	c3 = c3 + 3
	arr := []int{c0, c1, c2, c3}
	sort.SliceStable(arr, func(i, j int) bool { return arr[i] > arr[j] })
	_, _ = fmt.Fprintln(w, arr[len(arr)-1])
}

/////////////////////////////
// Scan
/////////////////////////////

func scanAsInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		_ = fmt.Errorf("[scanAsInt] error scanning integer: %s", e)
	}
	return i
}

func scanAsSplitInt(digit int) []int {
	sc.Scan()
	input := sc.Text()

	var arr []int
	if len(input) < digit {
		arr = make([]int, digit)
	} else {
		arr = make([]int, len(input))
	}

	if len(arr) == 0 {
		return arr
	}

	offset := len(arr) - len(input)
	if offset < 0 {
		_ = fmt.Errorf("[scanAsSplitInt] offset too small: %d", offset)
	}

	for i, c := range input {
		var e error
		arr[offset+i], e = strconv.Atoi(string(c))
		if e != nil {
			_ = fmt.Errorf("[scanAsSplitInt] error scanning integer: %s", e)
		}
	}
	return arr
}

func scanAsString() string {
	sc.Scan()
	return sc.Text()
}

/////////////////////////////
// Sort
/////////////////////////////

func reverseString(s string) string {
	b := []byte(s)
	size := len(b)
	// あまりは切り捨てられる
	for i := 0; i < size/2; i++ {
		b[i], b[size-1-i] = b[size-1-i], b[i]
	}
	return string(b)
}

func mergeSort(unsorted []int, isAsc bool) []int {
	if len(unsorted) < 2 {
		return unsorted
	}

	mid := len(unsorted) / 2

	// 再帰的に関数を実行してソート済のスライスを返す
	left := mergeSort(unsorted[:mid], isAsc)
	right := mergeSort(unsorted[mid:], isAsc)

	sorted := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// マージ
	for i < len(left) && j < len(right) {
		if (isAsc && left[i] <= right[j]) || (!isAsc && left[i] >= right[j]) {
			sorted = append(sorted, left[i])
			i++
		} else {
			sorted = append(sorted, right[j])
			j++
		}
	}

	// 残りの要素を追加
	sorted = append(sorted, left[i:]...)
	sorted = append(sorted, right[j:]...)

	return sorted
}

/////////////////////////////
// distinct
/////////////////////////////

func distinct[T comparable](dup []T) []T {
	m := make(map[T]struct{}, len(dup))
	var uniq []T
	for _, e := range dup {
		if _, ok := m[e]; ok {
			continue
		}
		m[e] = struct{}{}
		uniq = append(uniq, e)
	}
	return uniq
}

/////////////////////////////
// Config
/////////////////////////////

func flush() {
	e := w.Flush()
	if e != nil {
		_ = fmt.Errorf("[flush] error flushing: %s", e)
	}
}
