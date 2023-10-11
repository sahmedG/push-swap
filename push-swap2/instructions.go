package main

func executeInstruction(a *Stack, b *Stack, instruction string) bool {
	switch instruction {
	case "sa":
		if sa(*a) {
			return true
		}
	case "sb":
		if sb(*b) {
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
		if ra(*a) {
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