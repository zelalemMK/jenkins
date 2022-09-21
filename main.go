package main

import "fmt"

const (
	pi    = 3.123231
	Truth = false
	big   = 1 << 62
	small = big >> 61
)

func main() {
	fmt.Println(big, small, Truth, pi)
}
