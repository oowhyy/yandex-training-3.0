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
		graph[b] = append(graph[b], a)
	}
	colors := map[int]int{}
	var dfs func(int, int)

	res := true
	dfs = func(start, color int) {
		colors[start] = color
		for _, v := range graph[start] {
			if vColor, ok := colors[v]; !ok {
				dfs(v, 3-color)
			} else {
				if vColor == color {
					res = false
					return
				}
			}

		}
	}
	for i := 1; i < n+1; i++ {
		if colors[i] == 0 {
			dfs(i, 1)
		}
	}
	if res {
		fprintln("YES")
	} else {
		fprintln("NO")
	}
}
