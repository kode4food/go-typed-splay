package splaytree

// A stack to remember parent nodes when
// traversing a tree.
type stack[Item any] struct {
	nodes []*node[Item]
}

func newStack[Item any]() *stack[Item] {
	return &stack[Item]{nodes: make([]*node[Item], 0, 10)}
}

func (stack *stack[Item]) push(node *node[Item]) {
	stack.nodes = append(stack.nodes, node)
}

func (stack *stack[_]) length() int {
	return len(stack.nodes)
}

func (stack *stack[_]) empty() bool {
	return len(stack.nodes) == 0
}

func (stack *stack[Item]) tos() *node[Item] {
	return stack.nodes[len(stack.nodes)-1]
}

func (stack *stack[Item]) pop() *node[Item] {
	node := stack.nodes[len(stack.nodes)-1]
	stack.nodes = stack.nodes[:len(stack.nodes)-1]
	return node
}
