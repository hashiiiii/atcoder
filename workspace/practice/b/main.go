package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
	n, q := scanInt(), scanInt()
	alphaLong := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}
	alphaShort := []string{
		"A", "B", "C", "D", "E",
	}

	var alpha []string
	var queryMax int
	if q == 1000 {
		alpha = alphaLong
		queryMax = q
	} else if q == 100 {
		alpha = alphaShort
		queryMax = q
	} else {
		alpha = alphaShort
		queryMax = q
	}

	for i := 0; i < queryMax; i++ {
		c1 := alpha[i%(n-1)]
		c2 := alpha[(i+1)%(n-1)]
		_, err := fmt.Fprintf(w, "? %v %v\n", c1, c2)
		if err != nil {
			panic(err)
		}
		s := scanString()
		if s == ">" {
			alpha[i%(n-1)] = c2
			alpha[(i+1)%(n-1)] = c1
		}
	}
	_, err := fmt.Fprintf(w, "! %v\n", strings.Join(alpha, ""))
	if err != nil {
		panic(err)
	}
}

func scanInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		_ = fmt.Errorf("[scanInt] error scanning integer: %s", e)
	}
	return i
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func flush() {
	e := w.Flush()
	if e != nil {
		_ = fmt.Errorf("[flush] error flushing: %s", e)
	}
}

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
