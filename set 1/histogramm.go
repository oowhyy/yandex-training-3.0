package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
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

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func solve() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	mp := map[rune]int{}
	max := 0
	for scanner.Scan() {
		for _, v := range scanner.Text() {
			if v != ' ' {
				mp[v]++
				if mp[v] > max {
					max = mp[v]
				}
			}
		}
	}
	i := 0
	keys := make([]rune, len(mp))
	for k := range mp {
		keys[i] = k
		i++
	}

	sort.Sort(sortRunes(keys))
	for ; max > 0; max-- {
		for _, v := range keys {
			if mp[v] >= max {
				fmt.Fprint(out, "#")
			} else {
				fmt.Fprint(out, " ")
			}
		}
		fmt.Fprintf(out, "\n")
	}
	keyStrings := make([]string, len(keys))
	for k, v := range keys {
		keyStrings[k] = string(v)
	}
	fprintln(strings.Join(keyStrings, ""))
}
