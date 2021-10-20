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
	key   string
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if i, ok := c.items[key]; ok {
		i.Value = &cacheItem{
			key:   string(key),
			value: value,
		}
		c.queue.MoveToFront(i)
		return true
	}

	c.queue.PushFront(&cacheItem{
		key:   string(key),
		value: value,
	})
	c.items[key] = c.queue.Front()

	if c.capacity < c.queue.Len() {
		lru := c.queue.Back()
		delete(c.items, Key(lru.Value.(*cacheItem).key))
		c.queue.Remove(lru)
	}
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if e, ok := c.items[key]; ok {
		c.queue.MoveToFront(e)
		return e.Value.(*cacheItem).value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = NewList()
}
