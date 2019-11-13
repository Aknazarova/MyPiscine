package main
import "fmt"
func StrLen(str string) int {
	a := []rune(str)
	count := 0
	for range a {
		count++
	}
	return count
}

func main() {
	str := "Hello World!"
	nb := StrLen(str)
	fmt.Println(nb)
}
