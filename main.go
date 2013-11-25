package main

import (
  "encoding/json"
  "fmt"
  "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello world %s!", r.URL.Path[1:])
}

type cordinate_group struct {
  cordinates [3]Cordinate
}

func postTest(w http.ResponseWriter, r *http.Request) {

  decoder := json.NewDecoder(r.Body)
  var c cordinate_group
  err := decoder.Decode(&c)
  if err != nil {
    panic(err)
  }
  fmt.Println(c)
  fmt.Fprintf(w, "Hello world ", c)

}
func main() {

  a := Cordinate{100, 100, 10.0}
  b := Cordinate{160, 120, 36.06}
  c := Cordinate{70, 150, 60.83}
  fmt.Println("Cordinates of a:", a)
  fmt.Println("Cordinates of b:", b)
  fmt.Println("Cordinates of c:", c)
  // Get the trilateration using the derived formula that solves for x and y
  // From the equations of all 3 circles
  x, y := a.Trilaterate(b, c)
  fmt.Println("Cordinates of intersections center using formula x:", x, "y:", y)
  x, y = a.TrilaterateRast(b, c)
  fmt.Println("Cordinates of intersections center using Rasteration x:", x, "y:", y)

  http.HandleFunc("/", helloWorld)
  http.HandleFunc("/trilaterate", postTest)
  http.ListenAndServe(":8081", nil)
}
