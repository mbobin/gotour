package main

import (
  "fmt"
  "io"
  "strings"
)

func main() {
  reader := strings.NewReader("chunky bacon")

  b := make([]byte, 8)

  for {
    n, err := reader.Read(b)
    fmt.Printf("n = %v err= %v b = %v\n", n, err, b)
    fmt.Printf("b[:n] = %q\n", b[:n])

    if err == io.EOF {
      break
    }
  }
}
