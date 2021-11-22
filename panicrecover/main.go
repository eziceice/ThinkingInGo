package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	fmt.Println(div(1, 0))
}

func div(x, y int) int {
	return x / y
}
