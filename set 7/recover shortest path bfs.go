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
	//read graph
	for i := range graph {
		graph[i] = make([]int, n)
		for j := range graph[i] {
			fscan(&graph[i][j])
		}
	}
	// read start and end vertices
	var start, end int
	fscan(&start, &end)
	//	depth[i] - minimal distance from (start - 1) to i
	depth := make([]int, n)
	for i := range depth {
		depth[i] = -1
	}

	dCount := 0
	depth[start-1] = dCount
	// prev[i] - previous to i vertex
	prev := make([]int, n)
	prev[start-1] = -1
	// bfs
	q := []int{start - 1}
	for len(q) > 0 {
		nextQ := []int{}
		dCount++
		for _, cur := range q {
			for k, v := range graph[cur] {
				if v == 1 && depth[k] == -1 {
					depth[k] = dCount
					nextQ = append(nextQ, k)
					prev[k] = cur
				}
			}
		}
		q = nextQ
	}
	// print result
	fprintln(depth[end-1])
	if depth[end-1] == -1 || depth[end-1] == 0 {
		return
	}
	// recover path
	p := prev[end-1]
	path := []int{end - 1}
	for p != -1 {
		path = append(path, p)
		p = prev[p]
	}
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Fprint(out, path[i]+1, " ")
	}
}
