package main

import "fmt"

const MAX = 10

var intArray [MAX]int
var itemCount = 0

func peek() int {
	return intArray[itemCount-1]
}

func isEmpty() bool {
	return itemCount == 0
}

func isFull() bool {
	return itemCount == MAX
}

func size() int {
	return itemCount
}

func insert(data int) {
	i := 0

	if !isFull() {
		// if queue empty, insert the data
		if itemCount == 0 {
			intArray[itemCount] = data
			itemCount++
		} else {
			// start from the right end of the queue
			for i = itemCount - 1; i >= 0; i-- {
				if data > intArray[i] {
					intArray[i+1] = intArray[i]
				} else {
					break
				}
			}

			// insert the data
			intArray[i+1] = data
			itemCount++
		}
	}
}

func removeData() int {
	temp := intArray[itemCount]
	fmt.Println("temp:", temp)
	itemCount--
	return temp
}

func main() {
	insert(3)
	insert(5)
	insert(9)
	insert(1)
	insert(12)
	for _, v := range intArray {
		fmt.Println("value:", v)
	}
	num := removeData()
	fmt.Println("num:", num)
}
