package quere_test

import (
	"fmt"
	"testing"

	. "github.com/myyppp/go-demo/datastruct/quere"
)

func TestArrayQueue(t *testing.T) {
	q := new(ArrayQueue)
	q.Add(1)
	v := q.Remove()
	fmt.Printf("v: %v\n", v)
	q.Add(2)
	q.Add(3)
	q.Add(4)
	v = q.Remove()
	fmt.Printf("v: %v\n", v)
	q.Add(5)
	fmt.Printf("q: %v\n", q)
	q.Add(6)
	v = q.Remove()
	fmt.Printf("v: %v\n", v)
}

func TestLinkedQueue(t *testing.T) {
	q := new(LinkedQueue)
	q.Add(1)
	v := q.Remove()
	fmt.Printf("v: %v\n", v)
	q.Add(2)
	q.Add(3)
	q.Add(4)
	v = q.Remove()
	fmt.Printf("v: %v\n", v)
	q.Add(5)
	q.Add(6)
	v = q.Remove()
	fmt.Printf("v: %v\n", v)
}
