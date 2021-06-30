package stack

import "testing"
import "fmt"

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
    vals := []interface{}{1, 2, 3, 4, 5}
    s.Push(vals...)

    for i := s.Size() - 1; i >= 0; i-- {
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
