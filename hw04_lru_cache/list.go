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

// Len returns the length of the list
func (l *list) Len() int {
	return l.size
}

// Front вернет первый элемент списка
func (l *list) Front() *ListItem {
	if l.size == 0 {
		return nil
	}

	return l.head
}

// Back вернет последний элемент списка
func (l *list) Back() *ListItem {
	if l.size == 0 {
		return nil
	}

	return l.tail
}

// PushFront добавит значение в начало
func (l *list) PushFront(v interface{}) *ListItem {
	newListItem := &ListItem{Value: v}

	if l.size == 0 {
		l.head = newListItem
		l.tail = newListItem
	} else {
		newListItem.Next = l.head
		l.head.Prev = newListItem
		l.head = newListItem
	}

	l.size++
	return newListItem
}

// PushBack добавит значение в конец
func (l *list) PushBack(v interface{}) *ListItem {
	newListItem := &ListItem{Value: v}

	if l.size == 0 {
		l.head = newListItem
		l.tail = newListItem
	} else {
		newListItem.Prev = l.tail
		l.tail.Next = newListItem
		l.tail = newListItem
	}

	l.size++
	return newListItem
}

// Remove удалит элемент
func (l *list) Remove(i *ListItem) {
	//TODO implement me
	panic("implement me")
}

// MoveToFront переместит элемент в начало
func (l *list) MoveToFront(i *ListItem) {
	//TODO implement me
	panic("implement me")
}
