package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func fscan(a ...interface{}) {
	fmt.Fscan(in, a...)
}
func fprintln(a ...interface{}) {
	fmt.Fprintln(out, a...)
}

func main() {
	defer out.Flush()
	solve()
}

func solve() {
	var n int
	fscan(&n)

	var cur int
	res := 0
	fscan(&cur)
	for i := 1; i < n; i++ {
		var nxt int
		fscan(&nxt)
		res += min(nxt, cur)
		cur = nxt
	}
	fprintln(res)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
