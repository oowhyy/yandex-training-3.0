package main

import (
	"bufio"
	"fmt"
	"math"
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
	maxX := math.MinInt32
	maxY := math.MinInt32
	minX := math.MaxInt32
	minY := math.MaxInt32
	for i := 0; i < n; i++ {
		var x, y int
		fscan(&x, &y)
		if x > maxX {
			maxX = x
		}
		if x < minX {
			minX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}
	fprintln(minX, minY, maxX, maxY)
}
