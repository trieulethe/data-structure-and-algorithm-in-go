package main

import "fmt"

// using recursion
func binarySearch(arr []int, k int, low int, high int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if k == arr[mid] {
		return mid
	} else if k < mid {
		return binarySearch(arr, k, low, mid-1)
	} else {
		return binarySearch(arr, k, mid+1, high)
	}

}

// using loop
func binarySearchLoop(arr []int, k int, low int, high int) int {
	for low < high {
		mid := (low + high) / 2
		if k == arr[mid] {
			return mid
		} else if k < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println("hello world")
	arr := []int{1, 2, 3, 4, 5, 6}
	x := binarySearch(arr, 5, 0, 5)
	println(x)
	x = binarySearchLoop(arr, 3, 0, 5)
	println(x)

}
