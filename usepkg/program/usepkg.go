package main

import (
	"fmt"
	"go-study/usepkg/custompkg"

	"go-study/usepkg/exinit"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.PrintCustom()
	custompkg.Print2()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)

	s := custompkg.Student{}
	s.Name = "화랑"
	s.Are = 32
	fmt.Println(s.Name, s.Are)
	fmt.Println(custompkg.PI)

	exinit.PrintD()
}
