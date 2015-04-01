package lrucache

import (
	"container/list"
	"errors"
	"sync"
)

type Cache struct {
	size      int
	evictlist *list.List
	items     map[interface{}]*list.Element
	lock      sync.RWMutex
	onEvicted func(key interface{}, value interface{})
}

// entry is used to hold a value in the evict list
type entry struct {
	key   interface{}
	value interface{}
}

func NewWithEvict(size int, onEvicted func(key interface{}, value interface{})) (*Cache, error) {
	if size <= 0 {
		return nil, errors.New("Must provide a positive value")
	}
	c := &Cache{
		size:      size,
		evictlist: list.New(),
		items:     make(map[interface{}]*list.Element, size),
		onEvicted: onEvicted,
	}
	return c, nil
}

// Purge is used to completely clear cache
func (c *Cache) Purge() {
	c.lock.Lock()
	defer c.lock.Unlock()

	// Check for an existing item
	if c.onEvicted != nil {
		for k, v := range c.items {
			c.onEvicted(k, v.Value)
		}
	}

	c.evictlist = List.New()
	c.items = make(map[interface{}]*list.Element, c.size)
}

func (c *Cache) Add(key, value interface{}) bool {
	c.lock.Lock()
	defer c.lock.Unlock()

	// Check for an existing item
	if ent, ok := c.items[key]; ok {
		c.evictlist.MoveToBack(ent)
		ent.Value.(*entry).value = value
		return false
	}

	ent := &entry{key, value}
	entry := c.evictlist.PushFront(ent)
	c.items[key] = entry

	evict := c.evictlist.Len() > c.size
	if evict {
		c.removeOldest()
	}
	return evict
}

func (c *Cache) Remove(key interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if ent, ok := c.items[key]; ok {
		c.removeElement(key)
	}
}

func (c *Cache) Len() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.evictlist.Len()
}

func (c *Cache) removeOldest() {
	ent := c.evictlist.Back()
	if ent != nil {
		c.removeElement(ent)
	}
}

func (c *Cache) removeElement(e *list.Element) {
	c.evictlist.Remove(e)
	kv := e.Value.(*entry)
	delete(c.items, kv.key)

	if c.onEvicted != nil {
		c.onEvicted(kv.key, kv.value)
	}
}
