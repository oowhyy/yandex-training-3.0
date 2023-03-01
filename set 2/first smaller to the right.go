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
	res := make([]int, n)
	for i := 0; i < n; i++ {
		fscan(&cur)
		for !stack.IsEmpty() && stack.Back().first > cur {
			pop := stack.Pop()
			res[pop.second] = i
		}
		stack.Push(pair{cur, i})
	}
	for !stack.IsEmpty() {
		pop := stack.Pop()
		res[pop.second] = -1
	}
	for _, v := range res {
		fmt.Fprint(out, v, " ")
	}

}

type pair struct {
	first  int
	second int
}

type Stack struct {
	len  int
	head *node
}

type node struct {
	value pair
	next  *node
}

func newStack() *Stack {
	return &Stack{0, nil}
}

func (this *Stack) Push(val pair) {
	n := &node{val, this.head}
	this.head = n
	this.len++
}

func (this *Stack) IsEmpty() bool {
	return this.len == 0
}

func (this *Stack) Pop() pair {
	n := this.head
	this.head = n.next
	this.len--
	return n.value
}

func (this *Stack) Back() pair {

	return this.head.value
}

func (this *Stack) Clear() {
	*this = *newStack()
}
