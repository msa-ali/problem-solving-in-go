package linkedlist

import (
	"reflect"
	"testing"
)

func CompareTwoSlices[T any](a []T, b []T) bool {
	return reflect.DeepEqual(a, b)
}

func TestNewNode(t *testing.T) {
	node := NewNode(12)
	if node.val != 12 {
		t.Error("incorrect result: expected node value: 12, got:", node.val)
	}
	if node.next != nil {
		t.Error("incorrect result: expected node next: nil, got:", node.next)
	}
}

func TestNewLinkedListWithValue(t *testing.T) {
	ll := NewLinkedListWithValue(10)

	if ll.head == nil || ll.tail == nil {
		t.Error("incorrect result: expected head/tail : not nil, got: nil")
	}

	if ll.head != ll.tail {
		t.Error("incorrect result: head and tail doesnt point to the same address")
	}

	if ll.head.val != 10 {
		t.Error("incorrect result: expected head value: 10, got:", ll.head.val)
	}
}

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList[int]()

	if ll.head != nil || ll.tail != nil {
		t.Error("incorrect result: expected head/tail : nil, got: not nil")
	}
}

func Test_isNil(t *testing.T) {
	var ll *LinkedList[string]

	if !ll.isNil() {
		t.Error("incorrect result: expected linked list : nil, got: not nil")
	}

	ll = NewLinkedList[string]()

	if !ll.isNil() {
		t.Error("incorrect result: expected linked list : nil, got: not nil")
	}

	ll.InsertAtFront("abc")

	if ll.isNil() {
		t.Error("incorrect result: expected linked list : not nil, got: nil")
	}
}

func TestInsertAtFront(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.InsertAtFront(9)
	ll.InsertAtFront(8)
	ll.InsertAtFront(7)
	outputArr := ll.Array()
	expectedArr := []int{7, 8, 9}
	if !CompareTwoSlices(expectedArr, outputArr) {
		t.Error("incorrect result: expected value:", expectedArr, "got: ", outputArr)
	}
	ll = nil
	err := ll.InsertAtFront(12)

	if err == nil {
		t.Error("incorrect result: expected error: not nil, got: nil ")
	}
}

func TestDeleteAtFront(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.InsertAtFront(9)
	ll.InsertAtFront(8)
	ll.InsertAtFront(7)
	ll.DeleteAtFront()
	ll.DeleteAtFront()
	outputArr := ll.Array()
	expectedArr := []int{9}
	if !CompareTwoSlices(expectedArr, outputArr) {
		t.Error("incorrect result: expected value:", expectedArr, "got: ", outputArr)
	}
	ll.DeleteAtFront()
	err := ll.DeleteAtFront()

	if err == nil {
		t.Error("incorrect result: expected error: not nil, got: nil ")
	}
}

func TestInsertAtRear(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.InsertAtRear(9)
	ll.InsertAtRear(8)
	ll.InsertAtRear(7)
	outputArr := ll.Array()
	expectedArr := []int{9, 8, 7}
	if !CompareTwoSlices(expectedArr, outputArr) {
		t.Error("incorrect result: expected value:", expectedArr, "got: ", outputArr)
	}
	ll = nil
	err := ll.InsertAtFront(12)

	if err == nil {
		t.Error("incorrect result: expected error: not nil, got: nil ")
	}
}

func TestDeleteAtRear(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.InsertAtFront(9)
	ll.InsertAtFront(8)
	ll.InsertAtFront(7)
	ll.DeleteAtRear()
	ll.DeleteAtRear()
	outputArr := ll.Array()
	expectedArr := []int{7}
	if !CompareTwoSlices(expectedArr, outputArr) {
		t.Error("incorrect result: expected value:", expectedArr, "got: ", outputArr)
	}
	ll.DeleteAtFront()
	err := ll.DeleteAtFront()

	if err == nil {
		t.Error("incorrect result: expected error: not nil, got: nil ")
	}
}

func TestInsertAtPosition(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.InsertAtRear(11)
	ll.InsertAtRear(12)
	ll.InsertAtRear(13)
	ll.InsertAtRear(14)
	ll.InsertAtRear(15)
	ll.InsertAtRear(16)
	ll.InsertAtPosition(10, 0)
	ll.InsertAtPosition(17, -1)
	ll.InsertAtPosition(18, 3)
	outputArr := ll.Array()
	expectedArr := []int{10, 11, 12, 18, 13, 14, 15, 16, 17}
	if !CompareTwoSlices(expectedArr, outputArr) {
		t.Error("incorrect result: expected value:", expectedArr, "got: ", outputArr)
	}
	ll = nil
	err := ll.InsertAtPosition(123, 5)

	if err == nil {
		t.Error("incorrect result: expected error: not nil, got: nil ")
	}
}

func TestDeleteAtPosition(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.InsertAtRear(11)
	ll.InsertAtRear(12)
	ll.InsertAtRear(13)
	ll.InsertAtRear(14)
	ll.InsertAtRear(15)
	ll.InsertAtRear(16)
	ll.InsertAtPosition(10, 0)
	ll.InsertAtPosition(17, -1)
	ll.InsertAtPosition(18, 3)
	ll.DeleteAtPosition(0)
	ll.DeleteAtPosition(-2)
	ll.DeleteAtPosition(2)
	outputArr := ll.Array()
	expectedArr := []int{11, 12, 13, 14, 15, 16}
	if !CompareTwoSlices(expectedArr, outputArr) {
		t.Error("incorrect result: expected value:", expectedArr, "got: ", outputArr)
	}
	ll = nil
	err := ll.DeleteAtPosition(4)

	if err == nil {
		t.Error("incorrect result: expected error: not nil, got: nil ")
	}
}
