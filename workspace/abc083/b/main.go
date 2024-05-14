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
	n, a, b := scanAsInt(), scanAsInt(), scanAsInt()
	if n < 1 || 10000 < n {
		panic("n must be between 1 and 10000")
	}
	if a < 1 || 36 < a || b < 1 || 36 < b {
		panic("a must be between 1 and 36")
	}

	sum := 0
	for i := 1; i <= n; i++ {
		v := 0
		if i < 10 {
			v = i
		} else if i < 100 {
			xx := i / 10
			x := i % 10
			v = xx + x
		} else if i < 1000 {
			xxx := i / 100
			xx := i % 100 / 10
			x := i % 100 % 10
			v = xxx + xx + x
		} else if i < 10000 {
			xxxx := i / 1000
			xxx := i % 1000 / 100
			xx := i % 1000 % 100 / 10
			x := i % 1000 % 100 % 10
			v = xxxx + xxx + xx + x
		} else {
			// i == 10000
			v = 1
		}
		if a <= v && v <= b {
			sum += i
		}
	}
	_, e := fmt.Fprintln(w, sum)
	if e != nil {
		panic(e)
	}
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

	for i, c := range input {
		var e error
		arr[i], e = strconv.Atoi(string(c))
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

func mergeSort(unsorted []int) []int {
	if len(unsorted) < 2 {
		return unsorted
	}

	mid := len(unsorted) / 2

	// 再帰的に関数を実行してソート済のスライスを返す
	left := mergeSort(unsorted[:mid])
	right := mergeSort(unsorted[mid:])

	sorted := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// マージ
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			sorted = append(sorted, left[i])
			i++
		} else {
			sorted = append(sorted, right[j])
			j++
		}
	}

	// 残りを配列に追加
	sorted = append(sorted, left[i:]...)
	sorted = append(sorted, right[j:]...)

	return sorted
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
