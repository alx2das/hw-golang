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
	tail *ListItem
	size int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.size
}

func (l *list) Front() *ListItem {
	if l.size == 0 {
		return nil
	}

	return l.head
}

func (l *list) Back() *ListItem {
	if l.size == 0 {
		return nil
	}

	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	//TODO implement me
	panic("implement me")
}

func (l *list) PushBack(v interface{}) *ListItem {
	//TODO implement me
	panic("implement me")
}

func (l *list) Remove(i *ListItem) {
	//TODO implement me
	panic("implement me")
}

func (l *list) MoveToFront(i *ListItem) {
	//TODO implement me
	panic("implement me")
}
