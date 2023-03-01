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
	var n, m int
	fscan(&n)
	arr := make([]int, 0)
	set := map[int]struct{}{}
	for i := 0; i < n; i++ {
		var cur int
		fscan(&cur)
		if _, ok := set[cur]; !ok {
			arr = append(arr, cur)
		}
		set[cur] = struct{}{}
	}
	sort.Ints(arr)
	fscan(&m)
	for i := 0; i < m; i++ {
		var cur int
		fscan(&cur)
		res := countLess(&arr, cur)
		fprintln(res)
	}
}

func countLess(arr *[]int, target int) int {
	left := 0
	right := len(*arr)
	for left < right {
		mid := (left + right) / 2
		if target > (*arr)[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
