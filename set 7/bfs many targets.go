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
}

func solve() {
	var n, m, tx, ty, s, x, y int
	// read input
	fscan(&n, &m, &ty, &tx, &s)
	tx--
	ty--
	targets := make([]point, s)
	for i := 0; i < s; i++ {
		fscan(&x, &y)
		targets[i] = point{y - 1, x - 1}
	}
	//	depth[i][j] - minimal distance from (i, j) to (ty, tx)
	depth := make([][]int, n)
	for i := range depth {
		depth[i] = make([]int, m)
		for j := range depth[i] {
			depth[i][j] = -1
		}
	}
	depth[ty][tx] = 0
	dCount := 0
	dirs := [][]int{{1, 2}, {1, -2}, {2, 1}, {2, -1}, {-1, 2}, {-1, -2}, {-2, 1}, {-2, -1}}
	// bfs
	q := []point{{tx, ty}}
	for len(q) > 0 {
		nextQ := []point{}
		dCount++
		for _, cur := range q {
			// iterate all neighbours
			for _, d := range dirs {
				if cur.x+d[0] < m && cur.x+d[0] >= 0 && cur.y+d[1] >= 0 && cur.y+d[1] < n {
					if depth[cur.y+d[1]][cur.x+d[0]] == -1 {
						nextQ = append(nextQ, point{cur.x + d[0], cur.y + d[1]})
						depth[cur.y+d[1]][cur.x+d[0]] = dCount
					}
				}
			}

		}
		q = nextQ
	}
	// calculate the result
	res := 0
	for _, t := range targets {
		cur := depth[t.y][t.x]
		if cur == -1 {
			fprintln(-1)
			return
		}
		res += cur
	}
	// print result
	// fprintln(depth)
	fprintln(res)
}
