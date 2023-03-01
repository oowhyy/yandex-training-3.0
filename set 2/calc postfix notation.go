package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
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
	line, _ := in.ReadString('\n')
	stack := newStack()
	for _, v := range line {
		if v == ' ' {
			continue
		}
		cur := v
		//fprintln(string(cur))
		if unicode.IsDigit(cur) {
			stack.Push(int(cur - '0'))
			continue
		}
		switch cur {
		case '*':
			first := stack.Pop()
			second := stack.Pop()
			stack.Push(first * second)
		case '-':
			first := stack.Pop()
			second := stack.Pop()
			stack.Push(second - first)
		case '+':
			first := stack.Pop()
			second := stack.Pop()
			stack.Push(first + second)

		}
	}
	fprintln(stack.Pop())

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

func (this *Stack) Pop() int {
	n := this.head
	this.head = n.next
	this.len--
	return n.value
}

func (this *Stack) Back() int {

	return this.head.value
}

func (this *Stack) Clear() {
	*this = *newStack()
}
