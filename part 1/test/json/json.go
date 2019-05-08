package main

import (
	"encoding/json"
	"fmt"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// bolB, _ := json.Marshal(true)
	// fmt.Println("bolB:", string(bolB))

	// intB, _ := json.Marshal(1)
	// fmt.Println(string(intB))

	// fltB, _ := json.Marshal(2.34)
	// fmt.Println(fltB)
	// fmt.Println(string(fltB))

	// strB, _ := json.Marshal("gopher")
	// fmt.Println(strB)
	// fmt.Println(string(strB))

	// slcD := []string{"apple", "peach", "pear"}
	// slcB, _ := json.Marshal(slcD)
	// fmt.Println(slcB)
	// fmt.Println(string(slcB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
}
