package listnode

import "errors"

type ListNode struct {
	Val  any
	Next *ListNode
}

type List struct {
	Head *ListNode
}

func NewList() *List {
	return &List{Head: nil}
}

/*

list feture:
	-to array
	-append
	-prepend
	-delete
	-size
*/

func (l *List) ToArray() []any {
	var result []any
	current := l.Head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

func (l *List) Append(value any) {
	newNode := &ListNode{Val: value}
	if l.Head == nil {
		l.Head = newNode
		newNode.Next = nil
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (l *List) Prepend(value any) {
	newNode := &ListNode{Val: value}
	if l.Head == nil {
		l.Head = newNode
		newNode.Next = nil
		return
	}
	newNode.Next = l.Head
	l.Head = newNode
}

func (l *List) Size() int {
	var count int
	current := l.Head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

func (l *List) Delete(value any) (bool, any, error) {
	if l.Head == nil {
		return false, nil, errors.New("list is empty")
	}
	if l.Head.Val == value {
		l.Head = l.Head.Next
		return true, value, nil

	}
	current := l.Head
	for current.Next != nil {
		if current.Next.Val == value {
			current.Next = current.Next.Next
			return true, value, nil
		}
		current = current.Next
	}
	return false, nil, errors.New("value not found")
}
