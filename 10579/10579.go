// UVa 10579 - Fibonacci Numbers

package main

import (
	"fmt"
	"math/big"
	"os"
)

var f []big.Int

func prepare() {
	f = make([]big.Int, 2)
	f[0].SetInt64(0)
	f[1].SetInt64(1)

	var s string
	for i := 2;; i++ {
		var tmp big.Int
		tmp.Add(&f[i - 2], &f[i - 1])
		s = fmt.Sprintf("%v", &tmp)
		if len(s) > 1000 {
			break;
		}
		f = append(f, tmp)
	}
}

func do() {
	in, _ := os.Open("10579.in")
	out, _ := os.Create("10579.out")
	var n int
	for {
		_, err := fmt.Fscanf(in, "%d", &n)
		if err != nil {
			break;
		}
		fmt.Fprintln(out, &f[n])
	}
	in.Close()
	out.Close()
}

func main() {
	prepare()
	do()
}