package leetcode

import "fmt"

type LinkNode struct {
	data interface{}
	next *LinkNode
}

type LinkedList struct {
	head *LinkNode
	size int
}

func (list *LinkedList) Size() int {
	return list.size
}

func (list *LinkedList) IsEmpty() bool {
	return list.size == 0
}

func (list *LinkedList) Add(data interface{}) {
	newLinkNode := &LinkNode{data: data, next: nil}

	if list.head == nil {
		list.head = newLinkNode
	} else {
		curr := list.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = newLinkNode
	}
	list.size++
}

func (list *LinkedList) InsertAt(index int, data interface{}) error {
	if index < 0 || index > list.size {
		return fmt.Errorf("Index out of range")
	}

	newLinkNode := &LinkNode{data: data, next: nil}

	if index == 0 {
		newLinkNode.next = list.head
		list.head = newLinkNode
	} else {
		curr := list.head
		for i := 0; i < index-1; i++ {
			curr = curr.next
		}
		newLinkNode.next = curr.next
		curr.next = newLinkNode
	}
	list.size++
	return nil
}

func (list *LinkedList) RemoveAt(index int) error {
	if index < 0 || index >= list.size {
		return fmt.Errorf("Index out of range")
	}

	if index == 0 {
		list.head = list.head.next
	} else {
		curr := list.head
		for i := 0; i < index-1; i++ {
			curr = curr.next
		}
		curr.next = curr.next.next
	}
	list.size--
	return nil
}

func (list *LinkedList) IndexOf(data interface{}) int {
	curr := list.head
	index := 0
	for curr != nil {
		if curr.data == data {
			return index
		}
		curr = curr.next
		index++
	}
	return -1
}

func (list *LinkedList) Contains(data interface{}) bool {
	return list.IndexOf(data) != -1
}

func (list *LinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.size {
		return nil, fmt.Errorf("Index out of range")
	}

	curr := list.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	return curr.data, nil
}

type Ring struct {
	next, prev *Ring       // 前驱和后驱节点
	Value      interface{} // 数据
}

func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
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
	if r.next == nil {
		return r.init()
	}
	return r.prev
}

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
		for ; n > 0; n-- {
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

func (r *Ring) Unlink(n int) *Ring {
	if n < 0 {
		return nil
	}
	return r.Link(r.Move(n + 1))
}

func (r *Ring) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.next {
			n++
		}
	}
	return n
}
