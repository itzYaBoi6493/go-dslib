package deque

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	d := New()
	if d.Size() != 0 {
		t.Errorf("New: Size check: fail")
	}
	if !d.Empty() {
		t.Errorf("New: Empty check: fail")
	}
}

func TestPush(_ *testing.T) {
	d := New()

	for i := 0; i < 10; i++ {
		d.PushFront(i)
	}

	fmt.Println("intermediate: ", d)

	for i := 0; i < 10; i++ {
		d.PushBack(i)
	}

	fmt.Println("the deque is: ", d)
}

func TestBack(t *testing.T) {
	d := New()
	vals := []interface{}{1, 2, 3, 4, 5}

	for _, val := range vals {
		d.PushFront(val)
	}
	i := 0
	for !d.Empty() {
		back, _ := d.Back()
		got := (*back).(int)
		if got != vals[i] {
			t.Errorf("Back: got %v, wanted %v", got, vals[i])
		}
		d.PopBack()
		i++
	}
}
