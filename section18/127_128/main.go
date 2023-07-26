package main

import (
  "fmt"
)

type person struct {
  first string
  last string
  ice []string
}

func main() {
  p1 := person{
    first: "James",
    last: "Bond",
    ice: []string{"Vanilla", "Chocolate"},
  }

  p2 := person{
    first: "Dr",
    last: "No",
    ice: []string{"Vanilla"},
  }

  m := make(map[string]person)
  m[p1.last] = p1
  m[p2.last] = p2

  for _, person := range m {
    for _, v := range person.ice {
      fmt.Println(person.first, "fav ice cream is", v)
    }
  }
}
