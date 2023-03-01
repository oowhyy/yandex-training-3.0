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
	var n, m int
	fscan(&n, &m)
	inf := 100000000
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := range dp[0] {
		dp[0][i] = inf
	}
	for i := range dp {
		dp[i][0] = inf
	}
	dp[1][0] = 0
	cur := 0
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			fscan(&cur)
			dp[i][j] = cur + min(dp[i-1][j], dp[i][j-1])
		}
	}
	fprintln(dp[n][m])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
