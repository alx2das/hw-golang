package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	Clear()
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

// Len вернет длину списка
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
	if i == nil || l.size == 0 {
		return
	}

	// если узел в начале списка
	if i == l.head && i != l.tail {
		l.head = i.Next
		l.head.Prev = nil
	}

	// если узел в конце списка
	if i != l.head && i == l.tail {
		l.tail = i.Prev
		l.tail.Next = nil
	}

	// если узел находится где-то в середине
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	l.size--
}

// MoveToFront переместит элемент в начало
func (l *list) MoveToFront(i *ListItem) {
	// если узла нет или он уже в начале
	if i == nil || l.head == i {
		return
	}

	// если узел в начале или в конце списка
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Prev.Prev = i.Prev
	}

	// если узел был последним элементом
	if l.tail == i {
		l.tail = i.Prev
	}

	// перемещаем узел в начало
	i.Next = l.head
	if l.head != nil {
		l.head.Prev = i
	}
	l.head = i
	i.Prev = nil
}

// Clear очистит список
// сборщик мусора автоматически освободит память для всех остальных узлов
func (l *list) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}
