package main

import (
	"fmt"
)

func main() {
	i := 10

	for {
		fmt.Println(i)
		i--
		if i == 0 {
			break
		}
	}
}
