package singly_linkedlist

import "fmt"

type snode struct {
	Val  interface{}
	Next *snode
}

type SingleLinkedList struct {
	length int
	Head   *snode
}

func CreateList() *SingleLinkedList {
	return &SingleLinkedList{}
}

func NewNode(val interface{}) *snode {
	return &snode{val, nil}
}

func (ll *SingleLinkedList) AddAtBeg(val interface{}) {
	n := NewNode(val)
	n.Next = ll.Head
	ll.Head = n
	ll.length++
}

func (ll *SingleLinkedList) AddAtEnd(val int) {
	n := NewNode(val)

	if ll.Head == nil {
		ll.Head = n
		ll.length++
		return
	}

	cur := ll.Head
	for ; cur.Next != nil; cur = cur.Next {
	}

	cur.Next = n
	ll.length++
}

func (ll *SingleLinkedList) DelAtBeg() interface{} {
	if ll.Head == nil {
		return -1
	}

	cur := ll.Head
	ll.Head = cur.Next
	ll.length--

	return cur.Val
}

func (ll *SingleLinkedList) DelAtEnd() interface{} {
	if ll.Head == nil {
		return -1
	}

	if ll.Head.Next == nil {
		return ll.DelAtBeg()
	}

	cur := ll.Head

	for ; cur.Next.Next != nil; cur = cur.Next {
	}

	retval := cur.Next.Val
	cur.Next = nil
	ll.length--
	return retval
}

func (ll *SingleLinkedList) Count() int {
	return ll.length
}

func (ll *SingleLinkedList) Reverse() {
	var prev, Next *snode
	cur := ll.Head

	for cur != nil {
		Next = cur.Next
		cur.Next = prev
		prev = cur
		cur = Next
	}

	ll.Head = prev
}

func (ll *SingleLinkedList) Display() {
	for cur := ll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}

	fmt.Print("\n")
}
