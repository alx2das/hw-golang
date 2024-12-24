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

func (l list) Len() int {
	//TODO implement me
	panic("implement me")
}

func (l list) Front() *ListItem {
	//TODO implement me
	panic("implement me")
}

func (l list) Back() *ListItem {
	//TODO implement me
	panic("implement me")
}

func (l list) PushFront(v interface{}) *ListItem {
	//TODO implement me
	panic("implement me")
}

func (l list) PushBack(v interface{}) *ListItem {
	//TODO implement me
	panic("implement me")
}

func (l list) Remove(i *ListItem) {
	//TODO implement me
	panic("implement me")
}

func (l list) MoveToFront(i *ListItem) {
	//TODO implement me
	panic("implement me")
}

func NewList() List {
	return new(list)
}
