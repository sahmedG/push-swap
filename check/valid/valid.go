package valid

import "github.com/smslash/push-swap/check/handle"

func IsValid(s []string) {
	for i := 0; i < len(s); i++ {
		if !key(s[i]) {
			handle.Error()
		}
	}
}

func key(s string) bool {
	if s != "pa" && s != "pb" && s != "sa" && s != "sb" && s != "ss" && s != "ra" && s != "rb" && s != "rr" && s != "rra" && s != "rrb" && s != "rrr" {
		return false
	}
	return true
}
