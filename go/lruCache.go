package main

/*
import "fmt"

type CacheNode struct {
    v int
}

type DLLNode struct {
    v string
    next *DLLNode
    prev *DLLNode
}

type Cache struct {
    capacity int
    size int
    store map [string]*CacheNode
    list *DLLNode
}


type LRUCache interface {
    insert(key string, value  *CacheNode)
    get(key string) *CacheNode
    isFull() bool
    addFront(n *DLLNode)
    remove(key string) *DLLNode
    removeLast()
}

func NewCache(cap int) *Cache{
    return &Cache {
        capacity: cap,
        size: 0,
        store: map[string]*CacheNode{},
        list: nil,
    }
}

func NewDLLNode(k string) *DLLNode {
    return &DLLNode{
        v: k,
        prev: nil,
        next: nil,
    }
}

func (c *Cache) addFront(n *DLLNode) {
    if c.list != nil {
        n.next, c.list.prev = c.list, n
    }
    c.list = n
}

func (c *Cache) remove(key string) *DLLNode{
    q := c.list
    for q != nil {
        if q.v == key {
            p, n := q.prev, q.next
            if (n != nil) {
                n.prev = p
            }
            if (p != nil) {
                p.next = n
            }
            q.prev, q.next = nil, nil
            return q
        }
        q = q.next
    }
    return nil
}

func (c *Cache) removeLast() {
    q := c.list
    for q != nil {
        if q.next == nil {
            p := q.prev
            p.next = nil
            delete(c.store, q.v)
            break
        }
        q = q.next
    }
}

func (c *Cache) isFull() bool {
    return c.size >= c.capacity
}

func (c *Cache) insert(key string, value *CacheNode) {
    if c.isFull() {
        c.removeLast()
        c.size--
    }

    c.addFront(NewDLLNode(key))

    c.size++
    c.store[key] = value
}

func (c *Cache) get(key string) *CacheNode {
    if ret, ok := c.store[key]; ok {
        r := c.remove(key)
        c.addFront(r)
        return ret
    }
    return nil
}

func print(l *DLLNode) {
    for l != nil {
        fmt.Println(l.v)
        l = l.next
    }
}
*/
