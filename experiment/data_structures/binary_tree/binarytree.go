// Package binarytree basic binary tree and related operations
package main

import "fmt"

func main() {
	addAndDelete()
}

func addAndDelete() {
	t := &BTree{nil}
	InOrder(t.Root)
	t.Root = Insert(t.Root, 30)
	t.Root = Insert(t.Root, 20)
	t.Root = Insert(t.Root, 15)
	t.Root = Insert(t.Root, 10)
	t.Root = Insert(t.Root, 12)
	t.Root = Insert(t.Root, 9)
	t.Root = Insert(t.Root, 11)
	t.Root = Insert(t.Root, 17)
	fmt.Print(t.Depth(), "\n")
	InOrder(t.Root)
	fmt.Print("\n")
	t.Root = BstDelete(t.Root, 10)
	InOrder(t.Root)
	fmt.Print("\n")
	t.Root = BstDelete(t.Root, 30)
	InOrder(t.Root)
	fmt.Print("\n")
	t.Root = BstDelete(t.Root, 15)
	InOrder(t.Root)
	fmt.Print("\n")
	t.Root = BstDelete(t.Root, 20)
	InOrder(t.Root)
	fmt.Print("\n")
	fmt.Print(t.Depth(), "\n")
}
func kindsOfOrder() {
	t := BTree{nil}
	t.Root = NewNode(0)
	t.Root.left = NewNode(1)
	t.Root.right = NewNode(2)
	t.Root.left.left = NewNode(3)
	t.Root.left.right = NewNode(4)
	t.Root.right.left = NewNode(5)
	t.Root.right.right = NewNode(6)
	t.Root.right.right.right = NewNode(10)

	InOrder(t.Root)
	fmt.Print("\n")
	PreOrder(t.Root)
	fmt.Print("\n")
	PostOrder(t.Root)
	fmt.Print("\n")
	LevelOrder(t.Root)
	fmt.Print("\n")
	fmt.Print(t.Depth(), "\n")
	var list = AccessNodesByLayer(t.Root)
	fmt.Println("{")
	for i, v := range list {
		for _, v2 := range v {
			fmt.Print(" [", v2, "]")
		}
		if i != len(list)-1 {
			fmt.Print(",")
		}
		fmt.Println()
	}
	fmt.Println("}")
}
