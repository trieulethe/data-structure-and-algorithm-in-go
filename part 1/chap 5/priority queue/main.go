package main

import "fmt"

const MaxStack = 10

type priorityQueue struct {
	stacks [MaxStack]int
	size   int
}

func (p *priorityQueue) isEmpty() bool {
	return p.size == 0
}

func (p *priorityQueue) isFull() bool {
	return p.size == 10
}

func (p *priorityQueue) getMax() int {
	return p.stacks[0]
}

func (p *priorityQueue) getMin() int {
	return p.stacks[p.size-1]
}

func (p *priorityQueue) recoverQueue(element int) {
	// fmt.Println("queue", p)
	for i := p.size; i > 0; i-- {
		// fmt.Println("value:", p.stacks[i])
		if p.stacks[i] > p.stacks[i-1] {
			p.stacks[i], p.stacks[i-1] = p.stacks[i-1], p.stacks[i]
		} else {
			break
		}
	}
	return
}

func (p *priorityQueue) insert(value int) {
	p.stacks[p.size] = value
	// fmt.Println("p:", p)
	if p.size > 0 && (p.stacks[p.size] > p.stacks[p.size-1]) {
		p.recoverQueue(p.size)
	}
	p.size++
	return
}

func (p *priorityQueue) removeMin() {
	p.stacks[p.size-1] = 0
	p.size--
	return
}

func main() {
	arr := &priorityQueue{}
	arr.insert(2)
	arr.insert(3)
	arr.insert(1)
	arr.insert(5)
	arr.insert(-1)
	arr.insert(4)
	fmt.Println("arr:", arr)
	fmt.Println("max:", arr.getMax())
	fmt.Println("min:", arr.getMin())
	arr.removeMin()
	fmt.Println("arr:", arr)

}
