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
	dict := map[rune]rune{'(': ')', '{': '}', '[': ']'}
	stack := []rune{}
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
			continue
		}
		if len(stack) == 0 || v != dict[stack[len(stack)-1]] {
			fprintln("no")
			return
		}
		stack = stack[:len(stack)-1]
	}
	if len(stack) > 0 {
		fprintln("no")
	} else {
		fprintln("yes")
	}
}
