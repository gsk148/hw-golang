package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	head *ListItem
	last *ListItem
	size int
}

func NewList() List {
	return new(list)
}

func (l list) Len() int {
	return l.size
}

func (l list) Front() *ListItem {
	return l.head
}

func (l list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l.Len() == 0 {
		return l.pushFirst(v)
	}
	front := l.Front()
	item := &ListItem{
		Value: v,
		Next:  front,
		Prev:  nil,
	}
	front.Prev = item
	l.head = item
	l.size++

	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.Len() == 0 {
		return l.pushFirst(v)
	}
	back := l.Back()
	item := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  back,
	}
	back.Next = item
	l.last = item
	l.size++

	return item
}

func (l *list) Remove(i *ListItem) {
	if l.Len() == 0 {
		return
	}
	if l.head == i {
		l.head = i.Next
	}
	if l.last == i {
		l.last = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	l.size--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func (l *list) pushFirst(v interface{}) *ListItem {
	item := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  nil,
	}
	l.size = 1
	l.head = item
	l.last = item

	return item
}
