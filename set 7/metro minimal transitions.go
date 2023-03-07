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

type point struct {
	x int
	y int
	z int
}

func solve() {
	var n, m, c, st int
	fscan(&n, &m)
	lines := make([][]int, m)
	for l := range lines {
		fscan(&c)
		lines[l] = make([]int, c)
		for i := 0; i < c; i++ {
			fscan(&st)
			lines[l][i] = st - 1
		}
	}

	// V - lines, E - transitions
	graph := make([][]int, m)
	for i := range graph {
		graph[i] = []int{}
	}
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			if hasCommonStation(lines[i], lines[j]) {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
	}

	// one station can be on many lines => multiple starting positions, multiple targets
	var start, end int
	fscan(&start, &end)
	start--
	end--
	startLines := map[int]bool{}
	endLines := map[int]bool{}
	for lineId, l := range lines {
		if lineHasStation(l, start) {
			startLines[lineId] = true
		}
		if lineHasStation(l, end) {
			endLines[lineId] = true
		}
	}

	//depth[i] - minimal number of transitions form start lines to line i
	depth := make([]int, m)
	for i := range depth {
		depth[i] = -1
	}
	dCount := 0

	// bfs
	q := []int{}
	for k := range startLines {
		depth[k] = dCount
		q = append(q, k)
	}
	for len(q) > 0 {
		newQ := []int{}
		dCount++
		for _, cur := range q {
			// iterate all neighbours
			for _, v := range graph[cur] {
				if depth[v] == -1 {
					depth[v] = dCount
					newQ = append(newQ, v)
				}
			}
		}
		q = newQ
	}
	res := 1000000000
	for lineId := range endLines {
		if depth[lineId] != -1 && depth[lineId] < res {
			res = depth[lineId]
		}
	}
	if res == 1000000000 {
		fprintln(-1)
		return
	}
	fprintln(res)
}

func lineHasStation(arr1 []int, st int) bool {
	for _, v := range arr1 {
		if v == st {
			return true
		}
	}
	return false
}

func hasCommonStation(arr1, arr2 []int) bool {
	mp := map[int]bool{}
	for _, v := range arr1 {
		mp[v] = true
	}
	for _, v := range arr2 {
		if mp[v] {
			return true
		}
	}
	return false
}
