// Package that implements a data structure that supports 
// FIFO (First In First Out) operations in constant
// time complexity.
package queue

import (
    "fmt"
)

// Unexported struct type. A node for the doubly linked list
// data type.
type node struct {
    data interface{}
    next *node
    prev *node
}

// Struct for Queue type. Implemented as a Doubly Linked List.
// Additions at the back. Removals at the front.
type Queue struct {
    front *node
    back *node
    size int
}

// New returns an empty Queue.
func New() Queue {
    // default initialization to zero values
    return Queue {}
}

// Size returns the number of elements in the queue currently.
func (q *Queue) Size() (result int) {
    if q != nil {
        result = q.size
    }
    return
}

// Empty returns true if the queue has no elements and false otherwise.
func (q *Queue) Empty() bool {
    return q.Size() == 0
}

// Enqueue adds any number of elements to the back of the queue.
// The first argument specified is inserted first.
func (q *Queue) Enqueue(vals ...interface{}) bool {
    if q == nil {
        return false
    }
    for _, val := range vals {
        // for each value
        // create a node out of the value
        newNode := node {
            data: val,
            next: q.back,
        }

        // add newNode to the back of the queue.
        // first, connect q.back's prev field to newNode
        if q.back != nil {
            q.back.prev = &newNode
        } else {
            // first element in the queue
            q.front = &newNode
        }
        // new back
        q.back = &newNode
        q.size++
    }
    return true
}

// String returns a string representation of the queue from back to front from left to right.
// Can be used with fmt.Print and the likes.
func (q Queue) String() string {
    var result string
    sep := " "
    for curr := q.back; curr != nil; curr = curr.next {
        if curr.next == nil {
            sep = ""
        }
        result += fmt.Sprintf("%v%v", curr.data, sep)
    }
    return result
}

// Clear resets the queue, i.e, effectively removes every element from the
// queue without returning.
func (q *Queue) Clear() {
    if q != nil {
        q.size = 0
        q.front = nil
        q.back = nil
    }
}

// Dequeue removes the element from the front, if any. Returns the data and a bool(ok).
// If the front element can be removed successfully, then ok = true. False otherwise
func (q *Queue) Dequeue() (interface{}, bool) {
    if q == nil || q.Empty() {
        return nil, false
    }
    // preserve the current element at the front
    oldFront := q.front
    // new front
    q.front = oldFront.prev
    // nothing left??
    if q.front == nil {
        // then empty queue
        q.back = nil
    } else {
        // else, detach all connections
        q.front.next = nil
    }
    q.size--
    // retrieve data
    data := oldFront.data
    // clean up
    oldFront.next = nil
    oldFront.prev = nil
    return data, true
}

// Front returns a pointer to the data at the front, if any. Returns a pointer
// to the data and a bool(ok). If data was successfully returned, then ok = true.
// Otherwise, false.
func (q *Queue) Front() (*interface{}, bool) {
    if q == nil || q.Empty() {
        return nil, false
    }
    return &q.front.data, true
}

// Back returns a pointer to the data at the back of the queue, if any.
// Returns a pointer to the data and a bool(ok). If data was successfully returned,
// then ok = true. Otherwise, false.
func (q *Queue) Back() (*interface{}, bool) {
    if q == nil || q.Empty() {
        return nil, false
    }
    return &q.back.data, true
}
