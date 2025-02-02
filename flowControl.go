package main

import "fmt"

// Condition

func cond(day string) {
	if day == "Sunday" || day == "saturday" {
		rest()
	} else if day == "monday" && isTired() {
		work()
	} else {
		work()
	}
}

func isTired() bool {
	return true
}

func rest() {
	fmt.Println("u should rest!!")
}

func work() {
	fmt.Println("u should work!!")
}
