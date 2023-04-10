package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
  // Create a slice of slices of uint8
  pic := make([][]uint8, dy)
  // Loop through the slice of slices
  for y := range pic {
    // Create a slice of uint8
    pic[y] = make([]uint8, dx)
    // Loop through the slice of uint8
    for x := range pic[y] {
      // Set the value of the slice of uint8
      pic[y][x] = uint8((x + y) / 2)
    }
  }
  return pic

}

func main() {
  pic.Show(Pic)
}
