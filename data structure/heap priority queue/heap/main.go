package main

import "fmt"

var Heap []int
var size int
var last int

func insert(value int) {
	Heap = append(Heap, value)
	// fmt.Println("Heap:", Heap)
	size++
	if size > 1 {
		upHeap(size - 1)
	}
}

func parent(key int) int {
	if key&1 == 1 {
		return (key - 1) >> 1
	}
	return (key - 2) >> 1
}

func swap(a int, b int) {
	temp := Heap[a]
	Heap[a] = Heap[b]
	Heap[b] = temp
}

func upHeap(position int) {
	if position == 0 {
		return
	}
	parent := parent(position)
	// fmt.Println("position:", position)
	// fmt.Println("parent:", parent)
	if Heap[position] < Heap[parent] {
		swap(position, parent)
		upHeap(parent)
	}
	return
}

func downHeap(position int, check int) {
	child := position*2 + check
	if child >= size {
		return
	}
	if Heap[position] > Heap[child] {
		swap(position, child)
		downHeap(child, check)
	}
}

func handleHeapDown(position int) {
	leftChild := Heap[position*2+1]
	rightChild := Heap[position*2+2]
	if leftChild < rightChild {
		downHeap(position, 1)
	} else {
		downHeap(position, 2)
	}
}

func extractMin() {
	Heap[0] = Heap[size-1]
	Heap[size-1] = 0
	Heap = Heap[:(size - 1)]
	size--
	handleHeapDown(0)
}

func main() {
	// Heap = append(Heap, 1)
	// arr := []int{8, 6, 4, 5, 2}
	arr := []int{8, 6, 4, 5, 2, 2, 9, 7, 3}
	// arr := []int{2, 2, 3, 4, 5, 7, 9, 8, 6}
	// fmt.Println("size:", size)
	// p1 := parent(8)
	// fmt.Println("p1:", p1)
	// fmt.Println("arr[p1]:", arr[p1])
	for i := 0; i < len(arr); i++ {
		insert(arr[i])
	}
	fmt.Println("HEAP:", Heap)
	extractMin()
	fmt.Println("HEAP:", Heap)
	extractMin()
	fmt.Println("HEAP:", Heap)
	extractMin()
	// extractMin()
	// extractMin()
	// extractMin()
	fmt.Println("he nho heap z", Heap)
}
