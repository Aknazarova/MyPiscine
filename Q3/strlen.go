package main

import (
	"fmt"
)

func StrLen(str string) int {
	count := 0
	for range str {
		count++
	}
	return count
}

func main() {
	str := "Hello World!"
	nb := StrLen(str)
	fmt.Println(nb)
}
