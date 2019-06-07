package main

import (
	"fmt"
)

func main() {
	// a := 2
	// p := &a
	// *p = 3
	// fmt.Println("a: ", a)

	// // struct
	// type Book struct {
	// 	Pages int
	// 	title string
	// }
	// var book = Book{} // book is addressable
	// fmt.Println("book: ", book)
	// fmt.Println("book: ", &book)
	// p2 := &book
	// *p2 = Book{1, "hwk"}
	// // *p2 = Book{}
	// fmt.Println("p2:", p2)
	// fmt.Println("book", book)
	// pointer := unsafe.Pointer(&book)
	// fmt.Println("pointer: ", pointer)
	// // invalid indirect of pointer
	// // *pointer = Book{4, "change value"}

	fmt.Println("result:", 1 &= 1)

}
