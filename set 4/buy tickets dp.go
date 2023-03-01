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
	inf := 100000000
	dp := []int{0, 0, 0}
	a := make([]int, n+3)
	b := make([]int, n+3)
	c := make([]int, n+3)
	a[0] = inf
	a[1] = inf
	a[2] = inf
	b[0] = inf
	b[1] = inf
	b[2] = inf
	c[0] = inf
	c[1] = inf
	c[2] = inf
	for i := 3; i < n+3; i++ {
		fscan(&a[i], &b[i], &c[i])
	}
	for i := 3; i < n+3; i++ {
		best := min(dp[i-1]+a[i], min(dp[i-2]+b[i-1], dp[i-3]+c[i-2]))
		dp = append(dp, best)
	}
	fprintln(dp[n+2])
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
