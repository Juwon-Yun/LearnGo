package main

type (
	Stack02 struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new Stack02
func New() *Stack02 {
	return &Stack02{nil, 0}
}

// Return the number of items in the Stack02
func (this *Stack02) Len() int {
	return this.length
}

// View the top item on the Stack02
func (this *Stack02) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

// Pop the tope item of the Stack02 and return it
func (this *Stack02) Pop() interface{} {
	if this.length == 0 {
		return nil
	}
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the Stack02
func (this *Stack02) Push(value interface{}) {
	n := &node{value: value, prev: this.top}
	this.top = n
	this.length++
}
