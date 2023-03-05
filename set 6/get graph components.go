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
	var n, m, a, b int
	fscan(&n, &m)
	edges := map[int][]int{}
	for i := 0; i < m; i++ {
		fscan(&a, &b)
		edges[a] = append(edges[a], b)
		edges[b] = append(edges[b], a)
	}
	seen := map[int]bool{}
	var dfs func(int)
	comp := []int{}
	dfs = func(start int) {
		seen[start] = true
		comp = append(comp, start)
		for _, v := range edges[start] {
			if !seen[v] {
				dfs(v)
			}
		}
	}
	res := [][]int{}
	for i := 1; i <= n; i++ {
		if !seen[i] {
			dfs(i)
			res = append(res, comp[:])
			comp = []int{}
		}
	}
	fprintln(len(res))
	for _, arr := range res {
		fprintln(len(arr))
		for _, vv := range arr {
			fmt.Fprint(out, vv, " ")
		}
		fprintln()
	}
}
