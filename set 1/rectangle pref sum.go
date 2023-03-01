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
	var n, m, k int
	fscan(&n, &m, &k)
	pref := make([][]int, n+1)
	for i := range pref {
		pref[i] = make([]int, m+1)
	}
	var cur int
	for row := 1; row < n+1; row++ {
		for col := 1; col < m+1; col++ {
			fscan(&cur)
			pref[row][col] = cur + pref[row-1][col] + pref[row][col-1] - pref[row-1][col-1]
		}
	}
	//fprintln(pref)
	var x1, y1, x2, y2, res int
	for i := 0; i < k; i++ {
		fscan(&x1, &y1, &x2, &y2)
		res = pref[x2][y2] - pref[x1-1][y2] - pref[x2][y1-1] + pref[x1-1][y1-1]
		fprintln(res)
	}
}
