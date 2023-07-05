package containers

import (
	"errors"
	"fmt"
)

type ll_node[T comparable] struct {
	val  T
	next *ll_node[T]
	prev *ll_node[T]
}

type LinkedList[T comparable] struct {
	head *ll_node[T]
	tail *ll_node[T]
	len  uint
}

// NewLinkedList returns a pointer to a new LinkedList[T] with head and tail set to nil and len set to 0.
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
		tail: nil,
		len:  0,
	}
}

// GetLen returns the length of the linked list.
func (list *LinkedList[T]) GetLen() uint {
	return list.len
}

// Append adds a new node to the end of the linked list.
func (list *LinkedList[T]) Append(val T) {
	newNode := &ll_node[T]{
		val:  val,
		next: nil,
		prev: nil,
	}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		newNode.next = nil
		list.tail.next = newNode
		list.tail = newNode
	}

	list.len += 1
}

// Prepend adds a new node to the beginning of the linked list.
func (list *LinkedList[T]) Prepend(val T) {
	newNode := &ll_node[T]{
		val:  val,
		next: nil,
		prev: nil,
	}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head.prev = newNode
		list.head = newNode
	}

	list.len += 1
}

// InsertAt adds a new node at the specified index.
func (list *LinkedList[T]) InsertAt(val T, index uint) error {
	if index > list.GetLen() {
		return errors.New("index out of range")
	} else if index == list.GetLen() {
		list.Append(val)
		return nil
	} else if index == 0 {
		list.Prepend(val)
		return nil
	}

	newNode := &ll_node[T]{
		val:  val,
		next: nil,
		prev: nil,
	}

	i := uint(0)
	for currentNode := list.head; currentNode != nil; currentNode = currentNode.next {
		if i == index {
			newNode.next = currentNode
			newNode.prev = currentNode.prev
			currentNode.prev.next = newNode
			currentNode.prev = newNode
		}

		i += 1
	}

	list.len += 1

	return nil
}

// DeleteTail removes the head of the linked list.
func (list *LinkedList[T]) DeleteTail() {
	if list.tail != nil {
		if list.tail.prev == nil {
			list.head = nil
			list.tail = nil
		} else {
			list.tail.prev.next = nil
			list.tail.prev = nil
			list.tail.next = nil
			list.tail = list.tail.prev
		}

		list.len -= 1
	}
}

// DeleteAt removes the node at the specified index.
func (list *LinkedList[T]) DeleteAt(index uint) {
	if index > list.GetLen() {
		return
	} else if index == list.GetLen() {
		list.DeleteTail()
		return
	} else if index == 0 {
		list.DeleteHead()
		return
	}

	i := uint(0)
	for currentNode := list.head; currentNode != nil; currentNode = currentNode.next {
		if i == index {
			if currentNode.next == nil {
				list.DeleteTail()
				return
			}

			if currentNode.prev == nil {
				list.DeleteHead()
				return
			}

			currentNode.prev.next = currentNode.next
			currentNode.next.prev = currentNode.prev
			currentNode.next = nil
			currentNode.prev = nil
		}

		i += 1
	}

	list.len -= 1
}

// DeleteHead removes the head of the linked list.
func (list *LinkedList[T]) DeleteHead() {
	if list.head != nil {
		if list.head.next == nil {
			list.head = nil
			list.tail = nil
		} else {
			list.head = list.head.next
			list.head.prev = nil
		}

		list.len -= 1
	}
}

// Delete removes the first node with the specified value.
func (list *LinkedList[T]) Delete(val T) {
	didModifyList := false
	for currentNode := list.head; currentNode != nil; currentNode = currentNode.next {
		if currentNode.val == val {
			if currentNode.next == nil {
				list.DeleteTail()
				return
			}

			if currentNode.prev == nil {
				list.DeleteHead()
				return
			}

			currentNode.prev.next = currentNode.next
			currentNode.next.prev = currentNode.prev
			currentNode.next = nil
			currentNode.prev = nil

			didModifyList = true
		}
	}

	if didModifyList {
		list.len -= 1
	}
}

// Get returns the value at the specified index.
func (list *LinkedList[T]) Get(index uint) (T, error) {
	if index > list.GetLen() {
		var nilVal T
		return nilVal, errors.New("index out of range")
	}

	i := uint(0)
	for currentNode := list.head; currentNode != nil; currentNode = currentNode.next {
		if i == index {
			return currentNode.val, nil
		}

		i += 1
	}

	var nilVal T
	return nilVal, errors.New("value not found")
}

// String returns a string representation of the linked list.
func (list *LinkedList[T]) String() string {
	output := ""

	output += "values: ["
	for currentNode := list.head; currentNode != nil; currentNode = currentNode.next {
		output += fmt.Sprintf("%v", currentNode.val)

		if currentNode.next != nil {
			output += ", "
		}
	}
	output += "]"

	output += " ; len: " + fmt.Sprintf("%v", list.GetLen())

	return output
}

// Compare returns true if the two linked lists are equal.
func (list *LinkedList[T]) Compare(other *LinkedList[T]) bool {
	if list.GetLen() != other.GetLen() {
		return false
	}

	i := uint(0)
	for currentNode := list.head; currentNode != nil; currentNode = currentNode.next {
		otherVal, err := other.Get(i)
		if err != nil {
			return false
		}

		if currentNode.val != otherVal {
			return false
		}

		i++
	}

	return true
}
