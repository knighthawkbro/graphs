package linkedset

import "fmt"

// node (Private) - Defines the structure for each individual node in a linked list
type node struct {
	data string // Value of Node
	next *node  // Pointer to the next Node
}

// List (Public) - The container for all the linked nodes in a set
type List struct {
	head node // the begining node
	size int  // size of the list
}

// init (Private) - Generates a linked list with Size=0 and head pointing to itself
func (l *List) init() *List {
	l.head.next = &l.head
	l.size = 0
	return l
}

// New (Public) - Returns an initialized list.
func New() *List { return new(List).init() }

// Size (Public) - Returns the length variable for the list as an integer
func (l *List) Size() int { return l.size }

// Add (Public) - Returns the node in a singly linked list, just adds to the front of the list
func (l *List) Add(item string) {
	if item == "" {
		fmt.Println("Cannot add a nil value to set")
		return
	}
	if l.Contains(item) {
		return
	}
	new := &node{data: item, next: l.head.next}
	l.head.next = new
	l.size++
	return
}

// Remove (Public) - Removes the first item on a list and returns it
func (l *List) Remove() string {
	if l.size == 0 {
		return "<nil>"
	}
	removed := l.head.data
	l.head = *l.head.next
	l.size--
	return removed
}

// Get (Public) - Returns the first item list
func (l *List) Get() string {
	if l.size == 0 {
		return "<nil>"
	}
	return l.head.data
}

// Contains (Public) - Returns true or false whether an item was contained in the list
func (l *List) Contains(vert string) bool {
	for current := l.head.next; current != nil; current = current.next {
		if current.data == vert {
			return true
		}
	}
	return false
}

// String (Public) - Allows for the fmt.Print* functions to print the list struct as a string.
func (l *List) String() string {
	if l.size == 0 {
		return "[ ]"
	}
	result := "[ "
	for current := l.head.next; current != nil; current = current.next {
		result += fmt.Sprintf("%v ", current.data)
	}
	return result + "]"
}
