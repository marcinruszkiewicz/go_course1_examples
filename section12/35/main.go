package main

import (
  "fmt"
)

func main() {
  xi := []int{42, 43, 44, 45, 46, 47}

  for i, val := range xi {
    fmt.Printf("index %v: \t %v\n", i, val)
  }
}
