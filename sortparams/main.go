package main

import (
	"os"

	"github.com/01-edu/z01"
)

func bubleSort(arr []string) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	args := os.Args[1:]
	bubleSort(args)
	for _, v := range args {
		PrintStringWithNewLine(v)
	}
}

func PrintStringWithNewLine(mystring string) {
	for _, v := range mystring {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
