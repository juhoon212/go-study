package main

import "fmt"

func array() {

	nums := [...]int{10, 20, 30}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	var t [5]float64 = [5]float64{24.0, 23.11, 27.8, 24.2, 26.12}

	for i, v := range t {
		fmt.Printf("i = %v, v = %v ", i, v)
	}

	for _, v := range t {
		fmt.Println(v)
	}

	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}
