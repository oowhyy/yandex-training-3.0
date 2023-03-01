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
	stack := newStack()
	var n, cur int
	fscan(&n)
	p := 1
	for i := 0; i < n; i++ {
		fscan(&cur)
		if cur != p {
			stack.Push(cur)
			continue
		}
		p++
		for !stack.IsEmpty() && stack.Back() == p {
			stack.Pop()
			p++
		}
	}
	if p == n+1 {
		fprintln("YES")
	} else {
		fprintln("NO")
	}
}

type Stack struct {
	len  int
	head *node
}

type node struct {
	value int
	next  *node
}

func newStack() *Stack {
	return &Stack{0, nil}
}

func (this *Stack) Push(val int) {
	n := &node{val, this.head}
	this.head = n
	this.len++
}

func (this *Stack) IsEmpty() bool {
	return this.len == 0
}

func (this *Stack) Pop() interface{} {
	if this.IsEmpty() {
		return "error"
	}
	n := this.head
	this.head = n.next
	this.len--
	return n.value
}

func (this *Stack) Back() interface{} {
	if this.IsEmpty() {
		return "error"
	}
	return this.head.value
}

func (this *Stack) Clear() {
	*this = *newStack()
}
