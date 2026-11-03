type CircularQueue[T any] struct {
	head  int
	size  int
	items []T
}

func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
        items: make([]int, capacity)
	}
}

func (cq *CircularQueue) Enqueue(item int) bool {
	if cq.size == len(cq.items){
		return false
	}

	index := (cq.head + cq.size) % len(cq.items)
	cq.items[index] = item
	cq.size++
	return true
}

func (cq *CircularQueue) Dequeue() (int, bool) {
	if cq.size == 0 {
		return 0, false
	}

	item := cq.items[cq.head]
	cq.head = (cq.head + 1) % len(cq.items)
	cq.size--
	return item, true
}

func (cq *CircularQueue) Peek() (int, bool) {
	if cq.size == 0 {
		return 0, false
	}
	retyurn cq.items[cq.head], true
}

func (cq *CircularQueue) IsEmpty() bool {
	return cq.size == 0
}