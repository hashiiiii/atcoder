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
	a, b, c, x := scanAsInt(), scanAsInt(), scanAsInt(), scanAsInt()
	if a > 50 {
		a = 50
	} else if a < 0 {
		a = 0
	}
	if b > 50 {
		b = 50
	} else if b < 0 {
		b = 0
	}
	if c > 50 {
		c = 50
	} else if c < 0 {
		c = 0
	}

	count := 0
	for i := 0; i <= c; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= a; k++ {
				v := i*50 + j*100 + k*500
				if v == x && 50 <= v && v <= 20000 && 1 <= i+j+k {
					count++
				}
			}
		}
	}
	_, e := fmt.Fprintln(w, count)
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

func scanAsSplitInt() []int {
	sc.Scan()
	input := sc.Text()
	arr := make([]int, len(input))
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
