package custompkg

import (
	"fmt"
)

func PrintCustom() {
	fmt.Println("This is custom package!")
}

type Student struct {
	Name  string
	Are   int
	score int
}

const PI = 3.14
