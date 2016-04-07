package main

import "syscall"

type DocumentNode struct {
	Name string
	Structure StructureNode
	Dependencies []DocumentNode
}

type StructureNode struct {
	Name string
	Properties []PropertyNode
	Children []StructureNode
}

func NewDocumentNodes(bundler JsonSchemaBundler) []DocumentNode {
	nodes := []DocumentNode{}
	b := bundler.GetRootBundle()
	n := DocumentNode{}
	n.Structure = NewStructureNode(bundler, b, "")
}

func NewStructureNode(bundler JsonSchemaBundler, schema *JsonSchema, name string) (StructureNode) {
	if schema.Type != SchemaTypeObject {
		panic("schema must be object for create structure node")
	}
	props := []PropertyNode{}
	dependencies := map[string]JsonSchema{}
	children := []StructureNode{}
	for k, s := range schema.Properties {
		if s.Ref != "" {
			s = bundler.GetBundle(schema.Ref).schema
		}
		typ := schema.Type
		if typ == SchemaTypeObject {
			typ = k
			dependencies[typ] = s
		}
		props = append(props, PropertyNode{
			k,
			typ,
			schema.IsRequired(k),
		})
	}
	for k, s := range dependencies {
		children = append(children, NewStructureNode(bundler, s, k))
	}
	return StructureNode{
		name,
		props,
		children,
	}
}

type PropertyNode struct {
	Name string
	Type string
	IsRequired bool
}

func NewPropertyNode(key string, schema *JsonSchema, isRequired bool) {
	return PropertyNode{
		key,
		schema.Type,
		isRequired,
	}
}
