package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	arg := os.Args[1]
	numbers := strings.Fields(arg)
	StackA := make(Stack, len(numbers))
	for i, numStr := range numbers {
		num, err := parseInt(numStr)
		if err != nil {
			fmt.Println("Error")
			return
		}
		StackA[i] = num
	}

	var StackB []int
	if CheckDup(StackA) {
		fmt.Println("Error")
		return
	} else if sort.IntsAreSorted(StackA) {
		return
	} else {
		if len(StackA) <= 3 {
			SmallStack(StackA)
		} else {
			LargeStack(StackA, StackB)
		}
	}
}
