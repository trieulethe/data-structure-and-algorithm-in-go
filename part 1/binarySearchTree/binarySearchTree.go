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
	// addrPointer int
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

func (*Node) binarySearchNode(n *Node, k int) bool {
	if n == nil {
		return false
	}
	if k == n.value {
		return true
	} else if k < n.value {
		return n.binarySearchNode(n.left, k)
	} else {
		return n.binarySearchNode(n.right, k)
	}
}

func (*Tree) searchValue(t *Tree, k int) bool {
	if t.root == nil {
		return false
	}
	return t.root.binarySearchNode(t.root, k)
}

func (*Node) searchIndex(n *Node, id int) **Node {
	// fmt.Println()
	fmt.Println("node >> search: ", n)
	// fmt.Println(reflect.TypeOf(n))
	// fmt.Println(reflect.TypeOf(&n))
	fmt.Println("node >> search: ", &n)
	// fmt.Println()

	if n == nil {
		return nil
	}
	if id == n.index {
		fmt.Println("n index: ", n)
		fmt.Println("n index: ", &n)
		return &n
	}
	// else if  {
	// 	return n.searchIndex(n.left, id)
	// } else {
	// 	return n.searchIndex(n.right, id)
	// }

	node := n.searchIndex(n.left, id)
	if node == nil {
		node = n.searchIndex(n.right, id)
	}
	fmt.Println("node finded return: ", node)
	*node = nil
	// fmt.Println("node finded return: ", &node)
	// fmt.Println("node finded return: ", *node)
	// println(">> : ", &node)
	return node
}

func (*Node) checkTypeNode(n *Node) string {
	if n.left == nil && n.right == nil {
		return "leaf"
	} else if (n.left == nil && n.right != nil) || (n.left != nil && n.right == nil) {
		return "child"
	} else {
		return "internal"
	}
}

func (*Node) removeLeaf(n *Node) *Node {
	n = nil
	return n
}

func (*Node) removeChild(n *Node) *Node {
	if n.left == nil {
		// parent := n.parent
		n = n.right
		n.right = nil
		fmt.Println("delete child n right: ", n)
	} else {
		n = n.left
		n.left = nil
	}

	return n
}

func (*Node) removeInternal(n *Node) *Node {
	temp := n.getElementSideLeft(n)
	fmt.Println("node temp: ", temp)
	// temp.removeChild(n)
	return temp.removeChild(n)

	// n = temp
	// fmt.Println("node copy temp: ", n)
	// return n
}

func (*Node) getElementSideLeft(n *Node) *Node {
	if n.left == nil {
		return n
	}
	return n.getElementSideLeft(n.left)
}

func (*Tree) remove(t *Tree, id int) *Node {
	if id == 0 {
		fmt.Println("cureent can not remove root")
	}
	fmt.Println("root: ", t.root)
	fmt.Println("address root: ", *t.root)
	fmt.Println("address root: ", &t.root)
	node := t.root.searchIndex(t.root, id)
	fmt.Println("finded node: ", node)
	fmt.Println("finded node: ", *node)
	// println(node)
	// p := node
	// // *p = nil
	// fmt.Println("p: ", p)
	// fmt.Println(reflect.TypeOf(p))
	// fmt.Println(p)
	// fmt.Println(reflect.TypeOf(*p))
	// fmt.Println(*p)
	// fmt.Println(reflect.TypeOf(**p))
	// fmt.Println(**p)
	// *p = nil

	// checkType := node.checkTypeNode(node)
	// fmt.Println("check type node: ", checkType)

	// if checkType == "leaf" {
	// 	node = node.removeLeaf(node)
	// } else if checkType == "child" {
	// 	node = node.removeChild(node)
	// } else {
	// 	node = node.removeInternal(node)
	// 	fmt.Println("node removed internal: ", node)
	// }
	return nil
}

func (*Node) printNode(n *Node) {
	if n == nil {
		return
	}
	fmt.Println("node: ", n)
	n.printNode(n.left)
	n.printNode(n.right)
}

func (*Tree) printTree(t *Tree) {
	if t.root == nil {
		return
	}
	fmt.Println()
	t.root.printNode(t.root)
	fmt.Println()
}

func main() {
	t := Tree{}
	fmt.Println("tree > : ", t)

	i := t.insertTree(&t, 34)
	fmt.Println("tree > : ", t, i)

	i = t.insertTree(&t, 17)
	fmt.Println("tree > : ", t, i)

	i = t.insertTree(&t, 25)
	fmt.Println("tree > : ", t, i)

	i = t.insertTree(&t, 66)
	fmt.Println("tree > : ", t, i)

	i = t.insertTree(&t, 50)
	fmt.Println("tree > : ", t, i)

	i = t.insertTree(&t, 71)
	fmt.Println("tree > : ", t, i)

	i = t.insertTree(&t, 68)
	fmt.Println("tree > : ", t, i)

	i = t.insertTree(&t, 75)
	fmt.Println("tree > : ", t, i)

	fmt.Println()

	var n = Node{}
	fmt.Println("node: ", n)
	fmt.Println("&node: ", &n)
	p := &n
	fmt.Println("p: ", p)
	fmt.Println("*&node: ", *&n)
	fmt.Println("*&node: ", *&n)

	// t.printTree(&t)
	fmt.Println()
	// fmt.Println("t: ", t)
	// fmt.Println(reflect.TypeOf(t))
	// fmt.Println("&t: ", &t)
	// // fmt.Println(reflect.TypeOf(&t))
	// fmt.Println("t: ", t.root)
	// fmt.Println(reflect.TypeOf(t.root))
	// fmt.Println("&t: ", &t.root)

	// fmt.Println(reflect.TypeOf(&t.root))
	// fmt.Println("*t: ", *t.root)
	// fmt.Println(reflect.TypeOf(*t.root))
	// fmt.Println("&*t: ", &*t.root)
	// fmt.Println(reflect.TypeOf(&*t.root))
	// fmt.Println()
	// p := &t.root
	// *p = nil
	// t.remove(&t, 1)

	// fmt.Println("root : ", t.root)
	// fmt.Println(reflect.TypeOf(t.root))
	// fmt.Println("root : ", &t.root)
	// fmt.Println(reflect.TypeOf(&t.root))
	// fmt.Println("root : ", *t.root)
	// fmt.Println(reflect.TypeOf(*t.root))
	// fmt.Println(reflect.TypeOf(*&t.root))
	// // *&t.root = nil
	// fmt.Println("root : ", t.root)

	// fmt.Println()

	// test := &t.root
	// fmt.Println("test: ", test)
	// fmt.Println("test: ", &test)
	// fmt.Println("test: ", *test)
	// *test = nil
	// fmt.Println("root : ", t.root)

	// test1 := test
	// fmt.Println("test1: ", test1)
	// // *test1 = nil
	// fmt.Println("test1: ", test1)
	// // t.root = nil
	// fmt.Println("test change: ", test)
	// // test.value = 23
	// // test = &Node{}
	// // test.left = nil
	// fmt.Println("test changeed: ", test)
	// fmt.Println("root : ", t)

	// t.printTree(t)

	// n := t.root.searchIndex(t.root, 1)
	// fmt.Println("index  1: ", n)

	// n = t.root.searchIndex(t.root, 1)
	// fmt.Println("index  1: ", n)

	t.printTree(&t)

	// n = n.getElementSideLeft(t.root)
	// fmt.Println("node end of side left: ", n)

	// if n.parent.left == n {
	// 	fmt.Println("compare left true")
	// }
	// if n.parent.right == n {
	// 	fmt.Println("compare right true")
	// }

	// check := t.searchValue(t, 4)
	// fmt.Println("check search", check)

	// t.remove(t, 1)
	// check = t.searchValue(t, 4)
	// fmt.Println("check search", check)

}
