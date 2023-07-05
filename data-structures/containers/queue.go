package containers

import "fmt"

type q_node[T comparable] struct {
	val  T
	next *q_node[T]
}

type Queue[T comparable] struct {
	head *q_node[T]
	tail *q_node[T]
	len  uint
}

// NewQueue returns a pointer to a new Queue[T] with head and tail set to nil and len set to 0.
func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{
		head: nil,
		tail: nil,
		len:  0,
	}
}

// GetLen returns the length of the queue.
func (q *Queue[T]) GetLen() uint {
	return q.len
}

// Peek returns the value of the head of the queue and an error code.
func (q *Queue[T]) Peek() (T, int8) {
	if q.head == nil {
		var nilVal T
		return nilVal, -1
	}

	return q.head.val, 0
}

// Deque removes the head of the queue and returns its value and an error code.
func (q *Queue[T]) Deque() (T, int8) {
	if q.head == nil {
		var nilVal T
		return nilVal, -1
	}

	val := q.head.val
	q.head = q.head.next

	q.len -= 1
	return val, 0
}

// Enqueue adds a new node to the end of the queue.
func (q *Queue[T]) Enqueue(val T) {
	newNode := &q_node[T]{
		val:  val,
		next: nil,
	}

	q.len += 1

	if q.head == nil {
		q.head = newNode
		q.tail = newNode

		return
	}

	q.tail.next = newNode
	q.tail = newNode
}

// String returns a string representation of the queue.
func (q *Queue[T]) String() string {
	output := ""

	output += "values: ["
	for currentNode := q.head; currentNode != nil; currentNode = currentNode.next {
		output += fmt.Sprintf("%v", currentNode.val)

		if currentNode.next != nil {
			output += ", "
		}
	}
	output += "]"

	output += " ; len: " + fmt.Sprintf("%v", q.GetLen())

	return output
}

// Compare returns true if the all the queue values are equal to the other queue values.
func (q *Queue[T]) Compare(other Queue[T]) bool {
	if q.GetLen() != other.GetLen() {
		return false
	}

	selfNode := q.head
	otherNode := other.head
	for i := uint(0); i < q.GetLen(); i += 1 {
		if selfNode.val != otherNode.val {
			return false
		}

		selfNode = selfNode.next
		otherNode = otherNode.next
	}

	return true
}
