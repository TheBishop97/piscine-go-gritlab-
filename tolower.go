package piscine

func ToLower(s string) string {
	Runes := []rune(s)
	for i, char := range Runes {
		if char >= 'A' && char <= 'Z' {
			Runes[i] = char + 32
		}
	}
	return string(Runes)
}
