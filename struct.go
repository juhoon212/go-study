package main

import (
	"fmt"
	"unsafe"
)

func structStudy() {
	var house House
	house.Address = "서울시 강남구 논현동"
	house.Size = 28
	house.Price = 10000000
	house.Category = "아파트"

	fmt.Printf("%v\n", house)

	user := User{"홍길동", "1", 32}
	fmt.Printf("%v ", user)

	vip := VIPUser{
		User{"홍길동", "abc1234", 32},
		1,
		30000,
		4,
	}

	fmt.Printf("VIP 유저: %s ID: %s, 나이: %d VIP레벨: %d VIP가격: %d만원\n", vip.UserInfo.Name, vip.UserInfo.ID, vip.UserInfo.Age, vip.VIPLevel, vip.Price)

	example := Example{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(example)) // 40

}

type House struct {
	Address  string
	Size     int
	Price    int
	Category string
}

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
	House    int
}

type Example struct {
	A int8 // 1
	B int  // 8byte in 64bit os
	C int8 // 1
	D int  // 8
	E int8 // 1

}
