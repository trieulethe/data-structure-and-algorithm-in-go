package main

import "fmt"

type Node struct {
	key    int
	left   *Node
	right  *Node
	height int
}

func newNode(key int) *Node {
	return &Node{key, nil, nil, 1}
}

func print(n *Node) {
	if n == nil {
		return
	}
	fmt.Println("node >>", n)
	print(n.left)
	print(n.right)
	return
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}

func getBalance(n *Node) int {
	if n == nil {
		return 0
	}
	return height(n.left) - height(n.right)
}

func rightRotate(y *Node) *Node {
	// fmt.Println(" right rotate y", y)
	x := y.left
	T2 := x.right // y.left.right

	x.right = y
	y.left = T2

	y.height = max(height(y.left), height(y.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1

	return x
}

func leftRotate(y *Node) *Node {
	x := y.right
	T2 := x.left

	x.left = y
	y.right = T2

	y.height = max(height(y.left), height(y.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1

	return x
}

func insert(n *Node, key int) *Node {
	if n == nil {
		return newNode(key)
	}

	if key < n.key {
		n.left = insert(n.left, key)
	} else if key > n.key {
		n.right = insert(n.right, key)
	} else {
		return n
	}
	// update height of node ancestor
	n.height = 1 + max(height(n.left), height(n.right))

	// get balance of current node
	balance := getBalance(n)

	// left left case
	if balance > 1 && key < n.left.key {
		n = rightRotate(n)
	}

	// right right case
	if balance < -1 && key > n.right.key {
		n = leftRotate(n)
	}

	// Left Right Case
	if balance > 1 && key > n.left.key {
		n.left = leftRotate(n.left)
		return rightRotate(n)
	}

	// Right Left Case
	if balance < -1 && key < n.right.key {
		n.right = rightRotate(n.right)
		return leftRotate(n)
	}

	return n

}

func minValueNode(n *Node) *Node {
	current := n
	for current.left != nil {
		current = current.left
	}
	return current
}

func main() {
	n := newNode(10)
	fmt.Println("new node:", n)
	n = insert(n, 11)
	// n = insert(n, 9)
	fmt.Println("")
	print(n)
	n = leftRotate(n)
	// n = insert(n, 20)
	// n = insert(n, 30)
	// n = insert(n, 40)
	// n = insert(n, 50)
	// n = insert(n, 25)
	fmt.Println("")
	print(n)
}
