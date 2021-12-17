package stack

import "testing"

func TestStackLinkedList(t *testing.T) {
	var newStack Stack

	newStack.push(1)
	newStack.push(2)
	t.Run("Stack push", func(t *testing.T) {
		result := newStack.show()
		expected := []interface{}{2, 1}
		for x := range result {
			if result[x] != expected[x] {
				t.Errorf("Stack push is not work, got %v but expected %v", result, expected)
			}
		}
	})
	t.Run("Stack isEmpty", func(t *testing.T) {
		if newStack.isEmpty() {
			t.Error("Stack isEmpty is returned true but expected false", newStack.isEmpty())
		}
	})


	t.Run("Stack Length", func(t *testing.T) {
		if newStack.len() != 2 {
			t.Error("Stack Length should be 2 but got", newStack.len())
		}
	})

	newStack.pop()
	pop := newStack.pop()

	t.Run("Stack Pop", func(t *testing.T) {
		if pop != 1 {
			t.Error("Stack Pop should return 1 but is returned", pop)
		}

	})

	newStack.push(52)
	newStack.push(23)
	newStack.push(99)
	t.Run("Stack Peak", func(t *testing.T) {
		if newStack.peak() != 99 {
			t.Error("Stack Peak should return 99 but got ", newStack.peak())
		}
	})
}