package containers_test

import (
	"data_structures/containers"
	"testing"
)

func TestLinkedList(t *testing.T) {
	t.Run("Test LinkedList", func(t *testing.T) {
		lst := containers.NewLinkedList[string]()
		lst.Append("a")
		lst.Append("b")
		lst.Append("c")
		lst.Append("d")

		if lst.GetLen() != 4 {
			t.Error("Expected 4, got ", lst.GetLen())
		}

		lst.Prepend("e")
		lst.Prepend("f")

		if lst.GetLen() != 6 {
			t.Error("Expected 6, got ", lst.GetLen())
		}

		if r, _ := lst.Get(0); r != "f" {
			t.Error("Expected f, got ", r)
		}

		lst = containers.NewLinkedList[string]()
		lst.Append("a")
		lst.Append("b")
		lst.Append("c")

		lst.Delete("b")

		if lst.GetLen() != 2 {
			t.Error("Expected 2, got ", lst.GetLen())
		}

		if r, _ := lst.Get(1); r != "c" {
			t.Error("Expected c, got ", r)
		}

		lst.DeleteHead()

		if lst.GetLen() != 1 {
			t.Error("Expected 1, got ", lst.GetLen())
		}

		if r, _ := lst.Get(0); r != "c" {
			t.Error("Expected c, got ", r)
		}

		lst = containers.NewLinkedList[string]()

		lst.Append("a")
		lst.Append("b")
		lst.Append("c")
		lst.Append("d")

		lst.DeleteTail()

		if lst.GetLen() != 3 {
			t.Error("Expected 3, got ", lst.GetLen())
		}

		if r, _ := lst.Get(2); r != "c" {
			t.Error("Expected c, got ", r)
		}

		lst.DeleteAt(1)

		if lst.GetLen() != 2 {
			t.Error("Expected 2, got ", lst.GetLen())
		}

		if r, _ := lst.Get(1); r != "c" {
			t.Error("Expected c, got ", r)
		}

		lst1 := containers.NewLinkedList[string]()
		lst1.Append("a")
		lst1.Append("b")
		lst1.Append("c")

		lst2 := containers.NewLinkedList[string]()
		lst2.Append("a")
		lst2.Append("b")
		lst2.Append("d")

		equal := lst1.Compare(lst2)
		if equal {
			t.Error("Expected false, got ", equal)
		}
	})
}
