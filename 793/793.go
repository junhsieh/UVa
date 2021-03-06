// UVa 793 - Network Connections

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func unionFind(x int, f []int) int {
	if f[x] != x {
		f[x] = unionFind(f[x], f)
	}
	return f[x]
}

func main() {
	in, _ := os.Open("793.in")
	defer in.Close()
	out, _ := os.Create("793.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var a, b int
	var command byte
	var line string
	s.Scan()
	kase, _ := strconv.Atoi(s.Text())
	s.Scan()
	s.Text()
	for ; kase > 0 && s.Scan(); kase-- {
		num, _ := strconv.Atoi(s.Text())
		f := make([]int, num+1)
		for i := range f {
			f[i] = i
		}
		ok, ko := 0, 0
		for s.Scan() {
			if line = s.Text(); line == "" {
				break
			}
			fmt.Sscanf(line, "%c%d%d", &command, &a, &b)
			fa, fb := unionFind(a, f), unionFind(b, f)
			if command == 'c' {
				if fa != fb {
					f[fa] = fb
				}
			} else {
				if fa != fb {
					ko++
				} else {
					ok++
				}
			}
		}
		fmt.Fprintf(out, "%d,%d\n", ok, ko)
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
