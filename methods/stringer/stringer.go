package main

import "fmt"

type Person struct {
  Name string
  Age  int
}

func (p Person) String() string {
  return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
  monty := Person{"Ed Norton", 47}
  yoda := Person{"Yoda", 900}
  fmt.Println(monty, yoda)
}
