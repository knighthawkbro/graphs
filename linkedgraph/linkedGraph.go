package linkedgraph

import (
	"fmt"

	"github.com/knighthawkbro/graphs/linkedset"
)

// node (Private) - Defines the structure for each individual node in a linked list
type node struct {
	vertex    string // Value of Node
	neighbors linkedset.List
	next      *node // Pointer to the next Node
}

func (n *node) String() string {
	return fmt.Sprintf("%v: %v", n.vertex, n.neighbors)
}

// List (Public) - The container for all the linked nodes in a set
type List struct {
	head node // the begining node
	size int  // size of the list
}

// New (Public) - Returns an initialized list.
func New() *List { return new(List) }

// Size (Public) - Returns the length variable for the list as an integer
func (l *List) Size() int { return l.size }

// Add (Public) - Returns the node in a singly linked list, just adds to the front of the list
func (l *List) Add(vert, neigh string) {
	if vert == "" || neigh == "" {
		fmt.Println("Cannot add a nil value to set")
		return
	}
	for current := l.head.next; current != nil; current = current.next {
		if current.vertex == vert {
			current.neighbors.Add(neigh)
			return
		}
	}
	new := &node{vertex: vert, next: l.head.next}
	l.head.next = new
	new.neighbors.Add(neigh)
	l.size++
	return
}

// Remove (Public) - Removes the first item on a list and returns it
func (l *List) Remove() string {
	if l.size == 0 {
		return "<nil>"
	}
	removed := l.head.neighbors.String()
	l.head = *l.head.next
	l.size--
	return removed
}

// GetNeighbors (Public) - Gets the neighbors at a vertex
func (l *List) GetNeighbors(vert string) string {
	for current := l.head.next; current != nil; current = current.next {
		if current.vertex == vert {
			return current.neighbors.String()
		}
	}
	return "[]"
}

// Contains (Public) - Returns true or false whether an item was contained in the list
func (l *List) Contains(vert string) bool {
	for current := l.head.next; current != nil; current = current.next {
		if current.vertex == vert {
			return true
		}
	}
	return false
}

// RemoveVert (Public) - finds a vertex and removes it from the list, returns <nil> if nothing found
func (l *List) RemoveVert(vert string) string {
	if l.size == 0 {
		return "<nil>"
	}
	var current = l.head.next
	for current != nil {
		if current.vertex == vert {
			removed := current.neighbors.String()
			l.head.next = l.head.next.next
			l.size--
			return removed
		}
		if current.next.vertex == vert {
			break
		}
		current = current.next
	}
	if current == nil {
		return "<nil>"
	}
	removed := current.next.neighbors.String()
	current.next = current.next.next
	l.size--
	return removed
}

// ContainsNeighbor (Public) -
func (l *List) ContainsNeighbor(vert, neigh string) bool {
	for current := l.head.next; current != nil; current = current.next {
		if current.vertex == vert {
			return current.neighbors.Contains(neigh)
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
		result += fmt.Sprintf("%v: %v\n", current.vertex, current.neighbors.String())
		if current.next == nil {
			result = result[:len(result)-1]
		}
	}
	return result + "]"
}
