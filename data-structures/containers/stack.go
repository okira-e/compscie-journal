package containers

import "fmt"

type stackNode[T comparable] struct {
	val  T
	prev *stackNode[T]
}

type Stack[T comparable] struct {
	top *stackNode[T]
	len uint
}

// NewStack returns a pointer to a new Stack[T] with top set to nil and len set to 0.
func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{
		top: nil,
		len: 0,
	}
}

// GetLen returns the length of the stack.
func (stack *Stack[T]) GetLen() uint {
	return stack.len
}

// Peek returns the value of the top of the stack and an error code.
func (stack *Stack[T]) Peek() (T, int8) {
	if stack.top == nil {
		var nilVal T
		return nilVal, -1
	}

	return stack.top.val, 0
}

// Push adds a new node to the top of the stack.
func (stack *Stack[T]) Push(val T) {
	newNode := &stackNode[T]{
		val:  val,
		prev: nil,
	}

	stack.len += 1

	if stack.top == nil {
		stack.top = newNode

		return
	}

	newNode.prev = stack.top
	stack.top = newNode
}

// Pop removes the top of the stack and returns its value and an error code.
func (stack *Stack[T]) Pop() (T, int8) {
	if stack.top == nil {
		var nilVal T
		return nilVal, -1
	}

	val := stack.top.val
	stack.top = stack.top.prev

	stack.len -= 1

	return val, 0
}

// String returns a string representation of the stack.
func (stack *Stack[T]) String() string {
	output := ""

	output += "top: ["
	for currentNode := stack.top; currentNode != nil; currentNode = currentNode.prev {
		output += fmt.Sprintf("%v", currentNode.val)

		if currentNode.prev != nil {
			output += ", "
		}
	}
	output += "]"

	output += " ; len: " + fmt.Sprintf("%v", stack.GetLen())

	return output
}

// Compare returns true if the two stacks are equal.
func (stack *Stack[T]) Compare(other Stack[T]) bool {
	if stack.GetLen() != other.GetLen() {
		return false
	}

	selfNode := stack.top
	otherNode := other.top
	for i := uint(0); i < stack.GetLen(); i += 1 {
		if selfNode.val != otherNode.val {
			return false
		}

		selfNode = selfNode.prev
		otherNode = otherNode.prev
	}

	return true
}
