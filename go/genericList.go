package main

import "fmt"

// List  represnts a singly-list that holds
// value of any type.

type Node[T any] struct {
	val  T
	next *Node[T]
}

type GenericList[T any] struct {
	head *Node[T]
}

func (list *GenericList[T]) add(x T) {
	if list.head == nil {
		list.head = &Node[T]{val: x, next: nil}
		return
	}
	list.head = &Node[T]{val: x, next: list.head}
}

func (list *GenericList[T]) remove() {
	if list.head == nil {
		return
	}
	list.head = list.head.next
}

func (list *GenericList[T]) count() int {
	count := 0
	temp := list.head
	for temp != nil {
		temp = temp.next
		count++
	}
	return count
}

func (list *GenericList[T]) print() {
	t := list.head
	for t != nil {
		fmt.Print(t.val, " ")
		t = t.next
	}
}
