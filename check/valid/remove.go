package valid

func RemoveNewLine(s string) string {
	res := ""

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i] == '\\' && s[i+1] == 'n' {
			i++
			res += " "
		} else {
			res += string(s[i])
		}
	}

	return res
}
