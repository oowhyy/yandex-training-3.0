package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var a, b, c string
	fscan(&a, &b, &c)
	t1 := getTime(a)
	t2 := getTime(b)
	t3 := getTime(c)
	d := timeDiff(t1, t3)
	d = timeDiv2(d)
	d = timeSum(t2, d)
	printTime(d)
}

func timeDiv2(time []int) []int {
	res := make([]int, 3)
	q, rem := 0, 0
	for i := 0; i <= 2; i++ {
		cur := time[i]
		if rem == 1 {
			cur += 60
		}
		q = cur / 2
		rem = cur % 2
		res[i] = q
	}
	if rem == 1 {
		return timeSum(res, []int{0, 0, 1})
	}
	return res
}

func timeSum(start, end []int) []int {
	carry := 0
	res := make([]int, 3)
	for i := 2; i >= 0; i-- {
		d := end[i] + start[i] + carry
		if i == 0 {
			if d >= 24 {
				d -= 24
				carry = 1
			} else {
				carry = 0
			}
		} else {
			if d >= 60 {
				d -= 60
				carry = 1
			} else {
				carry = 0
			}

		}
		res[i] = d
	}
	return res
}

func timeDiff(start, end []int) []int {
	carry := 0
	res := make([]int, 3)
	for i := 2; i >= 0; i-- {
		d := end[i] - start[i] - carry
		if d < 0 {
			carry = 1
			if i == 0 {
				d += 24
			} else {
				d += 60
			}
		} else {
			carry = 0
		}
		res[i] = d
	}
	return res
}

func getTime(s string) []int {
	hh, _ := strconv.Atoi(s[:2])
	mm, _ := strconv.Atoi(s[3:5])
	ss, _ := strconv.Atoi(s[6:])
	return []int{hh, mm, ss}
}

func printTime(t []int) {
	fmt.Fprintf(out, "%02d:%02d:%02d\n", t[0], t[1], t[2])
}
