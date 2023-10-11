package main


func isSortedAscending(s Stack) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}

func CheckDup(a Stack) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] == a[i+1] {
			return true
		}
	}
	return false
}