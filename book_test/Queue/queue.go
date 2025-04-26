package queue

import "fmt"

type Node struct {
	next  *Node
	value int
}
type Queue struct {
	head   *Node
	tail   *Node
	length int
}

func (q *Queue) Enqueue(digit int) {
	newNode := &Node{value: digit, next: nil}
	if q.tail == nil {
		q.tail = newNode
		q.head = newNode

	} else {
		q.tail.next = newNode
		q.tail = newNode

	}
	q.length++
}

func (q *Queue) Peek() int {
	if q.isQueueEmpty() {
		return -1
	}
	return q.head.value
}

func (q *Queue) Dqueue() int {
	if q.isQueueEmpty() {
		fmt.Println("the queue is empty")
		return -1
	}
	value := q.head.value
	q.head = q.head.next
	if q.isQueueEmpty() {
		q.tail = nil
	}
	q.length--
	return value

}

func (q *Queue) isQueueEmpty() bool {
	return q.head == nil

}
