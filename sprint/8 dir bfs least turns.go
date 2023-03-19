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
	var n, m int
	var inp string
	fscan(&n, &m)
	grid := make([][]int, n)

	for i := range grid {
		grid[i] = make([]int, m)
		fscan(&inp)
		for j := range grid[0] {
			if inp[j] == '.' {
				grid[i][j] = 1
			} else {
				grid[i][j] = 0
			}
		}
	}

	var sx, sy int
	fscan(&sx, &sy)
	start := point{sx - 1, n - sy}
	fscan(&sx, &sy)
	end := point{sx - 1, n - sy}

	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, m)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, -1}, {-1, 1}, {1, 1}, {-1, -1}}
	// bfs
	depth := 1
	dist[start.y][start.x] = depth
	q := []point{start}
	for len(q) > 0 {
		newQ := make([]point, 0)
		for _, cur := range q {
			//iterate all rays
			for _, d := range dirs {
				dx := cur.x + d[0]
				dy := cur.y + d[1]
				// follow the ray
				for dx >= 0 && dy >= 0 && dx < m && dy < n {
					// hit the wall
					if grid[dy][dx] == 0 {
						break
					}
					// hit same depth cell
					if dist[dy][dx] == depth-1 {
						break
					}
					if dist[dy][dx] == -1 {
						dist[dy][dx] = depth
						newQ = append(newQ, point{dx, dy})
					}
					dx += d[0]
					dy += d[1]
				}
			}
		}
		q = newQ
		depth++
	}
	if dist[end.y][end.x] == 0 {
		fprintln(-1)
		return
	}
	fprintln(dist[end.y][end.x])
}

type point struct {
	x int
	y int
}
