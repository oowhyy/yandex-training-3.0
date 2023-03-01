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
	mp := map[rune]int64{}
	n := int64(len(s))
	for k, v := range s {
		k64 := int64(k)
		mp[v] += (n - k64) * (k64 + 1)
	}
	for i := rune('a'); i <= rune('z'); i++ {
		if v, ok := mp[i]; ok {
			fprintln(string(i)+":", v)
		}
	}
}
