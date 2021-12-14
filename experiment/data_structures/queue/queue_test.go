// Queue Test
// description: based on `geeksforgeeks` description A Queue is a linear structure which follows a particular order in which the operations are performed.
// 	The order is First In First Out (FIFO).
// details:
// 	Queue Data Structure : https://www.geeksforgeeks.org/queue-data-structure/
//  Queue (abstract data type) : https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
// author [Milad](https://github.com/miraddo)
// see queuearray.go, queuelinkedlist.go, queuelinkedlistwithlist.go

package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {

	// Handle Queue Linked List
	t.Run("Test Queue Linked List", func(t *testing.T) {

		t.Run("Test EnQueue", func(t *testing.T) {
			var newQueue Queue
			newQueue.enqueue(2)
			newQueue.enqueue(3)
			newQueue.enqueue(4)
			newQueue.enqueue(45)

			if newQueue.frontQueue() != 2 && newQueue.backQueue() != 45 {
				t.Errorf("Test EnQueue is wrong the result must be %v and %v but got %v and %v", 2, 45, newQueue.frontQueue(), newQueue.backQueue())
			}

		})

		t.Run("Test DeQueue", func(t *testing.T) {
			var newQueue Queue
			newQueue.enqueue(2)
			newQueue.enqueue(3)
			newQueue.enqueue(4)

			newQueue.dequeue()
			if newQueue.dequeue() != 3 {
				t.Errorf("Test DeQueue is wrong the result must be %v but got %v", 3, newQueue.dequeue())
			}

			//fmt.Println(newQueue.show())
		})

		t.Run("Test Queue isEmpty", func(t *testing.T) {
			var newQueue Queue
			if newQueue.isEmpty() != true {
				t.Errorf("Test Queue isEmpty is wrong the result must be %v but got %v", true, newQueue.isEmpty())
			}

			newQueue.enqueue(3)
			newQueue.enqueue(4)

			if newQueue.isEmpty() != false {
				t.Errorf("Test Queue isEmpty is wrong the result must be %v but got %v", false, newQueue.isEmpty())
			}
		})

		t.Run("Test Queue Length", func(t *testing.T) {
			var newQueue Queue
			if newQueue.len() != 0 {
				t.Errorf("Test Queue Length is wrong the result must be %v but got %v", 0, newQueue.len())
			}

			newQueue.enqueue(3)
			newQueue.enqueue(4)
			newQueue.dequeue()
			newQueue.enqueue(22)
			newQueue.enqueue(99)
			newQueue.dequeue()
			newQueue.dequeue()

			if newQueue.len() != 1 {
				t.Errorf("Test Queue Length is wrong the result must be %v but got %v", 1, newQueue.len())
			}

		})
	})

	// Handle Queue Array
	t.Run("Test Queue Array", func(t *testing.T) {
		t.Run("Test EnQueue", func(t *testing.T) {
			EnQueue(2)
			EnQueue(23)
			EnQueue(45)
			EnQueue(66)

			if FrontQueue() != 2 && BackQueue() != 66 {
				t.Errorf("Test EnQueue is wrong the result must be %v and %v but got %v and %v", 2, 66, FrontQueue(), BackQueue())
			}

		})

		t.Run("Test DeQueue", func(t *testing.T) {
			EnQueue(2)
			EnQueue(23)
			EnQueue(45)
			EnQueue(66)

			DeQueue()
			DeQueue()

			if DeQueue() != 45 {
				t.Errorf("Test DeQueue is wrong the result must be %v but got %v", 45, DeQueue())
			}
		})

		ListQueue = []interface{}{}

		t.Run("Test Queue isEmpty", func(t *testing.T) {

			if IsEmptyQueue() != true {
				t.Errorf("Test Queue isEmpty is wrong the result must be %v but got %v", true, IsEmptyQueue())
			}

			EnQueue(3)
			EnQueue(4)

			if IsEmptyQueue() != false {
				t.Errorf("Test Queue isEmpty is wrong the result must be %v but got %v", false, IsEmptyQueue())
			}
		})

		ListQueue = []interface{}{}
		t.Run("Test Queue Length", func(t *testing.T) {
			if LenQueue() != 0 {
				t.Errorf("Test Queue Length is wrong the result must be %v but got %v", 0, LenQueue())
			}

			EnQueue(3)
			EnQueue(4)
			DeQueue()
			EnQueue(22)
			EnQueue(99)
			DeQueue()
			DeQueue()

			if LenQueue() != 1 {
				t.Errorf("Test Queue Length is wrong the result must be %v but got %v", 1, LenQueue())
			}

		})

	})
}
