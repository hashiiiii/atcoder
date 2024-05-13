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

	// 1. scan
	arr := scanAsSplitInt()
	i := 0
	for _, x := range arr {
		if x == 1 {
			i++
		}
	}

	_, err := fmt.Fprintln(w, i)
	if err != nil {
		panic(err)
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

func scanAsSplitInt() []int {
	sc.Scan()
	s := sc.Text()
	arr := make([]int, len(s))
	for i, c := range s {
		var err error
		arr[i], err = strconv.Atoi(string(c))
		if err != nil {
			_ = fmt.Errorf("[scanAsSplitInt] error scanning integer: %s", c)
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
