package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, v := range []rune(os.Args[0])[2:] {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
