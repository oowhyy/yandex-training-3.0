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
	res := []int{}
	dfs = func(start int) {
		seen[start] = true
		res = append(res, start)
		for _, v := range edges[start] {
			if !seen[v] {
				dfs(v)
			}
		}
	}
	dfs(1)
	sort.Ints(res)
	fprintln(len(res))
	for _, v := range res {
		fmt.Fprint(out, v, " ")
	}
}
