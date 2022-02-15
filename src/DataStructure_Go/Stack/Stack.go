package Stack

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new Stack
func NewStack() *Stack {
	return &Stack{nil, 0}
}

// Return the number of items in the Stack
func (this *Stack) Len() int {
	return this.length
}

// View the top item on the Stack
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

// Pop the tope item of the Stack and return it
func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the Stack
func (this *Stack) Push(value interface{}) {
	n := &node{value: value, prev: this.top}
	this.top = n
	this.length++
}

// return true or false about Stack of length
func (this *Stack) isEmpty() bool {
	if this.length == 0 {
		return true
	}
	return false
}
