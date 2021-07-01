package stack

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	s := New()
	if s.size != 0 {
		t.Errorf("New: size test: got %v, wanted 0", s.size)
	}

	if s.head != nil {
		t.Errorf("New: head test: got non-empty head")
	}
}

func TestPush(t *testing.T) {
	s := New()
	vals := []interface{}{1, 2, 3, 4, 5}
	s.Push(vals...)
	fmt.Println(s.String())

	want := 5
	got := s.Size()
	if got != want {
		t.Errorf("Push: Size: got %v, wanted %v", got, want)
	}
}

func TestPop(t *testing.T) {
	s := New()
	n := 20
	rand.Seed(time.Time{}.Unix())
	vals := make([]interface{}, 0, n)
	for i := 0; i < n; i++ {
		vals = append(vals, rand.Int())
	}
	s.Push(vals...)
	fmt.Println("TestPop: Size: ", s.Size())

	for i := n - 1; i >= 0; i-- {
		data, ok := s.Pop()
		if !ok {
			t.Errorf("Pop: Spurious false value from non-empty stack")
		}
		got := data.(int)
		want := vals[i].(int)
		if got != want {
			t.Errorf("Pop: got %v, wanted %v", got, want)
		}
	}

	if !s.Empty() {
		t.Errorf("Pop: non-empty stack after all pops")
	}
}

func TestPeek(t *testing.T) {
	s := New()
	n := 15
	rand.Seed(time.Time{}.Unix())

	want := rand.Int()
	for i := 0; i < n-1; i++ {
		s.Push(rand.Int())
	}
	s.Push(want)

	if data, ok := s.Peek(); ok {
		if *data != want {
			t.Errorf("Peek: incorrect top: got %v, wanted %v", *data, want)
		}
	} else {
		t.Errorf("Peek: invalid return status")
	}
}
