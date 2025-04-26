package stack

type Node struct {
	Next  *Node
	value int
}

type Stack struct {
	head   *Node
	tail   *Node
	length int
}

func (s *Stack) Peek() int {
	if s.isEmpty() {
		return -1
	} else {
		return s.head.value
	}

}

func (s *Stack) push(data int) {
	newNode := &Node{value: data, Next: nil}
	if s.isEmpty() {
		s.head = newNode
		s.tail = newNode
	} else {
		newNode.Next = s.head
		s.head = newNode
	}
	s.length++

}

func (s *Stack) pop() int {
	if s.isEmpty() {
		return -1
	}
	value := s.head.value
	s.head = s.head.Next
	if s.isEmpty() {
		s.tail = nil
	}
	s.length--
	return value
}

func (s *Stack) isEmpty() bool {
	return s.head == nil
}
