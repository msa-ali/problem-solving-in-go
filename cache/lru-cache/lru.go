package cache

import (
	"github.com/google/uuid"
)

type node[T any] struct {
	id    string
	next  *node[T]
	prev  *node[T]
	value T
}

type LRU[T any] struct {
	capacity      int
	length        int
	head          *node[T]
	tail          *node[T]
	lookup        map[string]*node[T]
	reverseLookup map[string]string
}

func newNode[T any](value T) *node[T] {
	return &node[T]{
		id:    uuid.NewString(),
		next:  nil,
		prev:  nil,
		value: value,
	}
}

func New[T any](capacity int) *LRU[T] {
	return &LRU[T]{
		capacity:      capacity,
		lookup:        make(map[string]*node[T], capacity+1),
		reverseLookup: make(map[string]string, capacity+1),
	}
}

func (cache *LRU[T]) Update(key string, value T) {

	// does it exist?
	node, ok := cache.lookup[key]
	if !ok {
		node = newNode(value)
		cache.length++
		cache.prepend(node)
		cache.trimCache()
		cache.lookup[key] = node
		cache.reverseLookup[node.id] = key
	} else {
		cache.detach(node)
		cache.prepend(node)
		node.value = value
	}
}

func (cache *LRU[T]) Get(key string) (T, bool) {
	var value T
	// check the cache for existence
	node, ok := cache.lookup[key]
	if !ok {
		return value, false
	}
	// update the value and move it to front
	cache.detach(node)
	cache.prepend(node)

	// return the value or return zero value
	return node.value, true
}

func (cache *LRU[T]) detach(n *node[T]) {
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
	if cache.head == n {
		cache.head = cache.head.next
	}
	if cache.tail == n {
		cache.tail = cache.tail.prev
	}
	n.prev = nil
	n.next = nil
}

func (cache *LRU[T]) prepend(n *node[T]) {
	if cache.head == nil {
		cache.head = n
		cache.tail = n
	} else {
		n.next = cache.head
		cache.head.prev = n
		cache.head = n
	}
}

func (cache *LRU[T]) trimCache() {
	if cache.length <= cache.capacity {
		return
	}
	tail := cache.tail
	cache.detach(cache.tail)
	key := cache.reverseLookup[tail.id]
	delete(cache.lookup, key)
	delete(cache.reverseLookup, tail.id)
	cache.length--
}
