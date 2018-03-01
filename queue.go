package go_cpu

type Queue struct {
	data []uint8
}

func NewQueue() *Queue {
	return &Queue{data: []uint8{}}
}

func (q *Queue) Enqueue(data ...uint8) {
	q.data = append(q.data, data...)
}

func (q *Queue) Dequeue(n int) (data []uint8) {
	data, q.data = q.data[:n], q.data[n:]
	return data
}
