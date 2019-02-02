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
	for i := p.size; i > 0; i-- {
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
	if p.size > 0 && (p.stacks[p.size] > p.stacks[p.size-1]) {
		p.recoverQueue(p.size)
	}
	p.size++
	return
}

func (p *priorityQueue) removeMin() int {
	e := p.stacks[p.size-1]
	p.stacks[p.size-1] = 0
	p.size--
	return e
}

// priority queue sort
func pqSort(c []int, q priorityQueue) []int {
	for i := 0; i < len(c); i++ {
		q.insert(c[i])
	}
	for i := 0; i < len(c); i++ {
		e := q.removeMin()
		c[i] = e
	}
	return c
}

func main() {
	q := priorityQueue{}
	c := []int{3, 2, 1, 5, 7, 6, 1}
	c = pqSort(c, q)
	fmt.Println("c:", c)
}
