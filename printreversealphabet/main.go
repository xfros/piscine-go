package main

import "github.com/01-edu/z01"

func main() {
	i := 25
	for i >= 0 {
		var aRune string = "abcdefghijklmnopqrstuvwxyz"
		z01.PrintRune(rune(aRune[i]))
		i--
	}
	z01.PrintRune(10)
}