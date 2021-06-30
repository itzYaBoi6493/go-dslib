package queue

import (
    "testing"
    "fmt"
)

func TestNew(t *testing.T) {
    q := New()
    if q.front != nil || q.back != nil || q.size != 0 {
        t.Errorf("New: wrong initial values")
    }
}

func TestSize(t *testing.T) {
    got := (*Queue)(nil).Size()
    if got != 0 {
        t.Errorf("Size: failure on nil queue")
    }
}

func TestEnqueue(t *testing.T) {
    q := New()
    n := 10
    for i := 0; i < n; i++ {
        q.Enqueue(i)
    }

    fmt.Println(q)
    got := q.Size()
    if got != n {
        t.Errorf("Enqueue: Size check: got %v, wanted %v", got, n)
    }
}

func TestDequeue(t *testing.T) {
    q := New()
    n := 10
    for i := 0; i < n; i++ {
        q.Enqueue(i)
    }

    fmt.Println(q)
    for i, lim := 0, n / 2; i < lim; i++ {
        data, _ := q.Dequeue()
        fmt.Println("Removed ", data)
    }

    fmt.Println(q)

    left := n / 2
    got := q.Size()
    if got != left {
        t.Errorf("Dequeue: Size check: got %v, wanted %v", got, left)
    }

    for !q.Empty() {
        q.Dequeue()
    }

    fmt.Println("The queue is: ", q)
    if q.Size() != 0 {
        t.Errorf("Dequeue: Empty check: failed")
    }
}

func TestFront(t *testing.T) {
    q := New()
    vals := []interface{}{1, 2, 3, 4, 5, 6, 42, 83}
    q.Enqueue(vals...)
    i := 0
    for !q.Empty() {
        want := vals[i].(int)

        frontGot, _ := q.Front()
        got := (*frontGot).(int)
        if got != want {
            t.Errorf("Front: got %v, wanted %v", got, want)
        }
        i++
        q.Dequeue()
    }
    _, ok := q.Front()
    if ok {
        t.Errorf("Front: Empty Check: false")
    }
}

func TestBack(t *testing.T) {
    n := 10
    q := New()
    _, ok := q.Back()
    if ok {
        t.Errorf("Back: Empty check: false")
    }
    for i := 1; i <= n; i++ {
        q.Enqueue(i)
        data, _ := q.Back()
        got := (*data).(int)
        if got != i {
            t.Errorf("Back: got %v, wanted %v", got, i)
        }
    }
}
