package main

import "github.com/01-edu/z01"

func main() {
	var xrune rune = 'z'
	for xrune >= 'a' {
		z01.PrintRune(xrune)
		xrune--
	}
	z01.PrintRune(10)
}
