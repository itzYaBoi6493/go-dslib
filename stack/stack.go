// Package to implement a LIFO (Last In First Out) data structure
// that supports Push and Pop operations in constant time complexity.
package stack

import "fmt"

// Private Node for the singly linked list
type node struct {
    data interface{}
    next *node
}

// Struct for the Stack data type
type Stack struct {
    head *node
    size int
}

// New creates and returns a new empty stack.
func New() Stack {
    return Stack {
        head: nil,
        size: 0,
    }
}

// Size returns the number of elements in the stack currently.
func (s *Stack) Size() int {
    return s.size
}

// Push takes any number of arguments (variadic function) and adds each one
// of them to the stack top. The last argument supplied is at the top after
// the Push operation. Returns true on success and false on failure
func (s *Stack) Push(vals ...interface{}) bool {
    if s == nil {
        return false
    }

    for _, val := range vals {
        // for each val, create a new node
        // out of the val
        newNode := node {
            data: val,
            next: s.head,
        }
        
        s.head = &newNode
        s.size++
    }
    return true
}

// String allows to print the stack. Can be used with fmt.Print or the likes.
func (s Stack) String() string {
    res := ""
    sep := " "
    for curr := s.head; curr != nil; curr = curr.next {
        if curr.next == nil {
            sep = ""
        }
        res += fmt.Sprintf("%v%v", curr.data, sep);
    }
    return res
}

// Empty returns if a stack is empty. Returns true if stack has 0 elements
// and false otherwise
func (s *Stack) Empty() bool {
    if s == nil || s.size == 0 {
        return true
    }
    return false
}

// Pop returns the top of the Stack and removes it from the stack, if it exists.
// Returns the data and ok (bool). If the top exists, ok = true. False otherwise.
func (s *Stack) Pop() (interface{}, bool) {
    if s == nil || s.Empty() {
        return nil, false
    }

    oldHead := s.head
    s.head = oldHead.next
    s.size--
    return oldHead.data, true
}

// Clear resets the stack.
func (s *Stack) Clear() {
    if s != nil {
        // let the GC do all the work
        s.head = nil
        s.size = 0
    }
}

// Peek returns the top of the stack without popping off, if it exists.
// Returns a pointer to the data and a bool(ok), which denotes successful query of the top.
// If the top exists, ok = true. False otherwise
func (s *Stack) Peek() (*interface{}, bool) {
    if s == nil || s.Empty() {
        return nil, false
    }

    return &s.head.data, true
}
