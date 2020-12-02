package tools

func StringIndex(s string, i int) string {
	if i < 0 {
		return ""
	}
	if i > len(s) {
		return ""
	}
	return string([]rune(s)[i])
}

func StringCompareXOR(a string, b string, compare string) bool {
	return (a == compare || b == compare) && !(a == compare && b == compare)
}