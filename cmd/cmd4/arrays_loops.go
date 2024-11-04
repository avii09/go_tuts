package main

import "fmt"

func main() {
	// arrays
	var intarr [3]int //declaring array
	intarr[0] = 1     //indexing
	intarr[1] = 2
	intarr[2] = 3
	fmt.Println(intarr[1:3])

	arr := [...]int{4, 5, 6} //initializing array
	fmt.Println(arr)

	// slices
	intslice := []int{8, 9, 10} //variable length
	intslice = append(intslice, 7)
	fmt.Println(intslice)

	intslice2 := []int{2, 6}
	intslice = append(intslice, intslice2...)
	fmt.Println(intslice)

}

// arrays are declared as ==> var <array_name> [size]<data-type> or <arr_name> := [size]<data-type>
// arrays are of fixed length, instead of providing the size it can also infer by using [..] dots tht signify the size
// arrays contain elements of the same data type.
// arrays store elements in contigous memory locations.
// arrays can be indexed.

// slices are arrays with some additional functionalities.
// in slice, i can omit the size such that unlike arrays it doesnt have a fixed length and i can add more elements to the sequence
// using append() i can append an element to the end of an existing array.

// capacity ( cap() ) and length ( len() ), cap ==> the total capacity of a slice, len ==> the no of elements in the slice.
// eg. {8, 9, 10, 7, *, *}, here len=4 and cap=6

// another way of making slice is by using the make function make(<datatype>[], len, cap)
