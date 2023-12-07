package main

type Node struct {
	data interface{}
	next *Node
}

func AddNode(head *Node, data interface{}) *Node {
	newNode := &Node{data: data, next: nil}
	if head == nil {
		return newNode
	}

	currentNode := head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = newNode
	return head
}

func DelNode(head *Node, data interface{}) *Node {
	if head.data == data {
		return head.next
	}

	pre, cur := head, head.next
	for cur != nil {
		if cur.data == data {
			pre.next = cur.next
			break
		}
		pre, cur = cur, cur.next
	}
	return head
}
