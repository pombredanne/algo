// Copyright 2013–2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

// Caches with least-recently used (LRU) eviction strategy.
package lru

// LRU cache.
type Cache struct {
	// Used by the Get method to generate values from keys.
	Func  func(interface{}) interface{}
	index map[interface{}]*link
	list
}

// Create a new LRU cache for function f with the desired capacity.
//
// f may be nil, in which case Get won't work (use Add and Check instead).
func New(f func(interface{}) interface{}, capacity int) *Cache {
	if capacity < 1 {
		panic("capacity < 1")
	}
	c := &Cache{Func: f, index: make(map[interface{}]*link)}
	c.init(capacity)
	return c
}

// Add key with associated value to c.
//
// Overwrites any value for key currently in c.
func (c *Cache) Add(key interface{}, value interface{}) {
	p, stored := c.index[key]
	if stored {
		p.value = value
		c.moveToFront(p)
	} else {
		c.insert(key, value)
	}
}

// Check whether key is in the cache and if so, return its associated value.
func (c *Cache) Check(key interface{}) (value interface{}, stored bool) {
	p, stored := c.index[key]
	if stored {
		value = p.value
		c.moveToFront(p)
	}
	return
}

// Calls f for each key, value pair in the cache, in order of descending
// recency (MRU first, LRU last).
func (c *Cache) Do(f func(k, v interface{})) {
	for p := c.head; p != nil; p = p.next {
		f(p.key, p.value)
	}
}

// Fetch value for key in the cache, calling Func to compute it if necessary.
func (c *Cache) Get(key interface{}) (value interface{}) {
	p, stored := c.index[key]
	if stored {
		value = p.value
		c.moveToFront(p)
	} else {
		value = c.Func(key)
		c.insert(key, value)
	}
	return
}

// Number of items currently in the cache.
func (c *Cache) Len() int {
	return len(c.links)
}

// Maximum number of items that c will store.
func (c *Cache) Capacity() int {
	return cap(c.links)
}

func (c *Cache) insert(key interface{}, value interface{}) {
	var p *link
	if c.full() {
		// evict least recently used item
		p = c.popTail()
		delete(c.index, p.key)
	} else {
		p = c.grow()
	}
	p.key, p.value = key, value
	c.putFront(p)
	c.index[key] = p
}

// Doubly linked list containing key/value pairs.
type list struct {
	head, tail *link
	links       []link
}

type link struct {
	key, value interface{}
	prev, next *link
}

// Initialize l with capacity c.
func (l *list) init(c int) {
	l.links = make([]link, 0, c)
}

func (l *list) full() bool {
	return len(l.links) == cap(l.links)
}

// Grow list by one link.
func (l *list) grow() *link {
	n := len(l.links)
	l.links = l.links[:n+1]
	return &l.links[n]
}

// Make p the head of the list. Precondition: l is not empty.
func (l *list) moveToFront(p *link) {
	if p == l.head {
		return
	}
	if p == l.tail {
		l.tail = p.prev
	}
	next, prev := p.next, p.prev
	if next != nil {
		next.prev = prev
	}
	if prev != nil {
		prev.next = next
	}

	l.head.prev, p.next, p.prev = p, l.head, nil
	l.head = p
}

// Pop the tail off the list and return it. Precondition: l is full.
func (l *list) popTail() (t *link) {
	t = l.tail
	l.tail = t.prev
	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}
	t.next, t.prev = nil, nil
	return
}

// Make p the head of the list.
func (l *list) putFront(p *link) {
	if l.head == nil {
		l.head, l.tail = p, p
		p.next, p.prev = nil, nil
	} else {
		l.moveToFront(p)
	}
	return
}
