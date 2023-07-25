package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("initialized here")
}

func main() {
	x := rand.Intn(251)
	fmt.Printf("x has value of %v\n", x)

	/*
	if x <= 100 {
		fmt.Println("between 0 and 100")
	} else if x > 100 && x <= 200 {
		fmt.Println("between 101 and 200")
	} else if x > 200 && x <= 250 {
		fmt.Println("between 201 and 250")
	}
	*/

	switch {
		case x <= 100:
			fmt.Println("between 0 and 100")
		case x > 100 && x <= 200:
			fmt.Println("between 101 and 200")
		case x > 200 && x <= 250:
			fmt.Println("between 201 and 250")
		default:
			fmt.Println("something else")
	}
}
