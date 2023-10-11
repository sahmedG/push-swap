package proof

func Start(a []int, b []int, s []string) ([]int, []int) {
	// a := make([]int, len(arr))
	// copy(a, arr)
	// a := arr

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case "pa":
			pa(&a, &b)
		case "pb":
			pb(&a, &b)
		case "sa":
			sa(a)
		case "sb":
			sa(b)
		case "ss":
			ss(a, b)
		case "ra":
			ra(a)
		case "rb":
			rb(b)
		case "rr":
			rr(a, b)
		case "rra":
			rra(a)
		case "rrb":
			rrb(b)
		case "rrr":
			rrr(a, b)
		}
	}

	// if IsSorted(a) && len(b) == 0 {
	// 	fmt.Println("OK")
	// } else {
	// 	fmt.Println("KO")
	// }
	return a, b
}
