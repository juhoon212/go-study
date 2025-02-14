package main

import "fmt"

func bit() {
	fmt.Println("10^2 = ", 10^2)
	fmt.Println("10^&2 = ", 10&^2)

	var x int8 = 4
	var y int8 = 64

	fmt.Printf("x:%08b x<<2:%08b x<<2: %d\n", x, x<<2, x<<2)
	fmt.Printf("y:%08b y<<2:%08b y<<2: %d\n", y, y<<2, y<<2)

	var x1 int8 = 16
	var y1 int8 = -128
	var z1 int8 = -1
	var w1 uint8 = 128

	fmt.Printf("x1:%08b x1>>2: %08b x1>>2: %d\n", x1, x1>>2, x1>>2)
	fmt.Printf("y1:%08b y1>>2: %08b y1>>2: %d\n", uint8(y1), uint8(y1>>2), uint8(y1)>>2)

	fmt.Printf("z1:%08b z1>>2: %08b z1>>2: %d\n", uint8(z1), uint8(z1)>>2, uint8(z1)>>2)
	fmt.Printf("w1:%08b w1>>2: %08b w1>>2: %d\n", w1, w1>>2, w1>>2)

}
