package tools

func MapAll(m map[string]bool, b bool) bool {
	for _, v := range m {
		if v != b {
			return false
		}
	}
	return true
}
