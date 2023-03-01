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
	var k int
	fscan(&k)
	var s string
	fscan(&s)
	n := len(s)
	best := 0
	for i := 'a'; i <= 'z'; i++ {
		stars := k
		p := 0
		for q, r := range s {
			for p < n && (rune(s[p]) == i || stars > 0) {
				if rune(s[p]) != i {
					stars--
				}
				p++
			}
			best = max(best, p-q)
			if r != i {
				stars++
			}
		}
	}
	fprintln(best)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
