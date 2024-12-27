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
	items    sync.Map
	mu       sync.RWMutex // Для синхронизаций операций с очередью
}

type cacheItem struct {
	key   Key
	value interface{}
}

// Set добавит элемент в кеш по ключу.
func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// если элемент присутствует в словаре
	if existingNode, exist := c.items.Load(key); exist {
		node := existingNode.(*ListItem)
		node.Value.(*cacheItem).value = value

		// обновить значение и переместить элемент в начало очереди
		c.items.Store(key, node)
		c.queue.MoveToFront(node)

		// успешно обновлен
		return true
	}

	// добавить в словарь и в начало очереди
	if c.queue.Len() >= c.capacity {
		tailNode := c.queue.Back()
		if tailNode != nil {
			tailItem := tailNode.Value.(*cacheItem)

			c.items.Delete(tailItem.key)
			c.queue.Remove(tailNode)
		}
	}

	// добавляем элемент
	newItem := &cacheItem{key, value}
	newNode := c.queue.PushFront(newItem)

	c.items.Store(key, newNode)

	// успешно добавлен
	return false
}

// Get вернет элемент из кеша по ключу.
func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// если элемент присутствует в словаре
	if existingNode, exist := c.items.Load(key); exist {
		node := existingNode.(*ListItem)
		// переместить элемент в начало очереди
		c.queue.MoveToFront(node)

		return node.Value.(*cacheItem).value, true
	}

	return nil, false
}

// Clear полностью очистит кеш.
func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.queue.Clear()
	c.items.Clear()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    sync.Map{},
	}
}
