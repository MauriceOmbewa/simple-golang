package piscine

import "fmt"

func QuadC(x, y int) {
	if x > 0 && y > 0 {
		// First line
		// The program starts only when the value of x is greater than 0
		fmt.Print("A")
		for length := 1; length <= x-2; length++ {
			fmt.Print("B")
			if length == x-2 {
				fmt.Println("A")
			}
		}
		if !(x > 1) {
			fmt.Println()
		}

		// Middle part
		// Only works when the value of why is equal to 2 or greater
		if y >= 2 {
			for width := 0; width <= y-3; width++ {
				fmt.Print("B")
				if !(x > 1) {
					fmt.Println()
					for length := 1; length <= x-2; length++ {
						fmt.Print(" ")
						if length == x-2 {
							fmt.Print("B")
						}
					}
				} else {
					for length := 1; length <= x-2; length++ {
						fmt.Print(" ")
						if length == x-2 {
							fmt.Println("B")
						}
					}
				}
			}
		}

		// Final part
		// Only works when y is greater than 2
		if y > 2 {
			fmt.Print("C")
			if !(x > 1) {
				fmt.Println()
			}
			for length := 1; length <= x-2; length++ {
				fmt.Print("B")
				if length == x-2 {
					fmt.Println("C")
				}
			}
		}
	} else {
		fmt.Println()
	}
}
