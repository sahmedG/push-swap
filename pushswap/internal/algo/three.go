package algo

import (
	"fmt"
)

func ThreeElements(a []int) {
	s, m, l := compare(a)

	if m == a[0] && s == a[1] && l == a[2] {
		fmt.Println("sa")
	} else if l == a[0] && m == a[1] && s == a[2] {
		fmt.Println("sa")
		fmt.Println("rra")
	} else if l == a[0] && s == a[1] && m == a[2] {
		fmt.Println("ra")
	} else if s == a[0] && l == a[1] && m == a[2] {
		fmt.Println("sa")
		fmt.Println("ra")
	} else if m == a[0] && l == a[1] && s == a[2] {
		fmt.Println("rra")
	}

	a[0], a[1], a[2] = s, m, l
}

func compare(a []int) (int, int, int) {
	max := a[0]
	min := a[0]
	med := a[0]

	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}

		if a[i] < min {
			min = a[i]
		}
	}

	for i := 0; i < len(a); i++ {
		if a[i] != max && a[i] != min {
			med = a[i]
		}
	}

	return min, med, max
}
