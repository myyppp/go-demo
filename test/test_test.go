package test_test

import (
	"testing"
)

func TestLeetcode(t *testing.T) {
	m := Constructor()
	m.Push(-2)
	m.Push(0)
	m.Push(-3)
	m.GetMin()
	m.Pop()
	m.Pop()
	m.GetMin()
}

type MinStack struct {
	val  int
	next *MinStack
	min  int
}

func Constructor() *MinStack {
	return &MinStack{min: 1<<31 - 1}
}

func (s *MinStack) Push(val int) {
	min := val
	if s.min < val {
		min = s.min
	}
	temp := &MinStack{
		val:  val,
		next: s,
		min:  min,
	}
	s.val = temp.val
	s.next = temp
	s.min = temp.min

}

func (s *MinStack) Pop() {
	s.val = s.next.val
	s.min = s.next.min
	s.next = s.next.next
}

func (s *MinStack) Top() int {
	return s.val
}

func (s *MinStack) GetMin() int {
	return s.min
}
