package proof

func pa(a, b *[]int) {
	*a = append(*a, (*b)[0])
	reverse(*a)
	*b = (*b)[1:]
}

func pb(a, b *[]int) {
	*b = append(*b, (*a)[0])
	reverse(*b)
	*a = (*a)[1:]
}

func sa(a []int) {
	a[0], a[1] = a[1], a[0]
}

func sb(b []int) {
	b[0], b[1] = b[1], b[0]
}

func ss(a, b []int) {
	sa(a)
	sb(b)
}

func ra(a []int) {
	for i := 0; i < len(a)-1; i++ {
		a[i], a[i+1] = a[i+1], a[i]
	}
}

func rb(b []int) {
	for i := 0; i < len(b)-1; i++ {
		b[i], b[i+1] = b[i+1], b[i]
	}
}

func rr(a, b []int) {
	ra(a)
	rb(b)
}

func rra(a []int) {
	// fmt.Println("DO: ", a)
	for i := len(a) - 1; i > 0; i-- {
		a[i], a[i-1] = a[i-1], a[i]
	}
	// fmt.Println("POSLE: ", a)
	// os.Exit(1)
}

func rrb(b []int) {
	for i := len(b) - 1; i > 0; i-- {
		b[i], b[i-1] = b[i-1], b[i]
	}
}

func rrr(a, b []int) {
	rra(a)
	rrb(b)
}

func reverse(a []int) {
	for i := len(a) - 1; i > 0; i-- {
		a[i], a[i-1] = a[i-1], a[i]
	}
}
