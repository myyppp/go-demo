package stack_test

import (
	"fmt"
	"testing"

	. "github.com/myyppp/go-demo/datastruct/stack"
)

func TestLinkedStack(t *testing.T) {
	l := new(LinkedStack)
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Pop()
	fmt.Printf("l.Peek(): %v\n", l.Peek())
	l.Push(4)
	l.Push(5)
	fmt.Printf("l: %v\n", l)
}

func TestArrayStack(t *testing.T) {
	l := new(ArrayStack)
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Pop()
	fmt.Printf("l.Peek(): %v\n", l.Peek())
	l.Push(4)
	l.Push(5)
	fmt.Printf("l: %v\n", l)
}
