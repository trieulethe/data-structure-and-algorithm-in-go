package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

func main() {
	type Book struct {
		Pages int
		title string
	}
	var book = Book{} // book is addressable
	fmt.Println("book: ", book)
	fmt.Println("book: ", &book)
	s1 := fmt.Sprintf("%p", &book)
	fmt.Println(reflect.TypeOf(s1))
	i, err := strconv.ParseInt(s1, 16, 64)
	if err != nil {
		fmt.Println("i convert: ", i)
	} else {
		fmt.Println("err: ", err)
	}
	fmt.Println("sprintf", s1)
	fmt.Println("book title: ", book.title)
	fmt.Println("& book title: ", &book.title)
	b := &book
	p := &book.Pages
	s := &book.title
	pointer := unsafe.Pointer(&book)
	fmt.Println("&book: ", pointer)
	fmt.Println("&b: ", &b)
	fmt.Println("&p: ", &p)
	fmt.Println("s: ", s)
	fmt.Println("&s: ", &s)
	fmt.Println("&s: ", *&s)
	*p = 123
	fmt.Println(book) // {123}

	// The following two lines fail to compile,
	// for Book{}.Pages is not addressable.
	/*
		Book{}.Pages = 123
		p = &(Book{}.Pages) // <=> p = &(Book{}.Pages)
	*/
}
