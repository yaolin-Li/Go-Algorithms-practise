package doubly_linkedlist

import "fmt"

type Node struct {
	val  int
	next *Node
	prev *Node
}

type DoubleLinkedList struct {
	head *Node
}

func NewNode(val int) *Node {
	n := &Node{}
	n.val = val
	n.next = nil
	n.prev = nil
	return n
}

func (ll *DoubleLinkedList) AddAtBeg(val int) {
	n := NewNode(val)
	n.next = ll.head

	if ll.head != nil {
		ll.head.prev = n
	}

	ll.head = n
}

func (ll *DoubleLinkedList) AddAtEnd(val int) {
	n := NewNode(val)

	if ll.head == nil {
		ll.head = n
		return
	}

	cur := ll.head
	for ; cur.next != nil; cur = cur.next {
	}
	cur.next = n
	n.prev = cur
}

func (ll *DoubleLinkedList) DelAtBeg() int {
	if ll.head == nil {
		return -1
	}

	cur := ll.head
	ll.head = cur.next

	if ll.head != nil {
		ll.head.prev = nil
	}

	return cur.val
}

// DetAtEnd Delete a node at the end of the linkedlist
func (ll *DoubleLinkedList) DelAtEnd() int {
	// no item
	if ll.head == nil {
		return -1
	}

	// only one item
	if ll.head.next == nil {
		return ll.DelAtBeg()
	}

	// more than one, go to second last
	cur := ll.head
	for ; cur.next.next != nil; cur = cur.next {
	}

	retval := cur.next.val
	cur.next = nil
	return retval
}

// Count Number of nodes in the linkedlist
func (ll *DoubleLinkedList) Count() int {
	var ctr int = 0

	for cur := ll.head; cur != nil; cur = cur.next {
		ctr += 1
	}

	return ctr
}

// Reverse Reverse the order of the linkedlist
func (ll *DoubleLinkedList) Reverse() {
	var prev, next *Node
	cur := ll.head

	for cur != nil {
		next = cur.next
		cur.next = prev
		cur.prev = next
		prev = cur
		cur = next
	}

	ll.head = prev
}

// Display diplay the linked list
func (ll *DoubleLinkedList) Display() {
	for cur := ll.head; cur != nil; cur = cur.next {
		fmt.Print(cur.val, " ")
	}

	fmt.Print("\n")
}

// DisplayReverse Display the linkedlist in reverse order
func (ll *DoubleLinkedList) DisplayReverse() {
	if ll.head == nil {
		return
	}
	var cur *Node
	for cur = ll.head; cur.next != nil; cur = cur.next {
	}

	for ; cur != nil; cur = cur.prev {
		fmt.Print(cur.val, " ")
	}

	fmt.Print("\n")
}
