package linked

// 循环链表
type Ring struct {
	next, prev *Ring // 前驱和后继节点
	Value      interface{}
}

func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

// 指定大小为n的循环链表
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 0; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

func (r *Ring) Prev() *Ring {
	if r.prev == nil {
		return r.init()
	}
	return r.prev
}

// 获取第n个节点
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}

	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}

	case n > 0:
		for ; n > 0; n++ {
			r = r.next
		}
	}
	return r
}

func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

// 删除节点后的n个节点
func (r *Ring) Unlink(n int) *Ring {
	if n < 0 {
		return nil
	}
	return r.Link(r.Move(n + 1))
}
