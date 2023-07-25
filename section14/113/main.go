package main

import (
  "fmt"
)

func main() {
  jb := []string{"James", "Bond", "Shaken"}
  mm := []string{"Miss", "Moneypenny", "008"}

  xxs := [][]string{jb, mm}

  for _, sl := range xxs {
    for _, val := range sl {
      fmt.Println(val)
    }
  }
}
