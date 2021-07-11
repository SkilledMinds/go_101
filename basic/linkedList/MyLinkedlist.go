package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type SinglyLinkedList struct {
	head   *Node
	opNode *Node
}

// Add new nodes to LinkedList
func (sl *SinglyLinkedList) add(someNum int) {

	newNode := &Node{
		value: someNum,
	}

	if sl.head == nil {
		sl.head = newNode
		sl.opNode = newNode
	} else {
		sl.opNode.next = newNode
		sl.opNode = newNode
	}
}

// Traverse to display value of each node
func (sld *SinglyLinkedList) displayList()  {
	list := sld.head
	for list != nil {
		fmt.Printf("%+v", list.value)
		list = list.next
		if(list!=nil){
			fmt.Printf(" -> ")
		}
	}
	fmt.Println()
}

func main() {
	sLinkList := SinglyLinkedList{}
	sLinkList.add(1)
	sLinkList.add(4)
	sLinkList.add(6)
	sLinkList.add(7)
	sLinkList.add(8)
	sLinkList.add(8)

	sLinkList.displayList()
}
