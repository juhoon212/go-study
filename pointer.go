package main

import "fmt"

func Pointer() {
	b := *getPointer()
	fmt.Println("Value b is = ", b)
	fmt.Println("Value b address is = ", &b)

	c := getPointer()
	fmt.Println("Value c address is = ", c)
}

func getPointer() (myPointer *int) {
	a := 234
	return &a
}
