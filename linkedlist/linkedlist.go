package linkedlist

import "errors"

type LinkedList struct {
	head *Node
}

type Node struct {
	data string
	next *Node
}

func newNode(data string) *Node {
	return &Node{data: data}
}

func newLinkedList() *LinkedList {
	return &LinkedList{
		head: newNode("Root Node"),
	}
}

func (l *LinkedList) getLastNode() *Node {
	nodeCursor := l.head
	for nodeCursor.next != nil {
		nodeCursor = nodeCursor.next
	}
	return nodeCursor
}

func (l *LinkedList) IncrementNode(node *Node) error {
	if node == nil {
		return errors.New("a nil value cannot be incremented - a valid *Node pointer must be an argument")
	}

	leafNode := l.getLastNode()
	if leafNode == node {
		return errors.New("node increment error - the next node cannot be the same node")
	}

	leafNode.next = node
	return nil
}
