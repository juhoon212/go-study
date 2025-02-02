package main

import "fmt"

func Arrays() {
	numbers := [...]int{0, 0, 0, 0, 0}

	for i := 1; i <= len(numbers); i++ {
		fmt.Println("numbers = i", i)
	}
}
