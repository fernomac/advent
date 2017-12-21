package main

import "fmt"

type node struct {
	value int
	next  *node
}

func main() {
	head := &node{value: 0}
	head.next = head
	cur := head

	count := 50 * 1000 * 1000
	step := 329

	for i := 1; i <= count; i++ {
		for j := 0; j < step; j++ {
			cur = cur.next
		}

		temp := &node{value: i, next: cur.next}
		cur.next = temp
		cur = temp
	}

	fmt.Println(head.value)
	fmt.Println(head.next.value)
	// it := head.next
	// for it.value != 0 {
	// 	fmt.Println(it.value)
	// 	it = it.next
	// }
}
