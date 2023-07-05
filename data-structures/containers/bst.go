package containers

import "fmt"

type _comparable interface {
	// the built-in `comparable` interface doesn't define <, >, <=, and >= operators.
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

type Node[T _comparable] struct {
	Val   T
	left  *Node[T]
	right *Node[T]
}

type BST[T _comparable] struct {
	root *Node[T]
	len  uint
}

// NewBST returns a pointer to a new BST[T] with root set to nil and len set to 0.
func NewBST[T _comparable]() *BST[T] {
	return &BST[T]{
		root: nil,
		len:  0,
	}
}

// GetLen returns the length of the binary tree.
func (tree *BST[T]) GetLen() uint {
	return tree.len
}

// Lookup returns the value of the node with the given value and a boolean indicating whether the node was found.
func (tree *BST[T]) Lookup(val T) (*Node[T], bool) {
	if tree.root == nil {
		return nil, false
	}

	returnVal, code := tree.lookup(tree.root, val)

	return returnVal, code
}

func (tree *BST[T]) lookup(currentNode *Node[T], val T) (*Node[T], bool) {
	if val == currentNode.Val {
		return currentNode, true
	}

	if val < currentNode.Val {
		if currentNode.left == nil {
			return nil, false
		} else {
			return tree.lookup(currentNode.left, val)
		}
	} else {
		if currentNode.right == nil {
			return nil, false
		} else {
			return tree.lookup(currentNode.right, val)
		}
	}
}

// Insert adds a new node to the binary tree.
func (tree *BST[T]) Insert(val T) {
	newNode := &Node[T]{
		Val:   val,
		left:  nil,
		right: nil,
	}

	if tree.root == nil {
		tree.root = newNode
		tree.len += 1
	} else {
		tree.insert(tree.root, newNode)
	}
}

// Delete removes a node from the binary tree.
func (tree *BST[T]) Delete(val T) {
	tree.delete(nil, tree.root, val)
}

// Traverse prints the values of the binary tree in order.
func (tree *BST[T]) Traverse() {
	if tree.root == nil {
		return
	}

	tree.traverse(tree.root)
}

// Compare compares the values of two binary trees.
func (tree *BST[T]) Compare(other *BST[T]) bool {
	return tree.compareNodes(tree.root, other.root)
}

func (tree *BST[T]) compareNodes(currentNode *Node[T], otherNode *Node[T]) bool {
	if currentNode == nil && otherNode == nil {
		return true
	}

	if currentNode == nil || otherNode == nil {
		return false
	}

	if currentNode.Val != otherNode.Val {
		return false
	}

	return tree.compareNodes(currentNode.left, otherNode.left) && tree.compareNodes(currentNode.right, otherNode.right)
}

func (tree *BST[T]) insert(currentRoot *Node[T], newNode *Node[T]) {
	if currentRoot == nil {
		return
	}

	if newNode.Val == currentRoot.Val {
		return
	}

	if newNode.Val < currentRoot.Val {
		if currentRoot.left == nil {
			currentRoot.left = newNode
			tree.len += 1
		} else {
			tree.insert(currentRoot.left, newNode)
		}
	} else {
		if currentRoot.right == nil {
			currentRoot.right = newNode
			tree.len += 1
		} else {
			tree.insert(currentRoot.right, newNode)
		}
	}
}

// TODO: INCOMPLETE
func (tree *BST[T]) delete(parentNode *Node[T], currentRoot *Node[T], val T) {
	if val == currentRoot.Val {
		if currentRoot.right == nil {
			if parentNode != nil {

			}
		}
	}

	if val < currentRoot.Val {
		if currentRoot.left == nil {
			return
		}

		tree.delete(currentRoot, currentRoot.left, val)
	} else {
		if currentRoot.right == nil {
			return
		}

		tree.delete(currentRoot, currentRoot.right, val)
	}
}

// traverse prints the values of the binary tree in order.
func (tree *BST[T]) traverse(currentRoot *Node[T]) {
	if currentRoot == nil {
		return
	}

	tree.traverse(currentRoot.left)
	fmt.Println(currentRoot.Val)
	tree.traverse(currentRoot.right)
}
