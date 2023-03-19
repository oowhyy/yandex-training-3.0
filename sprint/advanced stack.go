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
	var n, t int
	fscan(&n)
	var inp string
	dict := map[string]int{}
	stack := []goods{}
	for i := 0; i < n; i++ {
		fscan(&inp)
		switch inp {
		case "add":
			fscan(&t)
			fscan(&inp)
			stack = append(stack, goods{inp, t})
			dict[inp] += t
		case "delete":
			fscan(&t)
			if t <= 0 {
				break
			}
			for {
				cur := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				dict[cur.name] -= cur.val
				if cur.val >= t {
					dict[cur.name] += cur.val - t
					stack = append(stack, goods{cur.name, cur.val - t})
					break
				}
				t -= cur.val
			}
		case "get":
			fscan(&inp)
			fprintln(dict[inp])
		}
	}
}

type goods struct {
	name string
	val  int
}
