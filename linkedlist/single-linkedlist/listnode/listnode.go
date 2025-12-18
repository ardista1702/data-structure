package listnode

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Head *ListNode
}

func NewList() *List {
	return &List{}
}

func (l *List) Insert(value int) {
	newNode := &ListNode{
		Val: value,
	}

	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head

	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode

}
