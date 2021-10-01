package main

import (
	"fmt"
	"math/cmplx"
)

var (
	toBe          = false
	maxInt uint64 = 1<<64 - 1
	z             = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const WORLD = "world"
	fmt.Printf("%v", toBe)
	fmt.Printf("%v", maxInt)
	fmt.Printf("%v", z)
	fmt.Println("hi", true, "12354")
}
