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
	var m, n int
	fscan(&m, &n)
	arr := make([][]int, n)
	for r := range arr {
		arr[r] = make([]int, 2)
	}
	res := 1
	for i := 0; i < n; i++ {
		fscan(&arr[i][0], &arr[i][1])
	}
	if n == 0 {
		fprintln(0)
		return
	}

	for i := n - 2; i >= 0; i-- {
		flag := false
		for j := i + 1; j < n; j++ {
			if !ok(arr[i][0], arr[i][1], arr[j][0], arr[j][1]) {
				flag = true
				break
			}
		}
		if !flag {
			res++
		}
	}
	fprintln(res)
}

func ok(l1, r1, l2, r2 int) bool {
	return (l1 > r2 || r1 < l2)
}
