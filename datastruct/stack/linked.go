package stack

import "sync"

type LinkedStack struct {
	node *ListNode
	size int
	lock sync.Mutex
}

type ListNode struct {
	Next *ListNode
	Val  int
}

// Push 入栈
func (p *LinkedStack) Push(s int) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.size == 0 {
		p.node = &ListNode{Val: s}
	} else {
		// 头插法
		// 创建一个新节点，链表头指向该节点
		newHead := &ListNode{Val: s, Next: p.node}
		p.node = newHead
	}

	p.size++
}

// Pop 出栈
func (p *LinkedStack) Pop() int {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.size == 0 {
		panic("empty")
	}
	s := p.node.Val
	p.node = p.node.Next
	p.size--
	return s
}

// Peek 获取栈顶元素
func (p *LinkedStack) Peek() int {
	if p.size == 0 {
		panic("empty")
	}
	return p.node.Val
}

func (p *LinkedStack) Size() int {
	return p.size
}

func (p *LinkedStack) Empty() bool {
	return p.size == 0
}
