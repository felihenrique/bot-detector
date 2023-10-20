package utils

func BoolToInt(val bool) int {
	conv := 0
	if val {
		conv = 1
	}
	return conv
}
