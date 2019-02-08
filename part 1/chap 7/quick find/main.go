package main

import "fmt"

var id []int

func quickFind(n int) {
	for i := 0; i < n; i++ {
		id = append(id, i)
	}
}

func connected(p int, q int) bool {
	return id[p] == id[q]
}

func union(p int, q int) {
	pid := id[p]
	qid := id[q]
	for i := 0; i < len(id); i++ {
		if id[i] == pid {
			id[i] = qid
		}
	}
}

func main() {
	quickFind(5)
	fmt.Println("id:", id)
}
