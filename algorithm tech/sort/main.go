package main

import "fmt"

func selectionSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		s := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[s] {
				s = j
			}
		}
		if i != s {
			arr[s], arr[i] = arr[i], arr[s]
		}
	}
	return arr
}

func insertSort(arr []int) []int {
	length := len(arr)
	for i := 1; i < length; i++ {
		x := arr[i]
		j := i
		for j > 0 && x < arr[j-1] {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = x
	}
	return arr
}

func main() {
	arr := []int{2, 1, 4, 3, 5, 7, 6, 9, 8}
	// arr = selectionSort(arr)
	arr = insertSort(arr)
	fmt.Println("arr:", arr)
}
