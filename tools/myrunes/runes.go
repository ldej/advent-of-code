package myrunes

func Equal(a []rune, b []rune) bool {
	return string(a) == string(b)
}

func ToInt(a rune) int {
	out := int(a) - 96
	if out < 1 {
		return out + 58
	}
	return out
}
