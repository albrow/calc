package ast

import "fmt"

type Node interface {
	Children() []Node
	Parent() Node
	AddChild(Node)
	AddChildren([]Node)
	Copy() Node
	Draw()
	Format(int) string
	setParent(Node)
}

type BaseNode struct {
	children []Node
	parent   Node
}

func New() Node {
	return &BaseNode{}
}

func (n BaseNode) Children() []Node {
	return n.children
}

func (n BaseNode) Parent() Node {
	return n.parent
}

func (n *BaseNode) setParent(parent Node) {
	n.parent = parent
}

func (n *BaseNode) AddChild(child Node) {
	child.setParent(n)
	n.children = append(n.children, child)
}

func (n *BaseNode) AddChildren(children []Node) {
	for _, child := range children {
		n.AddChild(child)
	}
}

func (old *BaseNode) Copy() Node {
	newNode := &BaseNode{}
	for _, child := range old.children {
		newNode.AddChild(child.Copy())
	}
	return newNode
}

func (n BaseNode) Draw() {
	fmt.Println(n.Format(0))
}

func (n BaseNode) Format(depth int) string {
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
	BaseNode
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
	BaseNode
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
