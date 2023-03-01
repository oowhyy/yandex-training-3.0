package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	if len(s) == 0 {
		fprintln("")
	}
	res := strings.Builder{}
	cur := s[0]
	count := 1
	for i := 1; i < len(s); i++ {
		if cur == s[i] {
			count++
		} else {
			res.WriteString(pack(cur, count))
			cur = s[i]
			count = 1
		}
	}
	res.WriteString(pack(cur, count))
	fprintln(res.String())
}

func pack(x byte, n int) string {
	if n > 1 {
		return string(x) + fmt.Sprint(n)
	}
	return string(x)
}
