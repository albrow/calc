package ast

import "fmt"

type Node interface {
	Children() []Node
	Parent() Node
	AddChild(Node)
	Copy() Node
	Draw()
	Format(int) string
	setParent(Node)
}

type baseNode struct {
	children []Node
	parent   Node
}

func New() Node {
	return &baseNode{}
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

func (old *baseNode) Copy() Node {
	newNode := &baseNode{}
	for _, child := range old.children {
		newNode.AddChild(child.Copy())
	}
	return newNode
}

func (n baseNode) Draw() {
	fmt.Println(n.Format(0))
}

func (n baseNode) Format(depth int) string {
	indent := ""
	output := ""
	for i := 0; i < depth; i++ {
		indent += "  "
	}
	output += fmt.Sprintf("%s|- base\n", indent)
	depth++
	for _, child := range n.Children() {
		output += child.Format(depth)
	}
	return output
}

type OpClass uint

const (
	OpAdd OpClass = iota
	OpSubtract
)

func (c OpClass) String() string {
	switch c {
	case OpAdd:
		return "+"
	case OpSubtract:
		return "-"
	}
	panic(fmt.Sprintf("Unknown OpClass: %d", uint(c)))
}

type Operator struct {
	baseNode
	Class OpClass
}

func (old *Operator) Copy() Node {
	newNode := &Operator{
		Class: old.Class,
	}
	for _, child := range old.Children() {
		newNode.AddChild(child.Copy())
	}
	return newNode
}

func (n Operator) Format(depth int) string {
	indent := ""
	output := ""
	for i := 0; i < depth; i++ {
		indent += "  "
	}
	output += fmt.Sprintf("%s|- %s\n", indent, n.Class)
	depth++
	for _, child := range n.Children() {
		output += child.Format(depth)
	}
	return output
}

type Number struct {
	baseNode
	Value string
}

func (old *Number) Copy() Node {
	newNode := &Number{
		Value: old.Value,
	}
	for _, child := range old.Children() {
		newNode.AddChild(child.Copy())
	}
	return newNode
}

func (n Number) Format(depth int) string {
	indent := ""
	output := ""
	for i := 0; i < depth; i++ {
		indent += "  "
	}
	output += fmt.Sprintf("%s|- %s\n", indent, n.Value)
	depth++
	for _, child := range n.Children() {
		output += child.Format(depth)
	}
	return output
}
