package main

type StructureNode struct {
	Name string
	Properties []PropertyNode
	Children []StructureNode
}

type PropertyNode struct {
	Name string
	Type TypeNode
	IsRequired bool
}

type TypeNode struct {
	Name string
	InnerType *TypeNode
}

func (t TypeNode) EntityName() (string) {
	if t.InnerType != nil {
		return t.InnerType.EntityName()
	}
	return t.Name
}
