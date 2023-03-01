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
	fscan(&n)
	a := make([]int, n)
	for i := range a {
		fscan(&a[i])
	}
	fscan(&m)
	b := make([]int, m)
	for i := range b {
		fscan(&b[i])
	}
	// dp[i][j] - lcs for a[:i] and b[:j]
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}

		}
	}
	// recover answer
	res := make([]int, 0, dp[n][m])
	i, j := n, m
	for i > 0 && j > 0 {
		if a[i-1] == b[j-1] {
			res = append(res, a[i-1])
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}
	for i := range res {
		fmt.Fprint(out, res[len(res)-1-i], " ")
	}

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
