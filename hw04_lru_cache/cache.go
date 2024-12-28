package hw04lrucache

import "sync"

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
	mu       sync.Mutex
}

type cacheItem struct {
	key   Key
	value interface{}
}

// NewCache создает LRU-кэш заданной емкости.
func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

// Set добавит элемент в кеш или обновляет существующий.
func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	// если элемент присутствует в словаре
	if node, exist := c.items[key]; exist {
		node.Value.(*cacheItem).value = value

		// перемещаем элемент в начало очереди
		c.queue.MoveToFront(node)
		return true
	}

	// если кеш достиг емкости, удалим самый старый элемент
	if c.queue.Len() >= c.capacity {
		tailNode := c.queue.Back()
		if tailNode != nil {
			tailItem := tailNode.Value.(*cacheItem)

			delete(c.items, tailItem.key)
			c.queue.Remove(tailNode)
		}
	}

	// добавляем новый элемент
	newItem := &cacheItem{key, value}
	newNode := c.queue.PushFront(newItem)
	c.items[key] = newNode

	return false
}

// Get возвращаем элемент по ключу, если он существует.
func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if node, exist := c.items[key]; exist {
		// перемещаем элемент в начало очереди
		c.queue.MoveToFront(node)
		return node.Value.(*cacheItem).value, true
	}

	return nil, false
}

// Clear полностью очищает кеш.
func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
