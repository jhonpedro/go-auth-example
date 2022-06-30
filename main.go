package main

import (
	"fmt"
)

func main() {
	x := 123

	if x := true; x {
		x := 124

		fmt.Println(x)
	}

	fmt.Println(x)
}
