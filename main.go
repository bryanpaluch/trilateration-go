package main

import (
  "fmt"
)

func main() {

  a := Cordinate{10, 0, 2.5}
  b := Cordinate{7.5, 2.5, 2.5}
  c := Cordinate{5, 0, 2.5}
  fmt.Println("Cordinates of a:", a)
  fmt.Println("Cordinates of b:", b)
  fmt.Println("Cordinates of c:", c)
  x, y := a.Trilaterate(b, c)
  fmt.Println("Cordinates of intersections center x:", x, "y:", y)
}
