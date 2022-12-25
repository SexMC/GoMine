package ai

import "math/rand"

const (
	Running = iota
	Success
	Failure
)

const (
	Sequence = iota
	Conditional
	Selector
)

const (
	Ascending = iota
	Descending
)

type Node interface {
	Status() int
	GetParent() *Node
	ShouldStart() bool
	TryEnd()
}

type ParentNode interface {
	Node

	GetChildren() []*Node
	HasChildren() bool
	Executing() *Node
	Next() *Node
}

type Leaf struct {
	Node

	status int
	parent *Node
}

type Tree struct {
	ParentNode

	status    int
	treeType  int
	parent    *Node
	children  []*Node
	executing *Node
}

func (b *Leaf) Status() int {
	return b.status
}

func (b *Leaf) GetParent() *Node {
	return b.parent
}

func (t *Tree) Status() int {
	return t.status
}

func (t *Tree) GetParent() *Node {
	return t.parent
}

func (t *Tree) GetChildren() []*Node {
	return t.children
}

func (t *Tree) HasChildren() bool {
	return len(t.children) != 0
}

func (t *Tree) Executing() *Node {
	return t.executing
}

func (t *Tree) Next() *Node {
	if !t.HasChildren() || t.Executing() == nil {
		return nil
	}

	switch t.treeType {
	case Sequence:
		for i, c := range t.GetChildren() {
			if c == t.executing && i != len(t.children)-1 {
				return t.children[i+1]
			}
		}
	case Conditional:
		return t.children[0]
	case Selector:
		i := rand.Intn(len(t.children))
		return t.children[i]
	}

	return nil
}
