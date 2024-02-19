package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var numbers []int

	for i := 1; i < len(os.Args); i++ {
		num, err := strconv.Atoi(os.Args[i])
		if err == nil {
			numbers = append(numbers, num)
		}
	}

	largest := numbers[0]
	for _, num := range numbers {
		if num > largest {
			largest = num
		}
	}

	fmt.Println("The largest number is:", largest)
}