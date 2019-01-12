package main

import (
	"fmt"
)

type Tree struct {
	root       *Node
	lengthNode int
}

type Node struct {
	index  int
	value  int
	parent *Node
	left   *Node
	right  *Node
}

func (*Node) insertNode(n *Node, index int, v int) bool {
	if v < n.value {
		if n.left == nil {
			n.left = &Node{index, v, n, nil, nil}
			return true
		}
		return n.insertNode(n.left, index, v)
	}
	if n.right == nil {
		n.right = &Node{index, v, n, nil, nil}
		return true
	}
	return n.insertNode(n.right, index, v)
}

func (*Tree) insertTree(t *Tree, v int) bool {
	length := t.lengthNode

	if t.root == nil {
		t.root = &Node{t.lengthNode, v, nil, nil, nil}
		t.lengthNode++
		return true
	}
	t.lengthNode++
	return t.root.insertNode(t.root, length, v)
}

func binarySearchNode(n *Node, k int) bool {
	if n == nil {
		return false
	}
	if k == n.value {
		return true
	} else if k < n.value {
		return binarySearchNode(n.left, k)
	} else {
		return binarySearchNode(n.right, k)
	}
}

func searchValue(t *Tree, k int) bool {
	if t.root == nil {
		return false
	}
	return binarySearchNode(t.root, k)
}

func searchIndex(n *Node, id int) *Node {
	if n == nil {
		return nil
	}

	if id == n.index {
		return n
	}

	node := searchIndex(n.left, id)
	if node == nil {
		node = searchIndex(n.right, id)
	}

	return node
}

func checkTypeNode(n *Node) string {
	if n.left == nil && n.right == nil {
		return "leaf"
	} else if (n.left == nil && n.right != nil) || (n.left != nil && n.right == nil) {
		return "child"
	} else {
		return "internal"
	}
}

func checkParent(n *Node) string {
	if n.parent.left == n {
		return "left"
	}
	return "right"
}

func removeLeaf(n *Node) {
	parent := n.parent
	if parent.left == n {
		parent.left = nil
		n = &Node{}
		return
	}
	parent.right = nil
	n = &Node{}
	return
}

func removeChild(n *Node) {
	fmt.Println("node n:", n)
	fmt.Println("node *n:", *n)
	temp := *n
	if n.left == nil {
		*n = *n.right
		n.parent = temp.parent
		return
	}

	*n = *n.left
	n.parent = temp.parent
	return
}

func removeInternal(n *Node) {
	fmt.Println("node internal:", n)
	parent := n.parent
	fmt.Println("parent node:", parent)
	fmt.Println("parent node:", parent.right)
	fmt.Println("parent node:", *parent.right)
	nChildLeft := n.getElementSideLeft(n.right)
	fmt.Println("left node:", nChildLeft)
	n.value = nChildLeft.value
	n.index = nChildLeft.index
	fmt.Println("node changed:", n)
	removeLeaf(nChildLeft)
	return
}

func (*Node) getElementSideLeft(n *Node) *Node {
	if n.left == nil {
		return n
	}
	return n.getElementSideLeft(n.left)
}

func (*Tree) remove(t *Tree, id int) {
	if id == 0 {
		fmt.Println("cureent can not remove root")
	}
	node := searchIndex(t.root, id)
	checkT := checkTypeNode(node)
	if checkT == "leaf" {
		removeLeaf(node)
	} else if checkT == "child" {
		removeChild(node)
	} else {
		removeInternal(node)
	}

	return
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

func rangeQuery(n *Node, k1 int, k2 int) {
	if n == nil {
		return
	}

	if k1 <= n.value && n.value <= k2 {
		fmt.Println("n in range", n)
		rangeQuery(n.left, k1, k2)
		rangeQuery(n.right, k1, k2)
		return
	} else if n.value < k1 {
		rangeQuery(n.right, k1, k2)
	} else if k2 < n.value {
		rangeQuery(n.left, k1, k2)
	}
}

func main() {
	t := &Tree{}

	t.insertTree(t, 34)
	t.insertTree(t, 17)
	t.insertTree(t, 25)
	t.insertTree(t, 66)
	t.insertTree(t, 50)
	t.insertTree(t, 71)
	t.insertTree(t, 68)
	t.insertTree(t, 75)
	printTree(t)

	fmt.Println()

	rangeQuery(t.root, 55, 72)
	fmt.Println()

	t.remove(t, 3)
	t.remove(t, 1)
	t.remove(t, 2)
	t.remove(t, 4)

	printTree(t)
	fmt.Println()

}
