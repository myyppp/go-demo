package quere

import (
	"sync"
)

type ArrayQueue struct {
	array []int
	size  int
	lock  sync.Mutex
}

func (q *ArrayQueue) Add(v int) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.array = append(q.array, v)
	q.size++
}

func (q *ArrayQueue) Remove() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.size == 0 {
		panic("empty")
	}

	v := q.array[0]
	q.array = q.array[1:]
	q.size--
	return v
}

func (l *ArrayQueue) Empty() bool {
	return l.size == 0
}
