package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	// если элемент присутствует в словаре
	if node, exists := c.items[key]; exists {
		node.Value.(*cacheItem).value = value
		c.queue.MoveToFront(node)
		return true
	}

	// если размер очереди больше ёмкости кэша
	if len(c.items) >= c.capacity {
		// удаляем последний элемент
		tailNode := c.queue.Back()
		if tailNode != nil {
			tailItem := tailNode.Value.(*cacheItem)
			delete(c.items, tailItem.key) // удаляем из словаря
			c.queue.Remove(tailNode)      // удаляем из списка
		}
	}

	// добавить в словарь и в начало очереди
	newItem := &cacheItem{key, value}
	newNode := c.queue.PushFront(newItem)
	c.items[key] = newNode

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if node, exists := c.items[key]; exists {
		c.queue.MoveToFront(node)
		return node.Value.(*cacheItem).value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.queue.Clear()
	c.items = make(map[Key]*ListItem)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
