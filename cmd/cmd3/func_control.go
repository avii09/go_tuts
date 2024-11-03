package main

import "fmt"

func main() {
	firstfun() // calling the function

	newint := 50
	fmt.Printf(checkfun(newint)) // calling the function
}

func firstfun() {
	fmt.Println("hello")
}

func checkfun(newint int) (string, int) { //taking the parameter of type int, it will return a string and int these are the return types.
	if newint%2 == 0 {
		return "%v is even", newint
	} else {
		return "%v is even", newint
	}
}

// we can make functions and pass parameters like this ==> func <func name> (<paramenter> <type) <return type> {}
// the no of returns tht are there, we need to specify the no of return types. eg. return int1, int2 , the function will have (int, int) as return type
// if - else are the control statments.
// the if - else statemnts need to be in the same line as the last bracket.
