package main

import "fmt"

var id []int

func makeSet(n int) {
	for i := 0; i < n; i++ {
		id = append(id, i)
	}
}

func root(i int) int {
	for i != id[i] {
		fmt.Println("i:", i, id[i])
		i = id[i]
	}
	return i
}

func connected(p int, q int) bool {
	return root(p) == root(q)
}

func union(p int, q int) {
	i := root(p)
	j := root(q)
	id[i] = j
}

func main() {
	makeSet(10)
	union(4, 3)
	union(3, 8)
	union(6, 5)
	union(9, 4)
	union(2, 1)
	union(5, 0)
	union(7, 2)
	union(6, 1)
	union(7, 3)
	fmt.Println("---------------------")
	root(6)
	// check := connected(6, 8)
	// fmt.Println("check:", check)
}
