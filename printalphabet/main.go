package main

import "github.com/01-edu/z01"
//import "fmt"

func main() {

	

	for i:= 0; i <= 25; i++ {

		//fmt.Println(i)
		
		var aRune string = "abcdefghijklmnopqrstuvwxyz"

		z01.PrintRune(rune(aRune[i]))
	}
	
}
