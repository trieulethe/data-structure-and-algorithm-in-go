package main

import "fmt"

type Tree struct {
	root       *Node
	lengthNode int
}

type Node struct {
	index  int
	value  int
	size   int // total node
	rank   int // height of node
	parent *Node
	left   *Node
	right  *Node
}

func (*Node) insertNode(n *Node, index int, v int) bool {
	if v < n.value {
		if n.left == nil {
			n.left = &Node{index, v, 1, n.rank + 1, n, nil, nil}
			n.size++
			return true
		}
		n.size++
		return n.insertNode(n.left, index, v)
	}
	if n.right == nil {
		n.right = &Node{index, v, 1, n.rank + 1, n, nil, nil}
		n.size++
		return true
	}
	n.size++
	return n.insertNode(n.right, index, v)
}

func (*Tree) insertTree(t *Tree, v int) bool {
	length := t.lengthNode

	if t.root == nil {
		t.root = &Node{t.lengthNode, v, 1, 1, nil, nil, nil}
		t.lengthNode++
		return true
	}
	t.lengthNode++
	return t.root.insertNode(t.root, length, v)
}

func loopSearch(n *Node, k int) *Node {
	node := n
	for node != nil {
		if k == node.value {
			return node
		} else if k <= node.value {
			node = node.left
		} else {
			node = node.right
		}
	}
	return node
}

func printNode(n *Node) {
	if n == nil {
		return
	}
	fmt.Println("node: ", n)
	printNode(n.left)
	printNode(n.right)
}

func printTree(t *Tree) {
	if t.root == nil {
		return
	}
	fmt.Println()
	printNode(t.root)
	fmt.Println()
}

func restructure(n *Node) {
	fmt.Println("current node:", n)
	parent := n.parent
	grandParent := n.parent.parent
	var list []*Node
	list = append(list, n.left)
	list = append(list, n.right)
	list = append(list, parent.left)
	list = append(list, parent.right)
	list = append(list, grandParent.left)
	list = append(list, grandParent.right)
	list = append(list, grandParent)
	fmt.Println("list: ", list)
	// t0 := n.left
	// t1 := n.right
	// var t2, t3 *Node
	// if n.value <= parent.value {
	// 	t2 = parent.right
	// } else {
	// 	t2 = parent.left
	// }
	// if parent.value <= grandParent.value {
	// 	t3 = grandParent.right
	// } else {
	// 	t3 = grandParent.left
	// }
	// fmt.Println("parent: ", parent)
	// fmt.Println("grandParent: ", grandParent)
	// fmt.Println("t0: ", t0)
	// fmt.Println("t1: ", t1)
	// fmt.Println("t2: ", t2)
	// fmt.Println("t3: ", t3)
}

func main() {
	t := &Tree{}

	t.insertTree(t, 44)
	t.insertTree(t, 23)
	t.insertTree(t, 17)
	t.insertTree(t, 63)
	t.insertTree(t, 51)
	t.insertTree(t, 33)
	t.insertTree(t, 8)
	t.insertTree(t, 20)
	t.insertTree(t, 74)
	t.insertTree(t, 71)
	t.insertTree(t, 72)
	t.insertTree(t, 48)
	t.insertTree(t, 58)
	t.insertTree(t, 46)
	printTree(t)

	node := loopSearch(t.root, 8)
	// fmt.Println("node:", node)

	restructure(node)

}
