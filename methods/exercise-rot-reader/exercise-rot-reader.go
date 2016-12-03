package main

import (
  "io"
  "os"
  "strings"
)

type rot13Reader struct {
  r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
  n, e := rot.r.Read(b)

  for i := range b {
    switch {
    case ((b[i] >= 'a' && b[i] < 'n') || (b[i] >= 'A' && b[i] < 'N')):
      b[i] += 13
    case ((b[i] >= 'n' && b[i] <= 'z') || (b[i] >= 'N' && b[i] <= 'Z')):
      b[i] -= 13
    }
  }
  return n, e
}

func main() {
  s := strings.NewReader("Lbh penpxrq gur pbqr!")
  r := rot13Reader{s}
  io.Copy(os.Stdout, &r)
}
