package ast

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestChildren(t *testing.T) {
	parent := &baseNode{}
	children := []Node{
		&Number{
			Value: "1",
		},
		&Number{
			Value: "2",
		},
		&Number{
			Value: "3",
		},
	}
	parent.children = children
	if gotChildren := parent.Children(); !reflect.DeepEqual(gotChildren, children) {
		t.Errorf(
			"Expected: %s\n  but got: %s",
			spew.Sdump(children),
			spew.Sdump(gotChildren),
		)
	}
}

func TestParent(t *testing.T) {
	parent := &baseNode{}
	child := &Operator{
		Class: OpAdd,
	}
	child.parent = parent
	if gotParent := child.Parent(); !reflect.DeepEqual(gotParent, parent) {
		t.Errorf(
			"Expected: %s\n  but got: %s",
			spew.Sdump(parent),
			spew.Sdump(gotParent),
		)
	}
}

func TestAddChild(t *testing.T) {
	parent := New()
	children := []Node{
		&Number{
			Value: "1",
		},
		&Number{
			Value: "2",
		},
		&Number{
			Value: "3",
		},
	}
	for _, child := range children {
		parent.AddChild(child)
	}
	if gotChildren := parent.Children(); !reflect.DeepEqual(gotChildren, children) {
		t.Fatalf(
			"Expected: %s\n  but got: %s",
			spew.Sdump(children),
			spew.Sdump(gotChildren),
		)
	}
	for _, child := range children {
		if gotParent := child.Parent(); !reflect.DeepEqual(gotParent, parent) {
			t.Fatalf(
				"Expected: %s\n  but got: %s",
				spew.Sdump(parent),
				spew.Sdump(gotParent),
			)
		}
	}
}

func TestCopy(t *testing.T) {
	original := New()
	other := original.Copy()
	child := &Number{
		Value: "3",
	}
	other.AddChild(child)
	if numChildren := len(original.Children()); numChildren != 0 {
		t.Errorf("Expected original to not have any children, but got: %d", numChildren)
	}
	childCopy := child.Copy()
	childCopy.(*Number).Value = "0"
	if origValue := child.Value; origValue != "3" {
		t.Errorf("Expected original child value to not be mutated, but got: %s", origValue)
	}
	otherCopy := other.Copy()
	if len(otherCopy.Children()) != len(other.Children()) {
		t.Errorf(
			"otherCopy did not have the correct number of children. Expected: %d but got %d",
			len(other.Children()),
			len(otherCopy.Children()),
		)
	}
	otherCopy.Children()[0].(*Number).Value = "5"
	if origValue := other.Children()[0].(*Number).Value; origValue != "3" {
		t.Errorf("Expected child of other to not be mutated, but got: %s", origValue)
	}
}

func TestFormat(t *testing.T) {
	base := New()
	base.AddChild(&Operator{
		Class: OpAdd,
	})
	secondLevel := []Node{
		&Operator{
			Class: OpSubtract,
		},
		&Number{
			Value: "1",
		},
	}
	for _, child := range secondLevel {
		base.Children()[0].AddChild(child)
	}
	thirdLevel := []Node{
		&Number{
			Value: "3",
		},
		&Number{
			Value: "2",
		},
	}
	for _, child := range thirdLevel {
		secondLevel[0].AddChild(child)
	}
	expected := `|- base
  |- +
    |- -
      |- 3
      |- 2
    |- 1
`
	if got := base.Format(0); got != expected {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n\n", expected, got)
	}
}
