package valid

import (
	"strconv"
	"strings"

	"github.com/smslash/push-swap/pushswap/internal/handle"
)

func IsValid(nums string) []int {
	s := strings.Fields(nums)
	arr := make([]int, len(s))
	var err error
	for i := 0; i < len(s); i++ {
		arr[i], err = strconv.Atoi(string(s[i]))
		if err != nil {
			handle.Error()
		}
	}

	duplicateExist(arr)
	return arr
}
