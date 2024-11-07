package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rect struct {
	width  float64
	height float64
}

func (c Circle) Area() float64 {
	return 2 * math.Pi * c.radius * c.radius
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func printArea(s Shape) {
	fmt.Print("The area is:", s.Area(), "\n")
}

func main() {
	var rec Rect = Rect{3, 5}
	var cir Circle = Circle{4}
	printArea(rec)
	printArea(cir)
}

// structs are user defined data types.
// struct ==> type <struct type> struct { field1 <datatype>
//                                        field2 <datatype>...}

// declaration n initialization ==> var <struct-variable> <struct type> = <strcut-type>{field1:<value>, ...}
// initialization ==> struct-variable.field1 = <value>

// methods are similar to functions except it is directly tied to the struct.
// func (<struct-var> <struct-type>) <func name> (return type){}
// this is similar to java where the struct acts like a class and the struct-var is like the instance or obj of the class.
// we can call the method using the struct var, just like we call the method using the obj in java

// interfaces is like a blueprint tht consists of abstract methods.
// we can access these methods from an interface variable.
// type <interface-type> interface { method1 <datatype>
//                                   method2 <datatype>....}
