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
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, m+2)
	}
	dp[0][1] = 1
	for i := 2; i < n+2; i++ {
		for j := 2; j < m+2; j++ {

			dp[i][j] = dp[i-2][j-1] + dp[i-1][j-2]
		}
	}
	fprintln(dp[n+1][m+1])
}
