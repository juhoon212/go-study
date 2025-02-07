package main

import "fmt"

func dataType() {
	// go 언어에는 복소수를 표현할 수 있는 complex64, complex128이 존재한다.

	var i float32 = 10.0
	var i2 int8 = 2
	var i3 int32 = 3

	fmt.Printf("i = %6.2f\n", i)
	fmt.Printf("i2 = %02d\n", i2)
	fmt.Printf("i3 = %02d\n", i3)

	var (
		str1, str2 = "string1", "string2"
	)

	fmt.Println("str1 = ", str1, "str2 = ", str2)
}
