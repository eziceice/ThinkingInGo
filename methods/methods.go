package main

import (
	"fmt"
	"math"
	"time"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	/**
	With a value receiver, the Scale method operates on a copy of the original Vertex value.
	(This is the same behavior as for any other function argument.)
	The Scale method must have a pointer receiver to change the Vertex value
	*/
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println(v.X)
	fmt.Println(v.Y)
}

func typeAssertions() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string { // this is actually just override toString() method in Java
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age) // use Sprintf()
}

type MyError struct {
	When time.Time
	What string
}

/**
type error interface {
    Error() string
}
*/
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	p := &Vertex{4, 3}
	fmt.Println(p.Abs()) // p.Abs() -> (*p).Abs()

	if err := run(); err != nil {
		fmt.Println(err) // when call println actually it's looking for Error() string method
	}
}
