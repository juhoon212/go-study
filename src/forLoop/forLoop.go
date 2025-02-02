package forLoop

import "fmt"

func ForLoop() {
	entry := []string{"Jack", "John", "Jones"}
	for i, val := range entry {
		fmt.Printf("At position %d, the character %s", i, val)
	}
}
