package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}
	args := os.Args[1:]
	isUpper := args[0] == "--upper"
	if isUpper {
		if len(args) < 2 {
			os.Exit(0)
		}
		args = args[1:]
	}
	if isUpper {
		for _, v := range args {
			v := makestringtoint(v)
			if v > 26 || v < 0 {
				z01.PrintRune(' ')
				continue
			}
			z01.PrintRune('A' + rune(v) - 1)
		}
	} else {
		for _, v := range args {
			v := makestringtoint(v)
			if v > 26 || v < 0 {
				z01.PrintRune(' ')
				continue
			}
			z01.PrintRune('a' + rune(v) - 1)
		}
	}
	z01.PrintRune('\n')
}

func makestringtoint(mystring string) int {
	var n int
	isneg := []rune(mystring)[0] == ' '
	for _, v := range mystring {
		n = n*10 + int(v-'0')
	}
	if isneg {
		return -n
	}
	return n
}

func PrintStringNow(mystring string) {
	for _, v := range mystring {
		z01.PrintRune(v)
	}
}
