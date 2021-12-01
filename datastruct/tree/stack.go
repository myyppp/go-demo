package tree

import "sync"

type Stack struct {
	array []*TreeNode
	size  int
	lock  sync.Mutex
}

// Push 入栈
func (p *Stack) Push(s *TreeNode) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.array = append(p.array, s)
	p.size++
}

// Pop 出栈
func (p *Stack) Pop() *TreeNode {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.size == 0 {
		panic("栈为空")
	}

	s := p.array[p.size-1]
	p.array = p.array[:p.size-1]
	p.size--

	return s
}

// Peek 栈顶元素
func (p *Stack) Peek() *TreeNode {
	if p.size == 0 {
		panic("empty")
	}
	return p.array[p.size-1]
}

// Size 栈大小
func (p *Stack) Size() int {
	return p.size
}

// Empty 栈是否为空
func (p *Stack) Empty() bool {
	return p.size == 0
}
