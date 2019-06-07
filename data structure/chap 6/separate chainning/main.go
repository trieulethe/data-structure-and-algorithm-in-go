package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type HashTable struct {
	listArray []*Node
	size      int
}

func (h *HashTable) init() {
	h.size = 101
	h.listArray = make([]*Node, h.size)

	for i := 0; i < h.size; i++ {
		h.listArray[i] = nil
	}
}

func (h *HashTable) computeHash(key int) int {
	hashValue := key
	return hashValue % h.size
}

func (h *HashTable) add(value int) {
	index := h.computeHash(value)
	temp := new(Node)
	temp.value = value
	fmt.Println("node:", h.listArray[index])
	temp.next = h.listArray[index]
	fmt.Println("temp node", temp)
	h.listArray[index] = temp
}

func (h *HashTable) remove(value int) bool {
	index := h.computeHash(value)
	var nextNode, head *Node
	head = h.listArray[index]
	if head != nil && head.value == value {
		h.listArray[index] = head.next
		return true
	}
	for head != nil {
		nextNode = head.next
		if nextNode != nil && nextNode.value == value {
			head.next = nextNode.next
			return true
		}
		head = nextNode
	}
	return false
}

func (h *HashTable) find(value int) bool {
	index := h.computeHash(value)
	head := h.listArray[index]
	for head != nil {
		if head.value == value {
			return true
		}
		head = head.next
	}
	return false
}

func (h *HashTable) print() {
	for i := 0; i < h.size; i++ {
		head := h.listArray[i]
		if head != nil {
			fmt.Print("\nValues at index :: ", i, " Are :: ")
		}
		for head != nil {
			fmt.Print(head.value, " ")
			head = head.next
		}
	}
	fmt.Println()
}

func main() {
	ht := new(HashTable)
	ht.init()
	ht.add(89)
	ht.add(89)
	ht.add(89)
	ht.add(82)
	ht.add(83)
	ht.print()
	fmt.Println("89 found : ", ht.find(89))
}
