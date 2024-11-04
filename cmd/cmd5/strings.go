package main

import "fmt"

func main() {
	newstr := "Hello"
	index := newstr[0]
	fmt.Println(index) // this will print the ASCII value of 'H'

	for i, v := range newstr {
		fmt.Printf("Index : %v, Value : %v \n", i, v) // this will print the index and the ASCII value of each letter.
	}

	newstr2 := []rune("Avantika")
	for i, v := range newstr2 {
		fmt.Printf("Index : %v, Value : %v \n", i, v) // this will print the index and the ASCII value of each letter.
	}

	newstr3 := "résumé"
	fmt.Println(len(newstr3)) // o/p is 8 as it counts the no. of bytes of each char

	newstr4 := []rune("résumé")
	fmt.Println(len(newstr4)) // o/p is 6 it counts the no. of char.

	newstr5 := []string{"a", "v", "a", "n", "t", "i", "k", "a"}
	catstr := ""
	for i := range newstr5 {
		catstr += newstr5[i]
	}
	fmt.Printf(" \n %v", catstr)

}

// go used utf-8 encoding to represent strings in your computer.
// strings are represented as ASCII values which are basically binary nos representing bytes.
// we use rune data type.
// mainly because when we have a special case letter like a symbol which is out of the spectrum of vanilla ASCII values(0-127), it usually takes 2 or more bytes. so to index it properly such tht we have contigous index order.
// len() gives us the no of bytes not the length of the string, however when we use the rune data type, we get the length of the string
