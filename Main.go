package main

import (
	"fmt"
)

func main() {
	a := 7 == 7
	b := 4 == 7

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
}
