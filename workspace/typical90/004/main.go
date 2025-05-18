package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func main() {
	defer flush()

	// Implement
	H, W := scanAsInt(), scanAsInt()
	i := 0
	j := 0
	arr := make([][]int, H)
	for i := 0; i < H; i++ {
		arr[i] = make([]int, W)
	}
	for i < H {
		for j < W {
			v := scanAsInt()
			arr[i][j] = v
			j++
		}
		i++
		j = 0
	}
	i = 0

	yoko := make([]int, H)
	tate := make([]int, W)
	for i < H {
		for j < W {
			yoko[i] += arr[i][j]
			tate[j] += arr[i][j]
			j++
		}
		i++
		j = 0
	}
	i = 0

	for i < H {
		for j < W {
			v := yoko[i] + tate[j] - arr[i][j]
			if (j+1)%W == 0 {
				fmt.Fprintf(w, " %d\n", v)
			} else if (j+1)%W == 1 {
				fmt.Fprint(w, v)
			} else {
				fmt.Fprintf(w, " %d", v)
			}
			j++
		}
		i++
		j = 0
	}
	i = 0
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
