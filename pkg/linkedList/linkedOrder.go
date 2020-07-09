package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	head := createLinkedList(a)
	travelLinkedList(head)
	nLink := insert(head, 3)
	travelLinkedList(nLink)
}

type Node struct {
	data int
	next *Node
}

var linkIndex = -1

func createLinkedList(a []int) *Node {
	linkIndex += 1
	if linkIndex >= len(a) {
		return nil
	}
	n := Node{data: a[linkIndex]}
	n.next = createLinkedList(a)
	return &n
}

// 遍历
func travelLinkedList(head *Node) {
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}

// 插入：升序
func insert(head *Node, num int) *Node {
	if num <= head.data {
		q := &Node{data: num}
		q.next = head
		return q
	}
	p := head
	for head.next != nil {
		if num <= head.next.data {
			q := &Node{data: num}
			q.next = head.next
			head.next = q
			break
		}
		head = head.next
	}
	return p
}

// 删除：多个值都删除
func delete(a int, head *Node) *Node {
	n := head
	for n != nil && n.next != nil {
		if n.next.data == a {
			n.next = n.next.next
			continue
		}
		n = n.next
	}
	return n
}
