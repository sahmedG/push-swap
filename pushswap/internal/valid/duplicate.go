package valid

import (
	"github.com/smslash/push-swap/pushswap/internal/handle"
)

func duplicateExist(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				handle.Error()
			}
		}
	}
}
