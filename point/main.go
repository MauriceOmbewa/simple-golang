package main

import "github.com/01-edu/z01"

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printString(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func Int(n int) {
	number := '0'
	if n/10 > 0 {
		Int(n / 10)
	}
	for i := 0; i < n%10; i++ {
		number++
	}
	z01.PrintRune(number)
}

func main() {
	points := &point{}

	setPoint(points)

	printString("x = ")
	Int(points.x)
	printString(", y = ")
	Int(points.y)
	z01.PrintRune('\n')
}
