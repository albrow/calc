package ast

type Node interface {
	Children() []Node
	Parent() Node
	AddChild(Node)
	setParent(Node)
}

type baseNode struct {
	children []Node
	parent   Node
}

func (n baseNode) Children() []Node {
	return n.children
}

func (n baseNode) Parent() Node {
	return n.parent
}

func (n *baseNode) setParent(parent Node) {
	n.parent = parent
}

func (n *baseNode) AddChild(child Node) {
	child.setParent(n)
	n.children = append(n.children, child)
}

type OpType uint

const (
	opAdd OpType = iota
	opSubtract
)

type Operator struct {
	baseNode
	Operator OpType
}

type Number struct {
	baseNode
	Value string
}
