package main

import "fmt"

func variables2() {
	var a, b, c int = 10, 11, 12
	s4, s5, s6 := "string1", "string2", "string3"

	fmt.Println(a, b, c)
	fmt.Println(s4, s5, s6)

	const (
		A = iota
		B
	)

	fmt.Println(A, B)
}
