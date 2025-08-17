package main

import (
	"container/list"
	//"fmt"
)

type LRUCache struct {
	capacity int
	size     int
	store    map[int]int
	list     *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		size:     0,
		store:    map[int]int{},
		list:     list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if ret, ok := c.store[key]; ok {
		c.list.MoveToFront(c.list.Back())
		return ret
	}
	return -1
}

func (c *LRUCache) isFull() bool {
	return c.size > c.capacity
}

func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.store[key]; !ok {
		c.size++
	}

	if c.isFull() {
		b := c.list.Back()
		delete(c.store, b.Value.(int))
		c.list.Remove(b)
		c.size--
	}

	c.list.PushFront(key)
	c.store[key] = value
}

/*
func (c *LRUCache) printList() {
    q := c.listHead
    for q != nil {
        fmt.Printf("%v ", q.v)
        q = q.next
    }
    fmt.Println()
}
*/
