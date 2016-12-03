package main

import "fmt"

func do(i interface{}) {
  switch v := i.(type) {
  case int:
    fmt.Printf("Twice %v is %v\n", v, 2*v)
  case string:
    fmt.Printf("%q is %v bytes long\n", v, len(v))
  default:
    fmt.Printf("I don't know about type %T!\n", v)
  }
}

func main() {
  do(21)
  do("chunky bacon")
  do(true)

  // use of .(type) outside type switch doesn't work
  // var i interface{}
  // fmt.Println(i.(type))
}
