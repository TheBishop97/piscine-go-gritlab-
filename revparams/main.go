package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		PrintStringWithNewLine("Fuck Off")
	}
	restArgs := os.Args[1:]
	for i := len(restArgs) - 1; i >= 0; i-- {
		PrintStringWithNewLine(restArgs[i])
	}
}

func PrintStringWithNewLine(mystring string) {
	for _, v := range mystring {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
