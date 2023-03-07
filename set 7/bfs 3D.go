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
	var n int
	var row string
	var start point

	// read input
	fscan(&n)
	graph := make([][][]bool, n)
	for i := 0; i < n; i++ {
		graph[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			graph[i][j] = make([]bool, n)
			fscan(&row)
			for k := 0; k < n; k++ {
				graph[i][j][k] = row[k] == '.'
				if row[k] == 'S' {
					start = point{i, j, k}
				}
			}
		}
	}

	//	depth[i][j][k] - minimal distance from start to (i, j, k)
	depth := make([][][]int, n)
	for i := range depth {
		depth[i] = make([][]int, n)
		for j := range depth[i] {
			depth[i][j] = make([]int, n)
			for k := range depth[i][j] {
				depth[i][j][k] = -1
			}

		}
	}
	depth[start.x][start.y][start.z] = 0
	dCount := 0

	// bfs
	dirs := [][]int{{0, 0, 1}, {0, 0, -1}, {0, 1, 0}, {0, -1, 0}, {1, 0, 0}, {-1, 0, 0}}
	q := []point{start}
	for len(q) > 0 {
		nextQ := []point{}
		dCount++
		for _, cur := range q {
			// iterate all neighbours
			for _, d := range dirs {
				xx := (cur.x+d[0] < n && cur.x+d[0] >= 0)
				yy := (cur.y+d[1] < n && cur.y+d[1] >= 0)
				zz := (cur.z+d[2] < n && cur.z+d[2] >= 0)
				if xx && yy && zz && graph[cur.x+d[0]][cur.y+d[1]][cur.z+d[2]] {
					if depth[cur.x+d[0]][cur.y+d[1]][cur.z+d[2]] == -1 {
						nextQ = append(nextQ, point{cur.x + d[0], cur.y + d[1], cur.z + d[2]})
						depth[cur.x+d[0]][cur.y+d[1]][cur.z+d[2]] = dCount
					}
				}
			}

		}
		q = nextQ
	}

	// calculate the result
	res := 1000000000
	// i == 0 - top layer
	for j := range depth {
		for k := range depth {
			if depth[0][j][k] != -1 && depth[0][j][k] < res {
				res = depth[0][j][k]
			}
		}
	}

	// print result
	//fprintln(depth)
	fprintln(res)
}
