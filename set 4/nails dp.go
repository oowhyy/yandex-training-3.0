package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	dp := []int{inf, inf}
	arr := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fscan(&arr[i])
	}
	sort.Ints(arr)
	dp = append(dp, arr[2]-arr[1])
	for i := 3; i <= n; i++ {
		diff := arr[i] - arr[i-1]
		dp = append(dp, diff+min(dp[i-1], dp[i-2]))
	}
	fprintln(dp[n])
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
