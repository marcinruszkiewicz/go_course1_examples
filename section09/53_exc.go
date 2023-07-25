package main

import (
	"fmt"
)

var x = 40
const y = 41

func main() {
	fmt.Printf("value is %v and type is %T\n", x, x)
	fmt.Printf("value is %v and type is %T\n", y, y)
	z := 42
	fmt.Printf("value is %v and type is %T\n", z, z)
}