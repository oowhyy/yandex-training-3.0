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
	if n == 1 {
		fprintln(0)
		fprintln(1)
		return
	}
	dp := make([]int, n+1)
	prev := make([]int, n+1)
	//dp[0] = 0, dp[1] = 0
	for i := 2; i <= n; i++ {
		if i%6 == 0 { ///////////////////////// %6
			if dp[i-1] <= dp[i/3] && dp[i-1] <= dp[i/2] {
				prev[i] = i - 1
				dp[i] = dp[i-1] + 1
			} else if dp[i/2] <= dp[i-1] && dp[i/2] <= dp[i/3] {
				prev[i] = i / 2
				dp[i] = dp[i/2] + 1
			} else if dp[i/3] <= dp[i-1] && dp[i/3] <= dp[i/2] {
				prev[i] = i / 3
				dp[i] = dp[i/3] + 1
			}
		} else if i%3 == 0 { //////////////////
			if dp[i/3] < dp[i-1] {
				prev[i] = i / 3
				dp[i] = dp[i/3] + 1
			} else {
				prev[i] = i - 1
				dp[i] = dp[i-1] + 1
			}
		} else if i%2 == 0 {
			if dp[i/2] < dp[i-1] {
				prev[i] = i / 2
				dp[i] = dp[i/2] + 1
			} else {
				prev[i] = i - 1
				dp[i] = dp[i-1] + 1
			}
		} else {
			dp[i] = 1 + dp[i-1]
			prev[i] = i - 1
		}
	}
	fprintln(dp[n])
	preverse(prev, n)
}

func preverse(prev []int, cur int) {
	if cur == 0 {
		return
	}
	preverse(prev, prev[cur])
	fmt.Fprint(out, cur, " ")
}
