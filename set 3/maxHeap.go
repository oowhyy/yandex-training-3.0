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
	heap := newHeap()
	var n, cur int
	fscan(&n)
	for i := 0; i < n; i++ {
		fscan(&cur)
		if cur == 0 {
			fscan(&cur)
			heap.Push(cur)
		} else {
			fprintln(heap.Pop())
		}
	}
}

type MaxHeap struct {
	arr []int
}

func newHeap() MaxHeap {
	return MaxHeap{}
}

func (this *MaxHeap) Push(x int) {
	this.arr = append(this.arr, x)
	j := len(this.arr) - 1
	for {
		i := (j - 1) / 2 // parent
		if i == j || this.arr[j] <= this.arr[i] {
			break
		}
		this.arr[i], this.arr[j] = this.arr[j], this.arr[i]
		j = i
	}

}

func (this *MaxHeap) Pop() int {
	n := len(this.arr) - 1
	this.arr[n], this.arr[0] = this.arr[0], this.arr[n]
	i := 0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && this.arr[j2] > this.arr[j1] {
			j = j2 // = 2*i + 2  // right child
		}
		if this.arr[j] <= this.arr[i] {
			break
		}
		this.arr[i], this.arr[j] = this.arr[j], this.arr[i]
		i = j
	}
	res := this.arr[n]
	this.arr = this.arr[:n]
	return res
}
