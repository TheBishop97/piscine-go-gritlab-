package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	for i, char := range runes {
		if char >= 'a' && char <= 'z' {
			runes[i] = char - 32
		}
	}
	return string(runes)
}
