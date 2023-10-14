package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// // Stack operations
// // func push(stack []int, num int) []int {
// // 	return append(stack, num)
// // }

// // func pop(stack []int) ([]int, int) {
// // 	if len(stack) == 0 {
// // 		return nil, 0
// // 	}
// // 	return stack[:len(stack)-1], stack[len(stack)-1]
// // }

// func sa(stackA []int) []int {
// 	if len(stackA) >= 2 {
// 		stackA[len(stackA)-1], stackA[len(stackA)-2] = stackA[len(stackA)-2], stackA[len(stackA)-1]
// 	}
// 	return stackA
// }

// func ra(stackA []int) []int {
// 	if len(stackA) > 1 {
// 		stackA = append(stackA[1:], stackA[0])
// 	}
// 	return stackA
// }

// func rra(stackA []int) []int {
// 	if len(stackA) > 1 {
// 		stackA = append([]int{stackA[len(stackA)-1]}, stackA[:len(stackA)-1]...)
// 	}
// 	return stackA
// }

// // func main() {
// // 	args := os.Args[1:]
// // 	if len(args) == 0 {
// // 		return
// // 	}

// // 	var stackA []int
// // 	for _, arg := range args {
// // 		num, err := strconv.Atoi(arg)
// // 		if err != nil {
// // 			fmt.Fprintln(os.Stderr, "Error")
// // 			return
// // 		}
// // 		stackA = append(stackA, num)
// // 	}

// // 	reader := bufio.NewReader(os.Stdin)
// // 	for {
// // 		instruction, err := reader.ReadString('\n')
// // 		if err != nil {
// // 			break
// // 		}

// // 		instruction = strings.TrimSpace(instruction)

// // 		switch instruction {
// // 		case "sa":
// // 			stackA = sa(stackA)
// // 		case "ra":
// // 			stackA = ra(stackA)
// // 		case "rra":
// // 			stackA = rra(stackA)
// // 		// Add other cases for remaining instructions
// // 		default:
// // 			fmt.Fprintln(os.Stderr, "Error")
// // 			return
// // 		}
// // 	}

// // 	// Check if stackA is sorted
// // 	isSorted := true
// // 	for i := 1; i < len(stackA); i++ {
// // 		if stackA[i] < stackA[i-1] {
// // 			isSorted = false
// // 			break
// // 		}
// // 	}

// // 	if isSorted && len(stackA) > 0 {
// // 		fmt.Println("OK")
// // 	} else {
// // 		fmt.Println("KO")
// // 	}
// // }

type Stack []int

// Instructions

func sa(a Stack) {
	if len(a) >= 2 {
		a[0], a[1] = a[1], a[0]
	}
}

func sb(b Stack) {
	if len(b) >= 2 {
		b[0], b[1] = b[1], b[0]
	}
}

func ss(a Stack, b Stack) {
	sa(a)
	sb(b)
}

func pa(a *Stack, b *Stack) {
	if len(*b) > 0 {
		*a = append([]int{(*b)[0]}, *a...)
		*b = (*b)[1:]
	}
}

func pb(a *Stack, b *Stack) {
	if len(*a) > 0 {
		*b = append([]int{(*a)[0]}, *b...)
		*a = (*a)[1:]
	}
}

func ra(a Stack) {
	if len(a) >= 2 {
		first := a[0]
		copy(a, a[1:])
		a[len(a)-1] = first
	}
}

func rb(b Stack) {
	if len(b) >= 2 {
		first := b[0]
		copy(b, b[1:])
		b[len(b)-1] = first
	}
}

func rr(a Stack, b Stack) {
	ra(a)
	rb(b)
}

func rra(a Stack) {
	if len(a) >= 2 {
		last := a[len(a)-1]
		copy(a[1:], a[:len(a)-1])
		a[0] = last
	}
}

func rrb(b Stack) {
	if len(b) >= 2 {
		last := b[len(b)-1]
		copy(b[1:], b[:len(b)-1])
		b[0] = last
	}
}

func rrr(a Stack, b Stack) {
	rra(a)
	rrb(b)
}

// Sorting algorithm
func sortStack(a Stack, b Stack) {
	if len(a) <= 3 {
		sortSmallStack(a)
		return
	}

	pivot := findPivot(a)
	for len(a) > 0 {
		if a[0] <= pivot {
			pb(&a, &b)
		} else {
			ra(a)
		}
	}

	for len(b) > 0 {
		pa(&a, &b)
	}
}

func sortSmallStack(a Stack) {
	if len(a) == 2 && a[0] > a[1] {
		sa(a)
	}
	if len(a) == 3 {
		if a[0] > a[1] && a[0] < a[2] {
			sa(a)
			rra(a)
		} else if a[0] > a[1] && a[1] > a[2] {
			sa(a)
			ra(a)
		} else if a[0] > a[1] {
			ra(a)
		} else if a[1] > a[2] {
			rra(a)
		}
	}
}

func findPivot(a Stack) int {
	low := a[0]
	high := a[0]

	for _, num := range a {
		if num < low {
			low = num
		}
		if num > high {
			high = num
		}
	}

	return (low + high) / 2
}

// Parsing utility function
func parseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// Execute an instruction on the stacks
func executeInstruction(a *Stack, b *Stack, instruction string) {
	switch instruction {
	case "sa":
		sa(*a)
	case "sb":
		sb(*b)
	case "ss":
		ss(*a, *b)
	case "pa":
		pa(a, b)
	case "pb":
		pb(a, b)
	case "ra":
		ra(*a)
	case "rb":
		rb(*b)
	case "rr":
		rr(*a, *b)
	case "rra":
		rra(*a)
	case "rrb":
		rrb(*b)
	case "rrr":
		rrr(*a, *b)
	}
}

// Check if a stack is sorted
func isSorted(s Stack) bool {
	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			return false
		}
	}
	return true
}
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

	stackB := make(Stack, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		instruction := scanner.Text()
		executeInstruction(&stackA, &stackB, instruction)
	}

	// Check if stackA is sorted and stackB is empty
	if isSorted(stackA) && len(stackB) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}