package List

import (
	"fmt"
)

func GetList() LinkedList {
	return LinkedList{head: nil}
}

func (l LinkedList) IsEmpty() bool {
	if l.head == nil {
		return true
	}
	return false
}

func (l LinkedList) Length() int {
	res, point := 0, l.head
	for point != nil {
		res++
		point = point.next
	}
	return res
}

func (l *LinkedList) Add(data int) {
	node := linkedlistnode{
		data: data,
		next: nil,
	}
	if l.IsEmpty() {
		l.head = &node
	} else {
		point := l.head
		for point.next != nil {
			point = point.next
		}
		point.next = &node
	}
}

func (l *LinkedList) Insert(index int, data Object) bool {
	if l.Length() < index || index < 0 {
		return false
	} else {
		point := l.head
		for i := 0; i < index; i++ {
			point = point.next
		}
		node := linkedlistnode{
			data: data,
			next: point.next,
		}
		point.next = &node
	}
	return true
}

func (l *LinkedList) Delete(data Object) bool {
	point := l.head
	if point.data == data {
		l.head = point.next
		return true
	} else {
		for point.next != nil {
			if point.next.data == data {
				point.next = point.next.next
				return true
			}
			point = point.next
		}
	}
	return false
}

func (l LinkedList) Coontain(data Object) bool {
	point := l.head
	for point != nil {
		if point.data == data {
			return true
		}
	}
	return false
}

func (l LinkedList) PrintList() {
	point := l.head
	fmt.Printf("[")
	for point != nil {
		fmt.Printf("\t%v", point.data)
		point = point.next
	}
	fmt.Printf("\t]\n")
}
