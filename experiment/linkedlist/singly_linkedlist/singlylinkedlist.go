package singly_linkedlist

import (
	"errors"
	"fmt"
)

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

// ReversePartition Reverse the linked list from the ath to the bth node
func (ll *SingleLinkedList) ReversePartition(left, right int) error {
	err := ll.CheckRangeFromIndex(left, right)
	if err != nil {
		return err
	}
	tmpNode := NewNode(-1)
	tmpNode.Next = ll.Head
	pre := tmpNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	ll.Head = tmpNode.Next
	return nil
}
func (ll *SingleLinkedList) CheckRangeFromIndex(left, right int) error {
	if left > right {
		return errors.New("left boundary must smaller than right")
	} else if left < 1 {
		return errors.New("left boundary starts from the first node")
	} else if right > ll.length {
		return errors.New("right boundary cannot be greater than the length of the linked list")
	}
	return nil
}
func (ll *SingleLinkedList) Display() {
	for cur := ll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}

	fmt.Print("\n")
}
