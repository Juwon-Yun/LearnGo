package Stack

type (
	stack02 struct {
		top    *node
		length int
	}
	node02 struct {
		value interface{}
		prev  *node
	}
)

// Create a new stack02
func NewStack02() *stack02 {
	return &stack02{nil, 0}
}

// Return the number of items in the stack02
func (this *stack02) Len() int {
	return this.length
}

// View the top item on the stack02
func (this *stack02) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

// Pop the tope item of the stack02 and return it
func (this *stack02) Pop() interface{} {
	if this.length == 0 {
		return nil
	}
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack02
func (this *stack02) Push(value interface{}) {
	n := &node{value: value, prev: this.top}
	this.top = n
	this.length++
}

func (this *stack02) isEmpty() bool {
	if this.length == 0 {
		return true
	}
	return false
}
