package main

import "fmt"

func Pointer() {
	b := *getPointer()
	fmt.Println("Value b is = ", b)
	fmt.Println("Value b address is = ", &b)

	// address
	c := getPointer()
	fmt.Println("Value c address is = ", c)

	var a int = 500
	var p *int

	p = &a

	fmt.Printf("p의 값 %p\n", p)
	fmt.Printf("p가 가르키는 메모리 값: %d\n", *p)

	*p = 100
	fmt.Printf("a의 값: %d\n", a)

}

func getPointer() (myPointer *int) {
	a := 234
	return &a
}
