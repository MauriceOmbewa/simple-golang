package main

import "fmt"

func lift(n int) {
	if n < 0 {
		return
	}
	fmt.Println(n)
	lift(n - 1)

}

func main() {
	lift(10)
	fmt.Println("Lift off")
}
