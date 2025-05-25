# キーワード
`累積和`

# コード
```go
type Test struct {
	Class int
	Score int
}

func main() {
	defer flush()

	// Implement
	n := scanAsInt()
	scores := make([]Test, n)
	for i := 0; i < n; i++ {
		c, p := scanAsInt(), scanAsInt()
		scores[i] = Test{
			Class: c,
			Score: p,
		}
	}
	q := scanAsInt()
	one := 0
	two := 0
	for i := 0; i < q; i++ {
		l, r := scanAsInt(), scanAsInt()
		tmp := scores[l-1 : r]
		for _, test := range tmp {
			if test.Class == 1 {
				one += test.Score
			} else {
				two += test.Score
			}
		}
		fmt.Fprintf(w, "%d %d\n", one, two)
		one = 0
		two = 0
	}
}
```

# 改善点
loop を全てに対して回して、尚且つ Q の数だけ内側で loop を回していたため O(NQ) になっていた。
累積和を求めておけば、例えば 1 組の 2 - 6 が欲しい時は S[5] - S[1] で求めることができる。O(N+Q) になる。