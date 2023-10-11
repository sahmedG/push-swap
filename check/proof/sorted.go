package proof

func IsSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j+1] < arr[j] {
				return false
			}
		}
	}
	return true
}
