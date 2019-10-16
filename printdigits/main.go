package main

import "github.com/01-edu/z01"

func main() {
	var xrune rune = '0'
	for xrune <= '9' {
		z01.PrintRune(xrune)
		xrune++
	}
	z01.PrintRune(10)
}
