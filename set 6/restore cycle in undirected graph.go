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
	graph := make([][]int, n)

	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fscan(&graph[i][j])
		}
	}
	colors := make([]int, n)
	cyclePath := []int{}
	seenVisited := -1 // cycle start
	stop := false     // path already restored
	// dfs
	var dfs func(int, int)
	dfs = func(start, prev int) {
		colors[start] = 1
		for i := 0; i < n; i++ {
			if graph[start][i] == 1 {
				if i == prev {
					continue
				}
				if colors[i] == 1 { // seeing already visited vertex during dfs => has cycles
					seenVisited = i
					cyclePath = append(cyclePath, start)
					return
				}
				dfs(i, start)
				if seenVisited != -1 {
					if !stop {
						cyclePath = append(cyclePath, start)
					}
					if start == seenVisited {
						stop = true
					}
					return
				}
			}
		}
	}
	// main loop
	for i := 0; i < n; i++ {
		if colors[i] == 0 {
			dfs(i, -1)
			if seenVisited != -1 {
				break
			}
		}
	}
	// print answer
	if seenVisited == -1 {
		fprintln("NO")
		return
	}
	fprintln("YES")
	fprintln(len(cyclePath))
	for i := range cyclePath {
		fmt.Fprint(out, cyclePath[i]+1, " ")
	}
}
