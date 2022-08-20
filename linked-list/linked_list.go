package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(newNode *node) {
	second := l.head
	l.head = newNode
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	nodeToPrint := l.head
	fmt.Printf("linked List: ")
	for l.length != 0 {
		fmt.Printf("{%v}-> ", nodeToPrint.data)
		nodeToPrint = nodeToPrint.next
		l.length--
	}
	fmt.Println()
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		fmt.Println("There's nothing to delete, the linked list is empty")
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			fmt.Printf("%v does not exist in the linked list\n", value)
			return
		}

		previousToDelete = previousToDelete.next
	}

	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	fmt.Println("testing the prepend method for a linked list")
	fmt.Println("--------------------------------------------")
	var myLinkedList linkedList
	node1, node2, node3, node4 := &node{data: 43}, &node{data: 22}, &node{data: 15}, &node{data: 8}
	node5, node6, node7, node8 := &node{data: 33}, &node{data: 108}, &node{data: 67}, &node{data: 5}
	myLinkedList.prepend(node1)
	myLinkedList.prepend(node2)
	myLinkedList.prepend(node3)
	myLinkedList.prepend(node4)
	myLinkedList.prepend(node5)
	myLinkedList.prepend(node6)
	myLinkedList.prepend(node7)
	myLinkedList.prepend(node8)

	fmt.Println(myLinkedList) // should print {0xc00009e240 4} (memory location of the head and the length of the linked list)
	fmt.Println()

	fmt.Println("testing the printListData method for a linked list")
	fmt.Println("--------------------------------------------------")
	myLinkedList.printListData() // should print linked List: {data: 8} {data: 15} {data: 22} {data: 43}
	fmt.Println()

	fmt.Println("testing the deleteWithValue method for a linked list")
	fmt.Println("----------------------------------------------------")
	myLinkedList.deleteWithValue(node6.data)
	myLinkedList.deleteWithValue(300)
	myLinkedList.deleteWithValue(node8.data)
	myLinkedList.printListData()
	var secondLinkedList linkedList
	secondLinkedList.deleteWithValue(100)
	secondLinkedList.printListData()

}
