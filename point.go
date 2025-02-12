package main

import "fmt"

type point struct {
	x, y int
}

func format() {
	p := point{1, 2}

	// 객체 출력 {1 2}
	fmt.Printf("%%v %v\n", p)
	// 필드명 추가{x:1 y:2}
	fmt.Printf("%%+v %+v\n", p)
	// 코드 스니펫 출력 main.point{x:1, y:2}
	fmt.Printf("%%#v %#v\n", p)
	// 타입 출력 main.point
	fmt.Printf("%%T %T\n", p)
	// boolean true
	fmt.Printf("%%t %t\n", true)
	// 10진수 출력 123
	fmt.Printf("%%d %d\n", 123)
	// binary 출력 1110
	fmt.Printf("%%b %b\n", 14)
	// 정수에 해당하는 문자 출력 ! ASCII
	fmt.Printf("%%c %c\n", 33)
	// 16진수 인코딩 출력 1C8
	fmt.Printf("%%x %x\n", 456)
	// 실수 출력
	fmt.Printf("%%.2f %.2f\n", 75.999999)
	// 실수 출력2
	fmt.Printf("%%e %e\n", 75.999999)
	// 문자열 출력
	fmt.Printf("%%s %s\n", "string")
	// 문자열에 쌍따옴표를 추가하여 출력
	fmt.Printf("%%q %q\n", "string")
	// 문자열 16진수 인코딩 출력
	fmt.Printf("%%x %x\n", "hex this")
	// 포인터 표현 출력 (주소값 출력))
	fmt.Printf("%%p %p\n", &p)
}
