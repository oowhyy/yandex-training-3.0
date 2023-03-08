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
	graph := make([][]int, n+1)
	for i := 0; i < m; i++ {
		fscan(&a, &b)
		graph[a] = append(graph[a], b)
		//graph[b] = append(graph[b], a)
	}
	colors := map[int]int{}
	hasCycle := false
	res := []int{}
	// dfs
	var dfs func(int)
	dfs = func(start int) {
		// color grey on enter
		colors[start] = 1
		for _, v := range graph[start] {
			if vColor, ok := colors[v]; !ok {
				dfs(v)
			} else if vColor == 1 { // seeing grey vertex - there is cycle
				hasCycle = true
				return
			}

		}
		// color black on extit and append to result
		colors[start] = 2
		res = append(res, start)
	}
	// main loop
	for i := 1; i < n+1; i++ {
		if colors[i] == 0 {
			dfs(i)
		}
	}
	if hasCycle {
		fprintln("-1")
		return
	}
	for i := range res {
		fmt.Fprint(out, res[n-i-1], " ")
	}
}
