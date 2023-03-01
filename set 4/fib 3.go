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
	dp := []int{1, 2, 4}
	for i := 2; i < n; i++ {
		l := len(dp)
		dp = append(dp, dp[l-1]+dp[l-2]+dp[l-3])
	}
	fprintln(dp[n])
}
