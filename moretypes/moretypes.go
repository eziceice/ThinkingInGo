package main

import (
	"fmt"
	"math"
)

func arrayExample() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	fmt.Println(primes[1:4]) //includes first, exclude last
}

func sliceWithMake() {
	a := make([]int, 5) // use make could create dynamic slices
	fmt.Println(a)
	b := make([]int, 0, 5)
	fmt.Println(b)
	c := b[:2]
	fmt.Println(c)
	d := c[2:5]
	fmt.Println(d)
}

func rangeExample() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Println(i, v)
	}
}

func mapExample() {
	type Vertex struct {
		Lat, Long float64
	}

	m := make(map[string]Vertex)

	m["test"] = Vertex{1, 2}

	fmt.Println(m)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func functionExample() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func pointer() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 10
	fmt.Println(v)

}
