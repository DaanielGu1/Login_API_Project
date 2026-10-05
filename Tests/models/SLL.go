package models

type SLL struct {
	lenght int
	rear   *node
	tail   *node
}

func NewList() *SLL {
	return &SLL{lenght: 0}
}

func (list *SLL) Add(data int) {
	n := newNode(data)

	if list.lenght == 0 {
		list.rear = n
		list.tail = n
	} else {
		list.tail.next = n
		list.tail = n
	}

	list.lenght++
}

func (list *SLL) Lenght() int {
	return list.lenght
}

func (list *SLL) Remove(index int) (int, bool) {
	lastIndex := list.lenght - 1

	if index > lastIndex || index < 0 || list.rear == nil {
		return 0, false
	}

	if index == 0 {
		deleted := list.rear.data
		list.lenght--

		if list.rear == list.tail {
			list.rear = nil
			list.tail = nil
			return deleted, true
		}

		list.rear = list.rear.next
		return deleted, true
	}

	previousNode := list.rear

	for i := 1; i < index-1; i++ {
		previousNode = previousNode.next
	}

	n := previousNode.next
	deleted := previousNode.next.data

	previousNode.next = n.next
	list.lenght--

	return deleted, true
}

func (list *SLL) Find(num int) int {
	previousNode := list.rear
	lastIndex := list.lenght - 1
	
	for i := 0; i < lastIndex; i++ {
		listNum := previousNode.data
		previousNode = previousNode.next

		if num == listNum {
			return i

		}
	}

	return -1
}

/* 
func (list *SLL) removeFirst() (int, bool) {
	lastIndex := list.lenght - 1
	deleted := list.rear.data

	if lastIndex == 0 {
		list.rear = nil
		list.tail = nil

	} else if lastIndex == 1{
		list.rear = list.tail

	} else if lastIndex > 1 {
		list.rear = list.rear.next
		
	} else {
		return 0, false
	}

	list.lenght--

	return deleted, true
}
*/