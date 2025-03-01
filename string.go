package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func StringStudy() {
	str1 := "Hello 월드"
	str2 := str1

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)

	str3 := "Hello world!"
	arr := []rune(str3) // rune type 8byte

	for i := 0; i < len(arr); i++ {
		fmt.Printf("type: %T, value: %d, string: %c\n", arr[i], arr[i], arr[i])
	}

	str4 := "안녕"
	runes := []rune(str4)

	fmt.Printf("str4 of length: %d\n", len(str4))
	fmt.Printf("runes of length: %d\n", len(runes))

	for v := range runes {
		fmt.Printf("type: %T value: %d rune: %c\n", v, v, v)
	}

	fmt.Println(len(arr))

	str5 := "Hello world"
	slice := []byte(str5)

	stringHeader3 := (*reflect.StringHeader)(unsafe.Pointer(&str5))
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))

	fmt.Printf("str5: \t%x\n", stringHeader3.Data)
	fmt.Printf("slice: \t%x\n", sliceHeader.Data)

}
