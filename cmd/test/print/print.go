package main

import (
	"fmt"
)

func main() {
	i := 0

	for true {
		fmt.Print("\033[G\033[K")

		fmt.Printf("Retrieved %d\n", i)
		fmt.Print("\033[A")
		i += 1

	}
}
