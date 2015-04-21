// Copyright 2013–2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

// Caches for arbitrary functions with least-recently used (LRU) eviction
// strategy.
package lru

// Cache for function Func.
type Cache struct {
	Func  func(interface{}) interface{}
	index map[interface{}]int // index of key in queue
	list
}

// Create a new LRU cache for function f with the desired capacity.
func New(f func(interface{}) interface{}, capacity int) *Cache {
	if capacity < 1 {
		panic("capacity < 1")
	}
	c := &Cache{Func: f, index: make(map[interface{}]int)}
	c.init(capacity)
	return c
}

// Fetch value for key in the cache, calling Func to compute it if necessary.
func (c *Cache) Get(key interface{}) (value interface{}) {
	i, stored := c.index[key]
	if stored {
		value = c.valueAt(i)
		c.moveToFront(i)
	} else {
		value = c.Func(key)
		c.insert(key, value)
	}
	return value
}

// Number of items currently in the cache.
func (c *Cache) Len() int {
	return len(c.links)
}

func (c *Cache) Capacity() int {
	return cap(c.links)
}

func (c *Cache) insert(key interface{}, value interface{}) {
	var i int
	if c.full() {
		// evict least recently used item
		var k interface{}
		i, k = c.popTail()
		delete(c.index, k)
	} else {
		i = c.grow()
	}
	c.putFront(key, value, i)
	c.index[key] = i
}

// Doubly linked list containing key/value pairs.
type list struct {
	front, tail int
	links       []link
}

type link struct {
	key, value interface{}
	prev, next int
}

// Initialize l with capacity c.
func (l *list) init(c int) {
	l.front = -1
	l.tail = -1
	l.links = make([]link, 0, c)
}

func (l *list) full() bool {
	return len(l.links) == cap(l.links)
}

// Grow list by one element and return its index.
func (l *list) grow() (i int) {
	i = len(l.links)
	l.links = l.links[:i+1]
	return
}

// Make the node at position i the front of the list.
// Precondition: the list is not empty.
func (l *list) moveToFront(i int) {
	nf := &l.links[i]
	of := &l.links[l.front]

	nf.prev = l.front
	of.next = i
	l.front = i
}

// Pop the tail off the list and return its index and its key.
// Precondition: the list is full.
func (l *list) popTail() (i int, key interface{}) {
	i = l.tail
	t := &l.links[i]
	key = t.key
	l.links[t.next].prev = -1
	l.tail = t.next
	return
}

// Put (key, value) in position i and make it the front of the list.
func (l *list) putFront(key, value interface{}, i int) {
	f := &l.links[i]
	f.key = key
	f.value = value
	f.prev = l.front
	f.next = -1

	if l.tail == -1 {
		l.tail = i
	} else {
		l.links[l.front].next = i
	}
	l.front = i
}

func (l *list) valueAt(i int) interface{} {
	return l.links[i].value
}
