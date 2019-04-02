package main

import "fmt"

func findNode(head *Node, index int, cnt *int) *Node {
	// recursive한 함수 안에서 반환값이 가지는 의미: basket
	if head.next == nil {
		return nil
	}
	node := findNode(head.next, index, cnt)
	*cnt += 1

	if *cnt == index {
		return head
	}
	return node
}

func findNodeWithRunner(list *LinkedList, index int) *Node {
	// runner 패턴은 아주 유용하다
	if list.head == nil {
		return nil
	}
	p1, p2 := list.head, list.head
	for i := 0; i < index; i++ {
		if p2.next == nil {
			return nil
		}
		p2 = p2.next
	}

	for p2.next != nil {
		p1 = p1.next
		p2 = p2.next
	}
	return p1
}

func main() {
	node5 := Node{data: 5, next: nil}
	node4 := Node{data: 4, next: &node5}
	node3 := Node{data: 3, next: &node4}
	node2 := Node{data: 2, next: &node3}
	node1 := Node{data: 1, next: &node2}
	llist := LinkedList{head: &node1}

	//node := findNodeWithRunner(&llist, 3)
	cnt := 0
	node := findNode(llist.head, 3, &cnt)
	fmt.Println(node)
}
