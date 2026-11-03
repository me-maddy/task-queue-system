package internal

type Queue struct {
	items []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}

	var value int = q.items[0]
	q.items = q.items[1:]
	return value, true
}

func (q *Queue) Peek() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}

	return q.items[0], true
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
