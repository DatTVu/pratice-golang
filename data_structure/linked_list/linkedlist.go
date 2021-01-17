package linked_list

import (
	"fmt"
	"log"
	"sync"
)

type node struct {
	value int
	next  *node
}

type LinkedList struct {
	mutex   sync.Mutex
	headptr *node
	size    int
}

//Print all elements of the list
func (l *LinkedList) Print() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	curPtr := l.headptr
	for curPtr != nil {
		fmt.Print(curPtr.value, " ")
		if curPtr.next != nil {
			curPtr = curPtr.next
		} else {
			break
		}
	}
}

//Append val to the end of the list
//by first finding the tail first
func (l *LinkedList) Append(val int) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	curPtr := l.headptr
	for curPtr != nil {
		if curPtr.next != nil {
			curPtr = curPtr.next
		} else {
			break
		}
	}

	newNode := node{
		value: val,
		next:  nil,
	}
	curPtr.next = &newNode
	l.size++
}

//Prepend val to the beginning of the list
func (l *LinkedList) Prepend(val int) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	newNode := &node{
		value: val,
		next:  l.headptr,
	}

	l.headptr = newNode
	l.size++
}

//Add a node at specific index
func (l *LinkedList) Add(val, index int) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if index > l.size {
		return
	}

	curPtr := l.headptr
	prevPtr := l.headptr

	for i := 0; i < index; i++ {
		if curPtr == nil {
			break
		} else {
			prevPtr = curPtr
			curPtr = curPtr.next
		}
	}

	prevPtr.next = &node{
		value: val,
		next:  curPtr,
	}

	l.size++
}

func (l *LinkedList) PopFront() (result int) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.headptr == nil {
		log.Fatal("Popping an empty linked list!")
	}
	curPtr := l.headptr
	result = curPtr.value
	l.headptr = l.headptr.next
	curPtr = nil // is this line necessary?
	l.size--
	return result
}

func (l *LinkedList) PopBack() (result int) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.headptr == nil {
		log.Fatal("Popping an empty linked list!")
	}
	prevPtr := l.headptr
	curPtr := l.headptr
	for curPtr != nil {
		if curPtr.next == nil {
			break
		} else {
			prevPtr = curPtr
			curPtr = curPtr.next
		}
	}
	result = curPtr.value
	l.size--
	prevPtr.next = nil
	curPtr = nil
	return result
}

func (l *LinkedList) RemoveAt(index int) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if index > l.size {
		log.Fatal("Out of range!")
	}

	prevPtr := l.headptr
	curPtr := l.headptr
	for i := 0; i < index; i++ {
		for curPtr != nil {
			prevPtr = curPtr
			curPtr = curPtr.next
		}
	}
	prevPtr.next = curPtr.next
	curPtr = nil
	l.size--
}

func (l *LinkedList) RemoveValue(value int) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	prevPtr := l.headptr
	curPtr := l.headptr
	for curPtr != nil {
		if curPtr.next == nil {
			log.Fatal("Element doesn't exist in the list")
		} else if curPtr.value == value {
			prevPtr.next = curPtr.next
			curPtr = nil
			return
		} else {
			prevPtr = curPtr
			curPtr = curPtr.next
		}
	}
}

func (l *LinkedList) Empty() (res bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.size == 0 {
		res = true
	} else {
		res = false
	}
	return res
}

func (l *LinkedList) Size() int {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return l.size
}

func (l *LinkedList) ValueAt(index int) (result int) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if index > l.size {
		log.Fatal("Out of Range")
	}
	curPtr := l.headptr
	i := 0

	for i < index {
		curPtr = curPtr.next
		i++
	}
	result = curPtr.value
	return result
}

func (l *LinkedList) Reverse() {

}

// Find an element in list and return boolean
// Find an element in list by index and return it
// Count how many times a key show up in list
// Detect if a list has been loop
// Check if single linked list is palindrome
// Remove duplicate element in list (sorted)
// Remove duplicate element in list (unsorted)
// Remove duplicate element in list (unsorted)
