package main

import "github.com/01-edu/z01"

func PrintStr(str string) {
	//var i rune
	for _, i := range str {
		z01.PrintRune(i)
	}
}
func main() {
	str := "Hello World!"
	PrintStr(str)
}
