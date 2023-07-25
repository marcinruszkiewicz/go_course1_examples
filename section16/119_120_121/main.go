package main

import (
  "fmt"
)

func main() {
  m := map[string][]string{
    "bond_james": []string{"shaken, not stirred", "martinis", "fast cars"},
    "moneypenny_jenny": []string{"intelligence", "literature", "computer science"},
    "no_dr": []string{"cats", "ice cream", "sunsets"},
  }

  m["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

  delete(m, "no_dr")

  for key, xs := range m {
    for i, v := range xs {
      fmt.Println(key, i, v)
    }
  }
}
