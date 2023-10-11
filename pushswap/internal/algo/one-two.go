package algo

import (
	"fmt"
	"os"
)

func TwoElements(a []int) {
	if len(a) == 2 {
		if a[0] > a[1] {
			fmt.Println("rra")
		} else {
			os.Exit(0)
		}
	} else if len(a) == 1 {
		os.Exit(0)
	}
}
