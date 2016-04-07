package main

import (
)

type Node struct {
	Id string
	Type SchemaType
	IsRoot bool
	Required bool
	Children []Node
}

type NodeParam struct {
	Id string
	Schema *JsonSchema
	RootSchema *JsonSchema
}

func NewNode(param NodeParam) (Node, error) {
	node := Node{}
	node.initialize(param)
	switch node.Type {
	case SchemaTypeObject:
		for k, v := range param.Schema.Properties {
			n, _ := NewNode(NodeParam{Id: k, Schema: &v, RootSchema:param.Schema})
			node.Children = append(node.Children, n)
		}
	case SchemaTypeArray:
		for _, v := range param.Schema.GetItemList() {
			n, _ := NewNode(NodeParam{Schema: &v, RootSchema:param.Schema})
			node.Children = append(node.Children, n)
		}
	}
	return node, nil
}

func (n *Node) initialize(param NodeParam) {
	if param.Schema.Ref != "" {

	}
	if param.Id != "" {
		n.Id = param.Id
	} else if param.Schema.Id != "" {
		n.Id = param.Schema.Id
	} else {
		n.Id = ""
	}
	n.Type = SchemaTypeFromString(param.Schema.Type)
	n.IsRoot = param.Schema == param.RootSchema
	n.Required = param.RootSchema != nil && param.RootSchema.IsRequired(param.Schema.Id)
	n.Children = []Node{}
}

