package main

import (
  "fmt"
)

func main() {
  m := map[string]int{
    "James": 42,
    "Notjames": 43,
  } 

  // for k, v := range m {
  //   fmt.Printf("%v: %v\n", k, v)
  // }

  m1 := m["James"]
  fmt.Println(m1)

  m2 := m["Q"]
  fmt.Println(m2)

  if _, ok := m["James"]; !ok {
    fmt.Println("not in")
  } else {
    fmt.Println("in")
  }
}
