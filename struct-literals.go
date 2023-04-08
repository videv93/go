package main

import "fmt"

type Vertext struct {
  X, Y int
}

var (
  v1 = Vertext{1, 2}  // has type Vertext
  v2 = Vertext{X: 1}  // Y:0 is implicit
  v3 = Vertext{}      // X:0 and Y:0
  p  = &Vertext{1, 2} // has type *Vertext
)

func main() {
  fmt.Println(v1, p, v2, v3)
}
