package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		PrintStringWithNewLine("Fuck Off")
	}
	for _, v := range os.Args[1:] {
		PrintStringWithNewLine(v)
	}
}

func PrintStringWithNewLine(mystring string) {
	for _, v := range mystring {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
