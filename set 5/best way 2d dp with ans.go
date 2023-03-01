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
	//inf := 100000000
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	cur := 0
	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, m)
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			fscan(&cur)
			arr[i-1][j-1] = cur
			dp[i][j] = cur + max(dp[i-1][j], dp[i][j-1])
		}
	}
	x, y := m, n
	res := []rune{}
	for x+y > 2 {
		if y > 1 && (x == 1 || dp[y][x]-arr[y-1][x-1] == dp[y-1][x]) {
			res = append(res, 'D')
			y--
		} else {
			res = append(res, 'R')
			x--
		}
	}
	fprintln(dp[n][m])
	for i := range res {
		fmt.Fprint(out, string(res[len(res)-i-1]), " ")
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
