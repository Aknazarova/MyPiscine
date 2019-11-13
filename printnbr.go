package main
import "github.com/01-edu/z01"

func PrintNbr(num int) {
	if num == 0 {
		z01.PrintRune('0')
	} else if num > 0 {
		if num >= 10 {
			PrintNbr(num / 10)
		}
		result := '0'
		for i := 0; i < num%10; i++ {
			result++
		}
		z01.PrintRune(result)
	} else {
		z01.PrintRune('-')
		PrintNbr((-1) * num)
	}
}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
}
