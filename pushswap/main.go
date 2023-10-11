package main

import (
	"os"

	"github.com/smslash/push-swap/pushswap/internal/algo"
	"github.com/smslash/push-swap/pushswap/internal/valid"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	s := os.Args[1]

	if !valid.IsSorted(valid.IsValid(s)) {
		a := valid.Parse(s)
		b := []int{}

		if len(a) < 3 {
			algo.TwoElements(a)
		} else if len(a) == 3 {
			algo.ThreeElements(a)
		} else {
			algo.LargeElements(a, b)
		}
	}
}
