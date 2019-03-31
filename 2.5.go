package main

import (
	"container/list"
	"fmt"
)

func addTwoList(one *list.List, two *list.List) *list.List {

	result := list.New()
	r := 0
	list1, list2 := one.Front(), two.Front()

	for list1 != nil || list2 != nil {
		v1, v2 := 0, 0
		if list1 != nil {
			v1 = list1.Value.(int)
		}

		if list2 != nil {
			v2 = list2.Value.(int)
		}

		v3 := v1 + v2 + r
		r = 0
		if v3 > 10 {
			r = 1
			v3 -= 10
		}

		result.PushFront(v3)

		if list1 != nil {
			list1 = list1.Next()
		}
		if list2 != nil {
			list2 = list2.Next()
		}
	}

	return result
}

func main() {
	list1 := list.New()
	list1.PushFront(6)
	list1.PushFront(1)
	list1.PushFront(7)

	list2 := list.New()
	list2.PushFront(2)
	list2.PushFront(9)
	list2.PushFront(5)
	list2.PushFront(4)

	result := addTwoList(list1, list2)
	for e := result.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
