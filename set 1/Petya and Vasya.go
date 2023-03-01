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
	var n, k, row, pos int
	fscan(&n, &k, &row, &pos)
	absPos := toAbs(row, pos)
	absAfter := absPos + k
	absBefore := absPos - k
	if absAfter <= n {
		rowAfter, posAfter := fromAbs(absAfter)
		if absBefore > 0 {
			rowBefore, posBefore := fromAbs(absBefore)
			if row-rowBefore < rowAfter-row {
				fprintln(rowBefore, posBefore)
				return
			} else {
				fprintln(rowAfter, posAfter)
				return
			}
		}
		fprintln(rowAfter, posAfter)
		return
	}
	if absBefore > 0 {
		rowBefore, posBefore := fromAbs(absBefore)
		fprintln(rowBefore, posBefore)
		return
	}
	fprintln(-1)

}

func toAbs(row, pos int) int {
	res := row * 2
	if pos == 1 {
		res -= 1
	}
	return res
}

func fromAbs(x int) (row, pos int) {
	row = (x + 1) / 2
	if x%2 == 0 {
		pos = 2
		return
	}
	pos = 1
	return
}

//2 1
//3 2
//1 3
//2 1
