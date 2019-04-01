package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func removeOverlap(list *LinkedList) {
	var arr [10]bool

	for e := list.head; e.next != nil; e = e.next {
		arr[e.data] = true
		for arr[e.next.data] == true {
			e.next = e.next.next
		}
	}
}

func main() {
	tail9 := Node{data: 6, next: nil}
	tail8 := Node{data: 5, next: &tail9}
	tail7 := Node{data: 5, next: &tail8}
	tail6 := Node{data: 4, next: &tail7}
	tail5 := Node{data: 4, next: &tail6}
	tail4 := Node{data: 4, next: &tail5}
	tail3 := Node{data: 3, next: &tail4}
	tail2 := Node{data: 3, next: &tail3}
	tail1 := Node{data: 2, next: &tail2}
	head := Node{data: 1, next: &tail1}
	llist := LinkedList{head: &head}

	removeOverlap(&llist)
	for e := llist.head; e.next != nil; e = e.next {
		fmt.Println(e.data)
	}
}
