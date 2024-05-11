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
	a, b, c := scanInt(), scanInt(), scanInt()
	s := scanString()

	_, err := fmt.Fprintf(w, "%v %s", a+b+c, s)
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
