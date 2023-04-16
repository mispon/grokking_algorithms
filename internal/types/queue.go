package types

import "sync"

type Queue[T any] struct {
	lock  sync.Mutex
	items []T
}

func NewQueue[T any](cap int) *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0, cap),
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.Len() == 0 {
		var empty T
		return empty, false
	}

	q.lock.Lock()
	defer q.lock.Unlock()

	item := q.items[0]
	q.items = q.items[1:]

	return item, true
}

func (q *Queue[T]) Len() int {
	return len(q.items)
}
