package main

import "fmt"

const n = 5

var rank [n]int
var parent [n]int

func makeSet() {
	for i := 0; i < n; i++ {
		parent[i] = i
	}
}

func find(x int) int {
	if parent[x] != x {
		parent[x] = find(parent[x])
	}
	return parent[x]
}

func union(x int, y int) {
	xRoot := find(x)
	yRoot := find(y)

	if xRoot == yRoot {
		return
	}

	if rank[xRoot] < rank[yRoot] {
		parent[xRoot] = yRoot
	} else if rank[yRoot] < rank[xRoot] {
		parent[yRoot] = xRoot
	} else {
		parent[yRoot] = xRoot
		rank[xRoot] = rank[xRoot] + 1
	}
}

func main() {
	makeSet()
	union(0, 2)
	union(4, 2)
	union(1, 3)

	if find(4) == find(0) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	if find(1) == find(0) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

}
