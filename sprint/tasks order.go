package main

import (
	"bufio"
	"container/heap"
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
	var n, w, a, b int
	fscan(&n, &w)
	starts := &Heap1{}
	ends := &Heap2{}

	for i := 0; i < n; i++ {
		fscan(&a, &b)

		heap.Push(starts, task{i + 1, a, b + a, -1})
	}
	tasks := [][]int{}
	tasks = append(tasks, []int{})
	res := 0
	for i := 0; i < n; i++ {
		cur := heap.Pop(starts).(task)
		//fprintln(cur)
		if ends.Len() == 0 {
			heap.Push(ends, task{cur.id, cur.start, cur.end, 0})
			tasks[0] = append(tasks[0], cur.id)
			continue
		}
		if (*ends)[0].end > cur.start {
			res++
			heap.Push(ends, task{cur.id, cur.start, cur.end, res})
			tasks = append(tasks, []int{})
			tasks[res] = append(tasks[res], cur.id)
		} else {
			empNum := heap.Pop(ends).(task).employee
			heap.Push(ends, task{cur.id, cur.start, cur.end, empNum})
			tasks[empNum] = append(tasks[empNum], cur.id)
		}

	}
	//fprintln(tasks)
	fprintln(res + 1)
	for _, v := range tasks {
		for _, nn := range v {
			fmt.Fprint(out, nn, " ")
		}
	}
}

type task struct {
	id       int
	start    int
	end      int
	employee int
}

type Heap1 []task

func (h Heap1) Len() int           { return len(h) }
func (h Heap1) Less(i, j int) bool { return h[i].start < h[j].start }
func (h Heap1) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap1) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(task))
}
func (h *Heap1) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Heap2 []task

func (h Heap2) Len() int           { return len(h) }
func (h Heap2) Less(i, j int) bool { return h[i].end < h[j].end }
func (h Heap2) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap2) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(task))
}
func (h *Heap2) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
