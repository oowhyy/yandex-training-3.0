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
	for {
		var com string
		fscan(&com)
		switch com {
		case "push":
			var val int
			fscan(&val)
			stack.Push(val)
			fprintln("ok")
		case "pop":
			res := stack.Pop()
			fprintln(res)
		case "back":
			res := stack.Back()
			fprintln(res)
		case "size":
			fprintln(stack.len)
		case "clear":
			stack.Clear()
			fprintln("ok")
		case "exit":
			fprintln("bye")
			return
		}

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
