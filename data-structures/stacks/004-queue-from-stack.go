package stacks

// Design a custom queue, MyQueue, using only two stacks. Implement the Push(), Pop(), Peek(), and Empty() methods:

// Void Push(int x): Pushes element at the end of the queue.
// Int Pop(): Removes and returns the element from the front of the queue.
// Int Peek(): Returns the element at the front of the queue.
// Boolean Empty(): Returns TRUE if the queue is empty. Otherwise FALSE.
// You are required to only use the standard stack operations,
// which means that only the Push() to top, Peek() and Pop() from the top, Size(), and Is Empty() operations are valid.

type MyQueue struct {
	s1 Stack[int]
	s2 Stack[int]
}

func NewQueue(size int) MyQueue {
	return MyQueue{
		s1: NewStack[int](size),
		s2: NewStack[int](size),
	}
}

func (q *MyQueue) push(x int) {
	for !q.s1.IsEmpty() {
		el, _ := q.s1.Pop()
		q.s2.Push(el)
	}
	q.s1.Push(x)
	for !q.s2.IsEmpty() {
		el, _ := q.s2.Pop()
		q.s1.Push(el)
	}
}

func (q *MyQueue) pop() int {
	popped, _ := q.s1.Pop()
	return popped
}

func (q *MyQueue) peek() int {
	first, _ := q.s1.Peek()
	return first
}

func (q *MyQueue) empty() bool {
	return q.s1.IsEmpty()
}
