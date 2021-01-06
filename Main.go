package main

import (
	"fmt"
)

func main() {
	// Integers: Example below shows type conversion
	var a int = 10
	var b int8 = 3
	fmt.Println(a + int(b))

	// Operations avaiable to Integers

	// Arthmetic
	e := 7
	f := 4
	fmt.Println(e + f)
	fmt.Println(e - f)
	fmt.Println(e * f)
	fmt.Println(e / f)

	fmt.Println(e % f) // % is the modulo operator and shows the remainder of integer division

	// Bit Operators
	fmt.Println(e & f)
	fmt.Println(e | f)
	fmt.Println(e ^ f)
	fmt.Println(e &^ f)

	// Bit Shifting Operators
	fmt.Println(e << 4)
	fmt.Println(e >> 4)

	// Floating point: In the below example n = 13.7e72 is too large for a float32
	var n float64 = 3.14
	n = 13.7e72
	n = 2.1e14
	fmt.Printf("%v, %T\n", n, n)

	// Arithmetic Operations available to floating point numbers
	c := 9.7
	d := 4.8
	fmt.Println(c + d)
	fmt.Println(c - d)
	fmt.Println(c * d)
	fmt.Println(c / d)

}
