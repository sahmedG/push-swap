package main

import (
	"fmt"
	"os"
	"strings"
)

type Stack []int

func main() {
	if len(os.Args) < 2 {
		return
	}

	arg := os.Args[1]
	numbers := strings.Fields(arg)
	stackA := make(Stack, len(numbers))

	for i, numStr := range numbers {
		num, err := parseInt(numStr)
		if err != nil {
			fmt.Println("Error")
			return
		}
		stackA[i] = num
	}

	var stackB Stack
	Instructions := []string{"pb", "sa", "sb", "ss", "pa", "ra", "rb", "rr", "rra", "rrb", "rrr"}

	if CheckDup(stackA) {
		fmt.Println("Error")
		return
	}

	if !isSortedAscending(stackA) {
		if len(stackA) == 3 {
			if stackA[0] > stackA[1]{
				sa(stackA)
				fmt.Println("sa")
				
			}
		}
	}


	if isSortedAscending(stackA) {
		return
	} else {
		for _, instruction := range Instructions {
			executeInstruction(&stackA, &stackB, instruction)
			if isSortedAscending(stackA) && len(stackB) == 0 {
				break
			}
		}
	}
}
