// Package which implements a data structure that allows constant time additions
// and removals from both the front and back of the data structure.
package deque

import (
	"fmt"
)

// Doubly linked list node
type node struct {
	data interface{}
	next *node
	prev *node
}

// Deque struct type.
// front of the struct is the tail of the doubly linked list and
// back of the struct is the head of the doubly linked list.
type Deque struct {
	size int
	front, back *node
}

// New creates and returns an empty Deque.
func New() Deque {
	return Deque {}
}

// Size returns the number of elements in the Deque currently.
func (d *Deque) Size() int {
	if d == nil {
		return 0
	}
	return d.size
}

// Empty returns whether the deque is empty or not. Returns true if empty
// and false otherwise.
func (d *Deque) Empty() bool {
	return d.Size() == 0
}

// PushFront adds an element to the front of the deque. Returns true if push is
// successful.
func (d *Deque) PushFront(data interface{}) bool {
	if d == nil {
		return false
	}
	// create a new node out of the data
	newNode := node {
		data: data,
		prev: d.front,
	}

	if d.front != nil {
		d.front.next = &newNode
	} else {
		d.back = &newNode
	}
	d.front = &newNode
	d.size++
	return true
}

// PushBack adds an element to the back of the queue. Returns true if push is successful.
func (d *Deque) PushBack(data interface{}) bool {
	if d == nil {
		return false
	}
	// create a new node
	newNode := node {
		data: data,
		next: d.back,
	}

	if d.back != nil {
		d.back.prev = &newNode
	} else {
		d.front = &newNode
	}
	d.back = &newNode
	d.size++
	return true
}

// String returns a string representation of the deque. The string returned is
// from back to front.
func (d Deque) String() string {
	result := ""
	sep := " "
	for curr := d.back; curr != nil; curr = curr.next {
		if curr.next == nil {
			sep = ""
		}
		result += fmt.Sprintf("%v%v", curr.data, sep)
	}
	return result
}

// PopFront returns a data and a bool(ok). If data was removed successfully from
// the front of the deque, then ok is true. Else, false.
func (d *Deque) PopFront() (interface{}, bool) {
	if d.Empty() {
		return nil, false
	}
	oldFront := d.front
	d.front = oldFront.prev

	if d.front == nil {
		d.back = nil
	} else {
		d.front.next = nil
	}
	// collect the data
	ret := oldFront.data
	// clean up
	oldFront.next = nil
	oldFront.prev = nil
	// size decrement
	d.size--
	return ret, true
}

// PopBack removes and returns the element at the back of the deque.
// Returns the data and a bool(ok). If data was successfully removed, then ok = true.
// Otherwise, false.
func (d *Deque) PopBack() (interface{}, bool) {
	if d.Empty() {
		return nil, false
	}
	oldBack := d.back
	d.back = oldBack.next
	if d.back == nil {
		d.front = nil
	} else {
		d.back.prev = nil
	}
	ret := oldBack.data
	// clean up
	oldBack.next = nil
	oldBack.prev = nil
	d.size--
	return ret, true
}

// Front returns a pointer to the element at the front if it exists. Returns a
// pointer and a bool, which is true if the front element can be successfully extracted
// and false otherwise.
func (d *Deque) Front() (*interface{}, bool) {
	if d.Empty() {
		return nil, false
	}
	return &d.front.data, true
}

// Back returns a pointer to the element at the back without removing the element, if it exists. 
// Returns a pointer to the data and a bool, which is true if the back element 
// can be successfully extracted and false otherwise.
func (d *Deque) Back() (*interface{}, bool) {
	if d.Empty() {
		return nil, false
	}
	return &d.back.data, true
}
