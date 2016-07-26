package main

import (
  "fmt"
)

func printer(msg string) {
  defer fmt.Printf("closing now...\n")
  defer fmt.Print("Then this \n")
  defer fmt.Printf("do this first at defer\n")
  noBytes, err := fmt.Printf("%s\n", msg)
  if err == nil {
      fmt.Printf("%d bytes printed...", noBytes)
  }

}
func main() {
  printer("Hello sam")

}
