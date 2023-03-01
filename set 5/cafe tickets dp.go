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
	dp := make([][]int, n+1) // days x tickets
	for i := range dp {
		dp[i] = make([]int, n+3)
	}
	for i := range dp[0] {
		dp[0][i] = inf
	}
	for i := range dp {
		dp[i][0] = inf
		dp[i][n+2] = inf
	}
	dp[0][1] = 0
	cur := 0
	path := make([][]int, n) //
	for i := range path {
		path[i] = make([]int, n+1)
	}
	for day := 1; day < n+1; day++ {
		fscan(&cur)
		for t := 1; t < n+2; t++ {
			if cur > 100 { // can get a coupon?
				if dp[day-1][t+1] < cur+dp[day-1][t-1] { // use coupon?
					dp[day][t] = dp[day-1][t+1]
					path[day-1][t-1] = 1
				} else {
					dp[day][t] = cur + dp[day-1][t-1] // pay and get a coupon
					path[day-1][t-1] = -1
				}

			} else {
				if dp[day-1][t+1] < cur+dp[day-1][t] { // use coupon?
					dp[day][t] = dp[day-1][t+1]
					path[day-1][t-1] = 1
				} else {
					dp[day][t] = cur + dp[day-1][t] // just pay
				}
			}
		}
	}
	// for _, v := range path {
	// 	fprintln(v)
	// }
	// for _, v := range dp {
	// 	fprintln(v)
	// }
	// get dp ans - last min value (to maximize coupons left)
	min := inf + 1
	id := -1
	for k, v := range dp[n] {
		if v <= min {
			min = v
			id = k - 1
		}
	}
	//fprintln(dp)
	endCoupons := id
	ans := dp[n][id+1]
	res := []int{}
	for n > 0 {
		p := path[n-1][id]
		if p == 1 {
			res = append(res, n)
		}
		id += p
		n--
	}
	fprintln(ans)
	fprintln(endCoupons, len(res))
	for i := range res {
		fprintln(res[len(res)-1-i])
	}
	//fprintln(couponDays)
}
