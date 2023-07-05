package containers_test

import (
	"data_structures/containers"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("Test Stack", func(t *testing.T) {
		s := containers.NewStack[int]()

		s.Push(1)
		s.Push(2)
		s.Push(3)

		if s.GetLen() != 3 {
			t.Error("Expected 3, got ", s.GetLen())
		}

		if r, _ := s.Peek(); r != 3 {
			t.Error("Expected 3, got ", r)
		}

		if r, _ := s.Pop(); r != 3 {
			t.Error("Expected 3, got ", r)
		}

		if r, _ := s.Pop(); r != 2 {
			t.Error("Expected 2, got ", r)
		}

		s.Push(4)

		if s.GetLen() != 2 {
			t.Error("Expected 2, got ", s.GetLen())
		}

		if r, _ := s.Pop(); r != 4 {
			t.Error("Expected 4, got ", r)
		}

		if r, _ := s.Pop(); r != 1 {
			t.Error("Expected 1, got ", r)
		}
	})
}
