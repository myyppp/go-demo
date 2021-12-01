package quere

import "sync"

type LinkedQueue struct {
	node *ListNode
	size int
	lock sync.Mutex
}

type ListNode struct {
	Next *ListNode
	Val  int
}

func (l *LinkedQueue) Add(v int) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.size == 0 {
		l.node = &ListNode{Val: v}
		l.size++
		return
	}

	newNode := l.node // 复制头节点
	for newNode.Next != nil {
		newNode = newNode.Next
	}
	newNode.Next = &ListNode{Val: v}
	l.size++
}

func (l *LinkedQueue) Remove() int {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.size == 0 {
		panic("empty")
	}
	v := l.node.Val
	l.node = l.node.Next
	l.size--
	return v
}

func (l *LinkedQueue) Empty() bool {
	return l.size == 0
}
