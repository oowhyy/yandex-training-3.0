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
	var s string
	fscan(&s)
	m := map[rune]int{}
	for _, v := range s {
		m[v]++
	}
	var res rune
	count := 0
	for k, v := range m {
		if v > count {
			count = v
			res = k
		}
	}
	fprintln(string(res))
}
