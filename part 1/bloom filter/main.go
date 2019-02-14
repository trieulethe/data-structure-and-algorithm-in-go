package main

import "fmt"

const b = 31

var HashTable [256]bool

func hash1(value string) int {
	hash := 0
	n := len(value)
	for i := 0; i < len(value); i++ {
		hash = hash + int(value[i]) + b*n
		n--
	}
	return hash % 256
}

func hash2(value string) int {
	hash := 0
	for i := 0; i < len(value); i++ {
		hash = hash + int(value[i])
	}
	return hash % 256
}

func hash3(value string) int {
	hash := 0
	for i := 0; i < len(value); i++ {
		hash = hash + int(value[i])<<5
	}
	return hash % 256
}

func getIndexHash(data string) (int, int, int) {
	h1 := hash1(data)
	h2 := hash2(data)
	h3 := hash3(data)
	return h1, h2, h3
}

func insert(data string) bool {
	h1, h2, h3 := getIndexHash(data)
	HashTable[h1] = true
	HashTable[h2] = true
	HashTable[h3] = true
	return true
}

func find(data string) bool {
	i1, i2, i3 := getIndexHash(data)
	h1 := HashTable[i1]
	h2 := HashTable[i2]
	h3 := HashTable[i3]
	if h1 == false || h2 == false || h3 == false {
		return false
	}
	return true
}

func main() {
	// hash1 := hash1("hello worlds")
	// fmt.Println("hash1:", hash1)

	// hash2 := hash2("hello worlds")
	// fmt.Println("hash2:", hash2)

	// hash33 := hash3("hello worlds")
	// fmt.Println("hash3:", hash33)
	insert("hello")
	check := find("hello")
	fmt.Println("check", check)
	check = find("hellos")
	fmt.Println("check", check)
}
