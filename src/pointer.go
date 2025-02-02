package main

import "fmt"

func Pointer() {
	b := *getPointer()
	fmt.Println("Value is = ", b)
}

func getPointer() (myPointer *int) {
	a := 234
	return &a
}
