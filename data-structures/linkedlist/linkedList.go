package linkedlist

import (
	"errors"
	"fmt"
)

type LinkedListNode[T any] struct {
	val  T
	next *LinkedListNode[T]
}

type LinkedList[T any] struct {
	head *LinkedListNode[T]
	tail *LinkedListNode[T]
}

// Create a New Linked List Node
func NewNode[T any](val T) *LinkedListNode[T] {
	return &LinkedListNode[T]{
		val:  val,
		next: nil,
	}
}

// Function to initialize linked list with a value
func NewLinkedListWithValue[T any](val T) *LinkedList[T] {
	newNode := NewNode(val)
	return &LinkedList[T]{
		head: newNode,
		tail: newNode,
	}
}

// Function to initialize empty linked list
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
		tail: nil,
	}
}

// Checks if linked list is not initialized or is empty
func (ll *LinkedList[T]) isNil() bool {
	if ll == nil || ll.head == nil || ll.tail == nil {
		return true
	}
	return false
}

// Function to insert value at front position in linked list
func (ll *LinkedList[T]) InsertAtFront(val T) error {
	// if linked list is not initialized
	if ll == nil {
		return errors.New("can't invode insert operation on nil linked list")
	}
	newNode := NewNode(val)
	// if linked list head is nil means linked list is empty
	if ll.head == nil {
		newNode.next = nil
		ll.tail = newNode
	} else {
		newNode.next = ll.head
	}
	ll.head = newNode
	return nil
}

// Function to delete value at front position in linked list.
// It returns error if linked list is not initialized or is empty
func (ll *LinkedList[T]) DeleteAtFront() error {
	if ll.isNil() {
		return errors.New("can't invode delete operation on nil linked list")
	}
	nextNode := ll.head.next
	// case 1: linked list has only element
	if nextNode == nil {
		ll.head = nil
		ll.tail = nil
	} else {
		// case 2: linked list has more than one element
		ll.head.next = nil
		ll.head = nextNode
	}
	return nil
}

// Function to insert value at rear position in linked list
func (ll *LinkedList[T]) InsertAtRear(val T) error {
	// if linked list is not initialized
	if ll == nil {
		return errors.New("can't invode insert operation on nil linked list")
	}
	newNode := NewNode(val)
	// if linked list tail is nil means linked list is empty
	if ll.tail == nil {
		ll.head = newNode
	} else {
		ll.tail.next = newNode
	}
	ll.tail = newNode
	return nil
}

// Function to delete value at rear position in linked list.
// It returns error if linked list is not initialized or is empty
func (ll *LinkedList[T]) DeleteAtRear() error {
	if ll.isNil() {
		return errors.New("can't invode delete operation on nil linked list")
	}
	// case 1: linked list has only element
	if ll.head == ll.tail {
		ll.head = nil
		ll.tail = nil
	} else {
		// case 2: linked list has more than one element
		temp := ll.head
		for {
			// iteration reaches one node before tail node
			if temp.next == ll.tail {
				temp.next = nil
				ll.tail = temp
				break
			}
			temp = temp.next
		}
	}
	return nil
}

// Function to insert node in a linked list at a given position.
// If position is set as 0, node will be inserted at front position.
// If position is set as any negative value or more than actual length of linked list, node will be inserted at rear position
func (ll *LinkedList[T]) InsertAtPosition(val T, pos int) error {
	if ll.isNil() {
		return errors.New("can't invoke insert operation on nil linked list")
	}
	// if position is 0, insert at front
	if pos == 0 {
		return ll.InsertAtFront(val)
	}
	//  if position is -ve, insert at rear
	if pos < 0 {
		return ll.InsertAtRear(val)
	}
	currPos := 0
	temp := ll.head
	for {
		if temp == nil || currPos == pos-1 {
			break
		}
		currPos += 1
		temp = temp.next
	}
	// if position is gt length of linked list, insert at rear
	if temp == nil {
		return ll.InsertAtRear(val)
	}
	newNode := NewNode(val)
	nextNode := temp.next
	temp.next = newNode
	newNode.next = nextNode
	// if newNode next is nil, set new node as linked list tail
	if newNode.next == nil {
		ll.tail = newNode
	}
	return nil
}

// Function to DELETE node in a linked list at a given position.
// If position is set as 0, node will be deleted at front position.
// If position is set as any negative value or more than actual length of linked list, node will be deleted at rear position
func (ll *LinkedList[T]) DeleteAtPosition(pos int) error {
	if ll.isNil() {
		return errors.New("can't invode delete operation on nil linked list")
	}
	// if position is 0, delete at front
	if pos == 0 {
		return ll.DeleteAtFront()
	}
	//  if position is -ve, insert at rear
	if pos < 0 {
		return ll.DeleteAtRear()
	}
	currPos := 0
	temp := ll.head
	for {
		if temp == nil || currPos == pos-1 {
			break
		}
		currPos += 1
		temp = temp.next
	}
	// if position is gt length of linked list, delete at rear
	if temp == nil {
		return ll.DeleteAtRear()
	}
	nextNode := temp.next
	if nextNode == nil {
		return ll.DeleteAtRear()
	}
	nextNode = nextNode.next
	temp.next = nextNode
	if nextNode.next == nil {
		ll.tail = nextNode
	}
	return nil
}

// Function to print linked list
func (ll *LinkedList[T]) Print() {
	if ll.isNil() {
		fmt.Println("Linked List is empty or nil. Try inserting value or reinitiative a new linked list")
		return
	}
	// if linked list is not empty, iterate it reaches nil
	temp := ll.head
	for {
		if temp == nil {
			fmt.Printf(" nil\n\n")
			return
		}
		fmt.Printf(" %v ->", temp.val)
		temp = temp.next
	}
}

// Convert linked list value to slice
func (ll *LinkedList[T]) Array() (result []T) {
	if ll.isNil() {
		return
	}
	result = make([]T, 0)
	temp := ll.head
	for {
		if temp == nil {
			return
		}
		result = append(result, temp.val)
		temp = temp.next
	}
}
