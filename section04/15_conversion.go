package main

import "fmt"

func main() {
	a := 42.3434
	fmt.Printf("%v of type %T\n", a, a)

	var b float32
	b = float32(a)
	fmt.Printf("%v of type %T\n", b, b)
}