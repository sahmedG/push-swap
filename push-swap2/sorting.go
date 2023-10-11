package main

import (
	"fmt"
	"strconv"
)

func sa(a Stack) bool {
	if len(a) >= 2 && a[0] > a[1] {
		a[0], a[1] = a[1], a[0]
		fmt.Println("sa", a)
		sa(a)
	}
	return true
}

func sb(b Stack) bool {
	if len(b) >= 2 && b[0] < b[1] {
		b[0], b[1] = b[1], b[0]
		fmt.Println("sb", b)
		sb(b)
	}
	return true
}

func ss(a Stack, b Stack) bool {
	if sa(a) || sb(b) {
		fmt.Println("ss", a, b)
	}
	return true
}

func pa(a *Stack, b *Stack) bool {
	if len(*b) > 0 && ((*a)[0] > (*b)[0]) {
		*a = append([]int{(*b)[0]}, *a...)
		*b = (*b)[1:]
		fmt.Println("pa", *a, *b)
		pa(a, b)
	}
	return true
}

func pb(a *Stack, b *Stack) bool {
	if len(*b) == 0 {
		*b = append([]int{(*a)[0]}, *b...)
		*a = (*a)[1:]
		fmt.Println("pb", *a, *b)
		pb(a, b)
	} else if len(*a) > 0 && ((*a)[0] < (*b)[0]) {
		*b = append([]int{(*a)[0]}, *b...)
		*a = (*a)[1:]
		fmt.Println("pb", *a, *b)
	}
	return true
}

func ra(a Stack) bool {
	if len(a) >= 2 {
		first := a[0]
		copy(a, a[1:])
		a[len(a)-1] = first
		fmt.Println("ra", a)
		rra(a)
	}
	return true
}

func rb(b Stack) bool {
	if len(b) >= 2 {
		first := b[0]
		copy(b, b[1:])
		b[len(b)-1] = first
		fmt.Println("rb", b)
		rb(b)
	}
	return true
}

func rr(a Stack, b Stack) bool {
	if ra(a) {
		return true
	}
	if rb(b) {
		return true
	}
	return true
}

func rra(a Stack) bool {
	if len(a) >= 2 {
		last := a[len(a)-1]
		copy(a[1:], a[:len(a)-1])
		a[0] = last
		fmt.Println("rra", a)
		rra(a)
	}
	return true
}

func rrb(b Stack) bool {
	if len(b) >= 2 {
		last := b[len(b)-1]
		copy(b[1:], b[:len(b)-1])
		b[0] = last
		fmt.Println("rrb", b)
		rra(b)
	}
	return true
}

func rrr(a Stack, b Stack) bool {
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
