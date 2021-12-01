package stack

import "sync"

type ArrayStack struct {
	array []int
	size  int
	lock  sync.Mutex
}

func (p *ArrayStack) Push(s int) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.array = append(p.array, s)
	p.size++
}

func (p *ArrayStack) Pop() int {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.size == 0 {
		panic("empty")
	}

	s := p.array[p.size-1]
	p.array = p.array[:p.size-1]
	p.size--
	return s
}

func (p *ArrayStack) Peek() int {
	return p.array[p.size-1]
}
