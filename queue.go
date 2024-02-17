package main
import "errors"

type Queue interface {
    enqueue(x int)
    dequeue() int
    peek() int
}

type List struct {
    val int
    next *List
}

func (l *List) enqueue(x int) {
    for l.next != nil {
        l = l.next
    }

    l.next = &List{
                val: x,
                next: nil,
            }
}

func (l *List) dequeue() (int, error) {
    if l == nil {
        return -1, errors.New("Can't dequeue from an empty list")
    }

    temp := l;
    l = l.next;
    return temp.val, nil
}

func (l *List) peek() (int, error) {
    if l == nil {
        return -1, errors.New("Can't peek into a an empty list")
    }
    return l.val, nil
}
