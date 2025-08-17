package main

/*
type DLLNode struct {
    v int
    next *DLLNode
    prev *DLLNode
}

type DLList struct {
    head *DLLNode
    tail *DLLNode
    length int
}

func New() *DLList {
    return &DLList {
        head: nil,
        tail: nil,
        length: 0,
    }
}

func (d *DLList) prepend(item *DLLNode) {
    // List is empty
    if d.head == nil {
        d.head, d.tail = item, item
    } else {
        // List is not empty
        item.next, d.head.prev = d.head, item
        d.head = item
    }
    d.length++
}

func (d *DLList) insertAt(item *DLLNode, idx int) {
    node := get(idx)
    if node.next == nil && node.prev == nil {
        item.next, node.prev = node, item
        d.head = item
    } else if node.next == nil && node.prev != nil {
    } else if node.next != nil && node.prev == nil {
    }
    d.length++
}

func (d *DLList) append(item *DLLNode) {
    // List is empty
    if d.tail == nil {
        d.head, d.tail = item, item
    } else {
        // List is not emtpy
        item.prev, d.tail.next = d.tail, item
        d.tail = item
    }
    d.length++
}

func (d *DLList) remove (item *DLLNode) *DLLNode {
    nextNode := d.head
    for nextNode != nil {

        if nextNode.v == item.v {

            if nextNode.prev != nil && nextNode.next != nil {
                nextNode.prev.next, nextNode.next.prev =
                nextNode.next, nextNode.prev
            } else if nextNode.prev == nil && nextNode.next != nil {

                nextNode.next.prev = nil
                d.head = nextNode.next
            } else if nextNode.prev != nil && nextNode.next == nil {

                nextNode.prev.next = nil
                d.tail = nextNode.prev
            }

            nextNode.prev, nextNode.next = nil, nil
            d.length--
            return nextNode
        }
    }
    return nil
}

func (d *DLList) removeAt(idx int) *DLLNode {
    node := get(idx)

    if node.prev != nil && node.next != nil {

        node.prev.next, node.next.prev =
        node.next, node.prev
    } else if node.prev == nil && node.next != nil {

        node.next.prev = nil
        d.head = nextNode.next
    } else if node.prev != nil && node.next == nil {

        node.prev.next = nil
        d.tail = nextNode.prev
    }

    node.prev, node.next = nil, nil
    d.lenght--
    return node
}

func (d *DLList) get (idx int) *DLLNode {
    c, nextNode := 0, d.head
    for nextNode != nil {
        if c == idx {
            return nextNode
        }
        c++
        nextNode = nextNode.next
    }
    return nil
}

func (d *DLList) length() int {
    return d.length
}
*/
