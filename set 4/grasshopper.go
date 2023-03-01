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
	var n, k int
	fscan(&n, &k)
	dp := []int{0, 1}
	for i := 2; i <= n; i++ {
		max(i-k, 1)
		dp = append(dp, sum(dp[max(i-k, 1):i]))
	}
	fprintln(dp[n])
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func sum(arr []int) int {
	res := 0

	for _, v := range arr {
		res += v
	}
	return res
}
