package queue

import "fmt"


type Node struct {
	Data interface{}
	Next *Node
}

type Queue struct{
	head *Node
	tail *Node
	length int
}

func (ll *Queue) enqueue(n interface{}) {
	var newNode Node
	newNode.Data = n

	if ll.tail != nil {
		ll.tail.Next = &newNode
	}

	ll.tail = &newNode

	if ll.head == nil {
		ll.head = &newNode
	}
	ll.length++
}

func (ll *Queue) dequeue() interface{} {
	if ll.isEmpty() {
		return -1
	}

	data := ll.head.Data

	ll.head = ll.head.Next

	if ll.head == nil {
		ll.tail = nil
	}

	ll.length--
	return data
}

func (ll *Queue) isEmpty() bool {
	return ll.length == 0
}

func (ll *Queue) len() int{
	return ll.length
}

func (ll *Queue) frontQueue() interface{} {
	return ll.head.Data
}

func (ll *Queue) backQueue() interface{} {
	return ll.tail.Data
}

func (ll *Queue) show() {
	p := ll.head
	for p != nil {
		fmt.Println(p.Data.(int))
		p = p.Next
	}
}