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
	deque := newDeque[int]()
	var inp string
	for {
		fscan(&inp)
		switch inp {
		case "push_back":
			var x int
			fscan(&x)
			deque.PushBack(x)
			fprintln("ok")
		case "push_front":
			var x int
			fscan(&x)
			deque.PushFront(x)
			fprintln("ok")
		case "pop_front":
			if deque.Size() == 0 {
				fprintln("error")
			} else {
				fprintln(deque.PopFront())
			}
		case "pop_back":
			if deque.Size() == 0 {
				fprintln("error")
			} else {
				fprintln(deque.PopBack())
			}
		case "back":
			if deque.Size() == 0 {
				fprintln("error")
			} else {
				fprintln(deque.Back())
			}
		case "front":
			if deque.Size() == 0 {
				fprintln("error")
			} else {
				fprintln(deque.Front())
			}
		case "size":
			fprintln(deque.Size())
		case "clear":
			deque.Clear()
			fprintln("ok")
		case "exit":
			fprintln("bye")
			return
		}

	}
}

type Deque[T any] struct {
	arr    []T
	head   int
	tail   int
	count  int
	minCap int
}

func newDeque[T any]() Deque[T] {
	return Deque[T]{arr: make([]T, 16), minCap: 16}
}

func (this *Deque[T]) Size() int {
	if this == nil {
		return 0
	}
	return this.count
}

func (this *Deque[T]) Clear() {
	this.arr = make([]T, 16)
	this.head = 0
	this.tail = 0
	this.count = 0
}

func (this *Deque[T]) PushBack(x T) {
	this.tryGrow()
	// push, then calculate new tail pos

	this.arr[this.tail] = x
	this.tail = this.next(this.tail)
	this.count++
}
func (this *Deque[T]) PushFront(x T) {
	this.tryGrow()
	// calculate new head pos, then push
	this.head = this.prev(this.head)
	this.arr[this.head] = x

	this.count++
}
func (this *Deque[T]) PopBack() T {
	if this.Size() == 0 {
		panic("pop on empty deque")
	}
	// calculate new tail pos, then pop
	this.tail = this.prev(this.tail)
	res := this.arr[this.tail]

	var zero T
	this.arr[this.tail] = zero
	this.count--
	this.tryShrink()
	return res
}

func (this *Deque[T]) PopFront() T {
	if this.Size() == 0 {
		panic("pop on empty deque")
	}
	// pop, then calculate new head pos
	res := this.arr[this.head]
	var zero T
	this.arr[this.head] = zero

	this.head = this.next(this.head)
	this.count--
	this.tryShrink()
	return res
}

func (this *Deque[T]) Front() T {
	if this.Size() == 0 {
		panic("front on empty deque")
	}
	return this.arr[this.head]
}

func (this *Deque[T]) Back() T {
	if this.Size() == 0 {
		panic("back on empty deque")
	}
	return this.arr[this.prev(this.tail)]
}

// grow if no more space (count == len)
func (this *Deque[T]) tryGrow() {
	if this.count != len(this.arr) {
		return
	}
	this.resize()

}

// shrink if count is exactly quarter of total length
func (this *Deque[T]) tryShrink() {
	if len(this.arr) > this.minCap && (this.count*4 == len(this.arr)) {
		this.resize()
	}
}

func (this *Deque[T]) next(x int) int {
	return (x + 1) % len(this.arr)
}
func (this *Deque[T]) prev(x int) int {
	return (x - 1 + len(this.arr)) % len(this.arr)
}

func (this *Deque[T]) resize() {
	newArr := make([]T, this.count*2)
	if this.tail > this.head {
		copy(newArr, this.arr[this.head:this.tail])
	} else {
		n := copy(newArr, this.arr[this.head:]) // length of this.arr[this.head:] // head is at index 0
		copy(newArr[n:], this.arr[:this.tail])  // copy tail to the end of newArr
	}
	this.head = 0
	this.tail = this.count
	this.arr = newArr
}
