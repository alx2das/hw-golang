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

// NewList создаст двусвязный список.
func NewList() List {
	return new(list)
}

// Len вернет длину списка.
func (l *list) Len() int {
	return l.size
}

// Front вернет первый элемент списка.
func (l *list) Front() *ListItem {
	return l.head
}

// Back вернет последний элемент списка.
func (l *list) Back() *ListItem {
	return l.tail
}

// PushFront добавит значение в начало.
func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v, Next: l.head}

	if l.head != nil {
		l.head.Prev = newItem
	} else {
		// список пуст, обновляем tail
		l.tail = newItem
	}

	l.head = newItem
	l.size++

	return newItem
}

// PushBack добавит значение в конец.
func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v, Prev: l.tail}

	if l.tail != nil {
		l.tail.Next = newItem
	} else {
		// список пуст, обновляем head
		l.head = newItem
	}

	l.tail = newItem
	l.size++

	return newItem
}

// Remove удалит элемент.
func (l *list) Remove(i *ListItem) {
	if i == nil || l.size == 0 {
		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		// удаляем head
		l.head = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		// удаляем tail
		l.tail = i.Prev
	}

	l.size--
}

// MoveToFront переместит элемент в начало.
func (l *list) MoveToFront(i *ListItem) {
	if i == nil || l.head == i {
		return
	}

	// удаляем элемент из текущей позиции
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	// если элемент был в tail, обновляем
	if l.tail == i {
		l.tail = i.Prev
	}

	// перемещаем элемент в начало
	i.Next = l.head
	i.Prev = nil
	if l.head != nil {
		l.head.Prev = i
	}
	l.head = i
}
