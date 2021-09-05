package main

import "fmt"

type BTree struct {
	Root *Node
}

func PreOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(root.val, " ")
	PreOrder(root.left)
	PreOrder(root.right)
}

func InOrder(root *Node) {
	if root == nil {
		return
	}
	InOrder(root.left)
	fmt.Print(root.val, " ")
	InOrder(root.right)
}

func PostOrder(root *Node) {
	if root == nil {
		return
	}
	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Print(root.val, " ")
}

func LevelOrder(root *Node) {
	var q []*Node
	var n *Node
	q = append(q, root)

	for len(q) != 0 {
		n, q = q[0], q[1:]
		fmt.Print(n.val, " ")
		if n.left != nil {
			q = append(q, n.left)
		}
		if n.right != nil {
			q = append(q, n.right)
		}
	}
}

func AccessNodesByLayer(root *Node) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var q []*Node
	var n *Node
	var idx = 0
	q = append(q, root)

	for len(q) != 0 {
		res = append(res, []int{})
		qLen := len(q)
		for i := 0; i < qLen; i++ {
			n, q = q[0], q[1:]
			res[idx] = append(res[idx], n.val)
			if n.left != nil {
				q = append(q, n.left)
			}
			if n.right != nil {
				q = append(q, n.right)
			}
		}
		idx++
	}
	return res
}

func (t *BTree) Depth() int {
	return calculateDepth(t.Root, 0)
}

func calculateDepth(n *Node, depth int) int {
	if n == nil {
		return depth
	}
	return Max(calculateDepth(n.left, depth+1), calculateDepth(n.right, depth+1))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Insert(root *Node, val int) *Node {
	if root == nil {
		return NewNode(val)
	}
	if val < root.val {
		root.left = Insert(root.left, val)
	} else {
		root.right = Insert(root.right, val)
	}
	return root
}

func BstDelete(root *Node, val int) *Node {
	if root == nil {
		return nil
	}
	if val < root.val {
		root.left = BstDelete(root.left, val)
	} else if val > root.val {
		root.right = BstDelete(root.right, val)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}
		n := root.right
		d := InOrderSuccessor(n)
		d.left = root.left
		return root.right
	}
	return root
}

// InOrderSuccessor Goes to the left
func InOrderSuccessor(root *Node) *Node {
	cur := root
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}
