package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	//for true {
	//	fmt.Print("\033[G\033[K")
	//
	//	fmt.Printf("Retrieved %d\n", i)
	//
	//	fmt.Print("\033[A")
	//	i += 1
	//
	//}

	fmt.Print("\033[s") // save the cursor position

	for true {
		fmt.Print("\033[u\033[K\n") // restore the cursor position and clear the line
		fmt.Printf("Retrieved %d", i)

		i += 1

		time.Sleep(1 * time.Second)
	}
}
