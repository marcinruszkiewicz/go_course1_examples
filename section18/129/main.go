package main

import (
  "fmt"
)

type engine struct {
  electric bool
}

type vehicle struct {
  engine
  make string
  model string
  doors int
  color string
}

func main(){
  v1 := vehicle{
    make: "Ford",
    model: "Fiesta",
    doors: 3,
    color: "White",
    engine: engine{
      electric: true,
    },
  }

  v2 := vehicle{
    make: "Toyota",
    model: "Corolla",
    doors: 5,
    color: "Red",
    engine: engine{
      electric: false,
    },
  }

  fmt.Println(v1)
  fmt.Println(v2)
  fmt.Println(v1.electric)
  fmt.Println(v2.electric)
}
