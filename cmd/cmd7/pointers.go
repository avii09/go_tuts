package main

import "fmt"

func main() {

	//pointers
	i := 10
	p := &i

	fmt.Println("value of p:", *p)

	*p = 20
	fmt.Println("value of i:", i)

	// allocating memory address to a pointer
	var q *int = new(int) // this means tht there is a pointer q of type int, new allocates a memory address for an int and then returns a pointer to tht address
	fmt.Println(*q)       // outputs 0 because there is no value stored in tht memory address
	*q = 8
	fmt.Println(*q) // will output 8 now cause now we are storing 8 at tht memory location.

	// functions and pointers
	arr1 := [5]float64{1, 2, 3, 4, 5}
	fmt.Println("The squares are:", squares(&arr1)) // here we are

}

func squares(arr2 *[5]float64) [5]float64 { // func taking in paras as a pointer to the arr1
	for i := range arr2 {
		arr2[i] = arr2[i] * arr2[i]
	}
	return *arr2
}

// var <pointer-var> *<datatype>
// pointers are special types tht store the memory address of another variable
// it is like a reference to a location, where a value is stored

// we use the & for storing the address of a variable and we use the * to get the value at tht address
// when we modify the value using the pointer (*p) the if suppose this pointer p is pointing to some variable i, then the value of i also changes. this is because p is pointing to the memory address of i.

// we can also allocate memory address of some type to a pointer. basically initializing a pointer with a memory address
// this is done using new(<datatype>)
