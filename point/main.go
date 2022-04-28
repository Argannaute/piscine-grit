package main

import (
	"github.com/01-edu/z01"
)

var text string

func main() {
	text = "x = 42, y = 21"
	printStr(text)
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
