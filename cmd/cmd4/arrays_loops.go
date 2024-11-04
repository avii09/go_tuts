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

}

// arrays are declared as ==> var <array_name> [size]<data-type> or <arr_name> := [size]<data-type>
// arrays are of fixed length, instead of providing the size it can also infer by using [..] dots tht signify the size
// arrays contain elements of the same data type.
// arrays store elements in contigous memory locations.
// arrays can be indexed.
