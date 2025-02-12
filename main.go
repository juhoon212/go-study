package main

import (
	"fmt"
	"go-study/src/forLoop"
)

func main() {
	message := greetMe("world")
	fmt.Println(message)

	varaibles()
	num()
	Pointer()
	Arrays()
	cond("monday")
	forLoop.ForLoop()
	variables2()
	dataType()
	format()
}

func greetMe(name string) string {
	return "Hello " + name + "!"
}

// variables

func varaibles() {
	var msg string
	msg = "Hello"
	fmt.Println(msg)
}

// constant

const Phi = 1.618 // boolean, string, numeric, char

// numbers

func num() {
	num1 := 3
	num2 := 3.
	num3 := 3 + 4i
	num4 := byte('a')

	var u uint = 7
	var p float32 = 22.7

	fmt.Println("num1 = ", num1)
	fmt.Println("num2 = ", num2)
	fmt.Println("num3 = ", num3)
	fmt.Println("num4 = ", num4)

	fmt.Println("u = ", u)
	fmt.Println("p = ", p)

}
