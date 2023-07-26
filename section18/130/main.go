package main

import (
  "fmt"
)

func main() {
  p1 := struct{
    first string
    friends map[string]int
    favDrinks []string
  }{
    first: "James",
    friends: map[string]int{
      "best": 5,
      "average": 1,
    },
    favDrinks: []string{"Martini", "Vodka"},
  }

  fmt.Println(p1)
}
