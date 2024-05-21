package piscine

func LastRune(s string) rune {
	length := len(s)
	runes := []rune(s)
	if length > 0 {
		return runes[length-1]
	}
	return 0
}
