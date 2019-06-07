package main

import "fmt"

type item struct {
	id      int
	weight  int
	benefit int
	value   int
	amount  int
}

func fractionKnapsack(S []item, W int) {
	for i := 0; i < len(S); i++ {
		// x := 0
		S[i].value = S[i].benefit / S[i].weight
	}
	w := 0
	for w < W && len(S) != 0 {

	}
}

func main() {
	s := []item{
		{1, 4, 12, 0, 0},
		{2, 8, 32, 0, 0},
		{3, 2, 40, 0, 0},
		{4, 6, 30, 0, 0},
		{5, 1, 50, 0, 0},
	}

	fmt.Println("items:", s)
}
