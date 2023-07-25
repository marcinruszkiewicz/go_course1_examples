package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		x := rand.Intn(11)
		y := rand.Intn(11)

		fmt.Printf("run %v\tx: %v\ty: %v\t", i, x, y)

		switch {
			case x < 4 && y < 4:
				fmt.Println("both less than 4")
			case x > 6 && y > 6:
				fmt.Println("both greater than 6")
			case x >= 4 && x <= 6:
				fmt.Println("x between 4 and 6")
			case y != 5:
				fmt.Println("y is not 5")
			default:
				fmt.Println("none met")
		}
	}
}
