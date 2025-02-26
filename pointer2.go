package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func Pointer2() {

	var a int = 10
	var b int = 20

	var p1 *int = &a
	var p2 *int = &a
	var p3 *int = &b

	fmt.Printf("p1 == p2: %v\n", p1 == p2)
	fmt.Printf("p2 == p3: %v\n", p2 == p3)

	var data Data
	ChangeData(data)
	fmt.Printf("value = %d\n", data.value)

	// 구조체 포인터 초기화
	var p *Data = &Data{}
	p.value = 3
	p.data[100] = 4
	fmt.Printf("p.value = %d, p.data[100] = %d\n", p.value, p.data[100])

	p4 := &Data{
		value: 3,
		data:  [200]int{1, 2, 3},
	}

	fmt.Printf("p4 = %d, %d\n", p4.value, p4.data[0])
}

func ChangeData(arg Data) {
	arg.value = 999
	arg.data[100] = 999
}
