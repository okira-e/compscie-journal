package containers_test

import (
	"data_structures/containers"
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("Test Queue", func(t *testing.T) {
		q := containers.NewQueue[string]()

		q.Enqueue("a")
		q.Enqueue("b")
		q.Enqueue("c")

		if q.GetLen() != 3 {
			t.Error("Expected 3, got ", q.GetLen())
		}

		if r, _ := q.Peek(); r != "a" {
			t.Error("Expected a, got ", r)
		}

		if r, _ := q.Deque(); r != "a" {
			t.Error("Expected a, got ", r)
		}

		if r, _ := q.Deque(); r != "b" {
			t.Error("Expected b, got ", r)
		}

		q.Enqueue("d")

		if q.GetLen() != 2 {
			t.Error("Expected 2, got ", q.GetLen())
		}

		if r, _ := q.Deque(); r != "c" {
			t.Error("Expected c, got ", r)
		}

		if r, _ := q.Deque(); r != "d" {
			t.Error("Expected d, got ", r)
		}
	})
}
