package go_cpu

import "testing"

func TestQueue_Enqueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1, 2, 3)
	if len(q.data) != 3 {
		t.Errorf("invalid: %v", q)
	}
	q.Enqueue(4, 5, 6)
	if len(q.data) != 6 {
		t.Errorf("invalid: %v", q)
	}
	t.Logf("%v", q)
}

func TestQueue_Dequeue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1,2,3)
	data := q.Dequeue(2)
	if len(data) != 2 {
		t.Errorf("invalid: %v", q)
	}
	t.Logf("%v", q)
}
