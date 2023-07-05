package containers_test

import (
	"data_structures/containers"
	"testing"
)

func TestBST(t *testing.T) {
	t.Run("Test BinaryTree", func(t *testing.T) {
		bt := containers.NewBST[int]()

		bt.Insert(5)
		bt.Insert(3)
		bt.Insert(7)
		bt.Insert(2)
		bt.Insert(4)
		bt.Insert(6)
		bt.Insert(8)

		if bt.GetLen() != 7 {
			t.Error("Expected 7, got ", bt.GetLen())
		}

		node, found := bt.Lookup(3)
		if !found {
			t.Error("Expected true, got ", found)
		} else if node.Val != 3 {
			t.Error("Expected 3, got ", node.Val)
		}

		//bt.Delete(5)
		//
		//if bt.GetLen() != 6 {
		//	t.Error("Expected 6, got ", bt.GetLen())
		//}
		//
		//node, found = bt.Lookup(5)
		//if found {
		//	t.Error("Expected false, got ", found)
		//}

		otherBt := containers.NewBST[int]()
		otherBt.Insert(5)
		otherBt.Insert(3)
		otherBt.Insert(7)
		otherBt.Insert(2)
		otherBt.Insert(4)
		otherBt.Insert(6)
		otherBt.Insert(8)

		if !bt.Compare(otherBt) {
			t.Error("Expected true, got false")
		}

		otherBt.Insert(9)

		if bt.Compare(otherBt) {
			t.Error("Expected false, got true")
		}
	})
}
