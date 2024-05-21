package piscine

func Split(s, sep string) []string {
	for i := 0; i < len(s); i++ {
		if index(s, sep) != -1 {
			s = s[0:index(s, sep)] + " " + s[index(s, sep)+len(sep):]
		}
	}
	return splitWhiteSpaces(s)
}

func splitWhiteSpaces(s string) []string {
	a := make([]string, 0)
	b := make([]byte, 0)
	for _, i := range s {
		if i == ' ' || i == '\t' || i == '\n' {
			if len(b) != 0 {
				a = append(a, string(b))
				b = make([]byte, 0)
			}
		} else {
			b = append(b, byte(i))
		}
	}
	if len(b) != 0 {
		a = append(a, string(b))
	}
	return a
}

func index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	if len(s) == 0 {
		return -1
	}
	for i := 0; i < len(s); i++ {
		if s[i] == toFind[0] {
			if len(toFind) == 1 {
				return i
			}
			if len(s[i:]) >= len(toFind) {
				if s[i:i+len(toFind)] == toFind {
					return i
				}
			}
		}
	}
	return -1
}
