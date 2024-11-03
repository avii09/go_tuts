package main

import "fmt"

func main() {
	// int type
	var Newvar int = 32767
	fmt.Println(Newvar)

	// float type
	var Newfloat float32 = 1238.9
	fmt.Println(Newfloat)

	// operation "+"
	var res float32 = float32(Newvar) + Newfloat
	fmt.Println(res)
	// res = Newvar + Newfloat (will throw an error because of different types)

	var num1 int = 3
	var num2 int = 2
	fmt.Println(num1 / num2) // int division rounds the result down, use % to get the remainder.

	//string
	var Mystr string = `heelo 
	world`
	fmt.Println(Mystr)

	// concatnate strings
	var Str string = "Hello" + " " + "World"
	fmt.Println(Str)

	// length of string
	fmt.Println(len("hello"))

	// inference, shorthand, not initialized
	Newstr := "Yoooo"
	fmt.Println(Newstr)
	Newint := 12
	fmt.Println(Newint)
	var try int
	fmt.Println(try)

	// initializing multiple variables
	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	// constants
	const mycon string = "yo whats up" //cannot change value
	fmt.Println(mycon)

}

// const <constant name> ==> declaring a constant
// var <variable name> <data type> ==> declaring a variable
// <variable name> := <value to be stored>

// data types ==> bool, float32, float64, int, int16, int32, int64, rune, string, uint, uint16, uint32, uint64.
// we need to use every variable we declare
// go doesnt have only "float" type, we need to specify either float32 or float64. just specifying float will show an error.
// you cant perform operations on mixed types, ie, you cant add a float data type and an int data type.
// we can cast one of the type such tht both the operands are of the same type and tht we can perform an operation.

// var <variable name> string ==> to store strings
// you can initialize a string value using " " or ` ` (back ticks).
// " " ==> cannot format directly,  ` ` ==> can format strings directly

// len() give the length of string, but it doesnt actually count the no of characters but instead gives the no of bytes.
// so if we have a special char in the string we import the utf8 encoding to ge the no of characters and also work with rune data type.

// it is not needed to specify the data type, we can also initialize the variable and it will infer the data type automatically.
// but it is good practice if we do specify the data type, as sometimes if it is returning a value and we dont know what the type of the value is.

// you cannot change the value of constants, and constants cannot be only declared you need to initialize it too.
