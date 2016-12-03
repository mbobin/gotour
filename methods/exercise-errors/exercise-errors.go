package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) (float64, error) {
  if x < 0 {
    return 0, ErrNegativeSqrt(x)
  }
  z := 2.0
  result := 0.0
  epison := 1e-15

  for {
    z = z - (z*z-x)/(2*z)
    if math.Abs(result-z) < epison {
      break
    }
    result = z
  }
  return result, nil
}

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(-2))
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
  return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}
