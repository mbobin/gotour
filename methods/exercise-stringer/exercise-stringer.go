package main

import "fmt"

type IPAddr [4]byte

// V1
// func (ip IPAddr) String() string {
//   return fmt.Sprintf("%v.%v.%v.%v",
//     ip[0], ip[1], ip[2], ip[3])
// }

// V2
func (ip IPAddr) String() string {
  str := fmt.Sprintf("%v", ip[0])
  for _, i := range ip[1:] {
    str += fmt.Sprintf(".%v", i)
  }
  return str
}

func main() {
  hosts := map[string]IPAddr{
    "loopback":  {127, 0, 0, 1},
    "googleDNS": {8, 8, 8, 8},
  }
  for name, ip := range hosts {
    fmt.Printf("%v: %v\n", name, ip)
  }
}
