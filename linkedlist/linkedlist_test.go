package linkedlist

import (
	"strconv"
	"testing"
)

func TestCreateNode(t *testing.T) {
	testCases := []struct {
		data string
	}{
		{data: ""},
		{data: "Test Node"},
	}

	for idx, test := range testCases {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := newNode(test.data)
			if got.data != test.data {
				t.Errorf("bad data on node - got %s want %s", got.data, test.data)
			}
		})
	}
}

func TestNewLinkedList(t *testing.T) {
	got := newLinkedList()

	if got.head == nil {
		t.Error("head node of linked list is nil")
	}
}

func TestIncrementNode(t *testing.T) {
	linkedList := newLinkedList()

	nodeA := newNode("Node A")
	nodeB := newNode("Node B")
	nodeC := newNode("Node C")

	var err error

	err = linkedList.IncrementNode(nodeA)
	if err != nil {
		t.Error(err)
	}

	err = linkedList.IncrementNode(nodeB)
	if err != nil {
		t.Error(err)
	}

	err = linkedList.IncrementNode(nodeC)
	if err != nil {
		t.Error(err)
	}

	if linkedList.getLastNode() != nodeC {
		t.Errorf("the final node in the linked list is not the last node that was incremented - got %+v want %+v", linkedList.getLastNode(), nodeC)
	} else if linkedList.head.next != nodeA {
		t.Errorf("the node after head node is not the expected node - got %+v want %+v", linkedList.head.next, nodeA)
	} else if nodeA.next != nodeB {
		t.Errorf("the node after head node is not the expected node - got %+v want %+v", nodeA.next, nodeB)
	} else if nodeB.next != nodeC {
		t.Errorf("the node after head node is not the expected node - got %+v want %+v", nodeB.next, nodeC)
	} else if nodeC.next != nil {
		t.Errorf("expected the \"next\" node of the final node (nodeC) to be nil - got %+v", nodeC.next)
	}
}
