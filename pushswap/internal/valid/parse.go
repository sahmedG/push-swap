package valid

import (
	"strconv"
	"strings"
)

func Parse(data string) []int {
	s := strings.Fields(data)
	arr := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		arr[i], _ = strconv.Atoi(string(s[i]))
	}

	return arr
}
