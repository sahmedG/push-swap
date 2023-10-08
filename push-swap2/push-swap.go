package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Stack represents a stack of integers.
type Stack []int

// Instructions

func sa(a Stack) bool{
	if len(a) >= 2 {
		a[0], a[1] = a[1], a[0]
		return true
	}
	return false
}

func sb(b Stack) bool{
	if len(b) >= 2 {
		b[0], b[1] = b[1], b[0]
		return true
	}
	return false
}

func ss(a Stack, b Stack)bool {
	if sa(a) {
		return true
	}
	if sb(b){
		return true
	}
	return false
}

func pa(a *Stack, b *Stack) bool{
	if len(*b) > 0 {
		*a = append([]int{(*b)[0]}, *a...)
		*b = (*b)[1:]
		return true
	}
	return false
}

func pb(a *Stack, b *Stack) bool{
	if len(*b) == 0 {
		*b = append([]int{(*a)[0]}, *b...)
		*a = (*a)[1:]
		return true
	} else if len(*a) > 0 && ((*a)[0] < (*b)[0]){
		*b = append([]int{(*a)[0]}, *b...)
		*a = (*a)[1:]
		return true
	}
	return false
}

func ra(a Stack)bool {
	if len(a) >= 2 {
		first := a[0]
		copy(a, a[1:])
		a[len(a)-1] = first
		return true
	}
	return false
}

func rb(b Stack) bool{
	if len(b) >= 2 {
		first := b[0]
		copy(b, b[1:])
		b[len(b)-1] = first
		return true
	}
	return false
}

func rr(a Stack, b Stack)bool {
	if ra(a) {
		return true
	}
	if rb(b) {
		return true
	}
	return false
}

func rra(a Stack)bool {
	if len(a) >= 2 {
		last := a[len(a)-1]
		copy(a[1:], a[:len(a)-1])
		a[0] = last
		return true
	}
	return false
}

func rrb(b Stack) bool{
	if len(b) >= 2 {
		last := b[len(b)-1]
		copy(b[1:], b[:len(b)-1])
		b[0] = last
		return true
	}
	return false
}

func rrr(a Stack, b Stack) bool{
	if rra(a) {
		return true
	}
	if rrb(b) {
		return true
	}
	return false
}

// Parsing utility function
func parseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// Execute an instruction on the stacks
func executeInstruction(a *Stack, b *Stack, instruction string) bool{
	switch instruction {
	case "sa":
		if sa(*a) {
			return true
		}
	case "sb":
		if sb(*b){
			return true
		}
	case "ss":
		if ss(*a, *b) {
			return true
		}
	case "pa":
		if pa(a, b) {
			return true
		}
	case "pb":
		if pb(a, b) {
			return true
		}
	case "ra":
		if ra(*a){
			return true
		}
	case "rb":
		if rb(*b) {
			return true
		}
	case "rr":
		if rr(*a, *b) {
			return true
		}
	case "rra":
		if rra(*a) {
			return true
		}
	case "rrb":
		if rrb(*b) {
			return true
		}
	case "rrr":
		if rrr(*a, *b) {
			return true
		}
	}
	return false
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

	var stackB Stack
	Instructions := []string{"pb","sa", "sb", "ss", "pa",  "ra", "rb", "rr", "rra", "rrb", "rrr"}
	// Output the instructions and the final state of stackA
	loop:
	for _, instruction := range Instructions {
		if executeInstruction(&stackA, &stackB, instruction) {
			goto loop
		}
		fmt.Println(instruction, stackA, stackB)
	}
	
}

func isSortedAscending(s Stack) bool {
	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			return false
		}
	}
	return true
}