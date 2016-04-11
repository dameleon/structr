package main

import (
	"strings"
)

type StructureNode struct {
	Name string
	Properties []PropertyNode
	Dependencies []StructureNode
}

type PropertyNode struct {
	Name string
	Type TypeNode
	IsRequired bool
}

type TypeNode struct {
	Name string
	ContainType *TypeNode
}

func (t TypeNode) EntityName() (string) {
	if t.ContainType != nil {
		return t.ContainType.EntityName()
	}
	return t.Name
}

func NewStructureNode(bundler JsonSchemaBundler, name string, rootSchema JsonSchema) (StructureNode) {
	if rootSchema.Type != SchemaTypeObject {
		panic("root schema must be type of object")
	}
	properties := []PropertyNode{}
	dependencyMap := make(map[string]StructureNode)
	for key, schema := range rootSchema.Properties {
		refString := ""
		if schema.Ref != "" {
			refString = schema.Ref
			schema = bundler.GetSchema(schema.Ref)
		}
		// create property
		prop := NewPropertyNode(key, schema, schema.IsRequired(key), refString)
		properties = append(properties, prop)
		// create dependency
		if containType := prop.Type.ContainType; containType != nil {
			name := containType.EntityName()
			s := schema
			if s.Type == SchemaTypeArray {
				s = schema.GetItemList()[0]
			}
			if s.Type == SchemaTypeObject {
				if _, ok := dependencyMap[name]; !ok {
					dependencyMap[name] = NewStructureNode(bundler, name, s)
				}
			}
		}
	}
	dependencies := []StructureNode{}
	for _, v := range dependencyMap {
		dependencies = append(dependencies, v)
	}
	return StructureNode{
		name,
		properties,
		dependencies,
	}
}

func NewPropertyNode(name string, schema JsonSchema, isRequired bool, refString string) (PropertyNode) {
	return PropertyNode{
		name,
		NewTypeNodeWithSchema(schema, name, refString),
		isRequired,
	}
}

func NewTypeNodeWithSchema(schema JsonSchema, additionalKey string, refString string) (TypeNode) {
	if IsPrimitiveSchemaType(schema.Type) {
		return NewSpecifiedTypeNode(schema.Type)
	} else if schema.Type == SchemaTypeArray {
		// NOTE: not support multiple item types
		item := schema.GetItemList()[0]
		return NewArrayTypeNode(NewTypeNodeWithSchema(item, additionalKey, item.Ref))
	} else if schema.Type == SchemaTypeObject {
		typ := additionalKey
		if schema.Id != "" {
			typ = schema.Id
		} else if refString != "" {
			refs := strings.Split(refString, "/")
			typ = refs[len(refs) - 1]
		}
		return NewObjectTypeNode(NewSpecifiedTypeNode(typ))
	} else {
		panic("undefined type")
	}
}

func NewSpecifiedTypeNode(typ string) (TypeNode) {
	return TypeNode{ typ, nil }
}

func NewArrayTypeNode(containType TypeNode) (TypeNode) {
	return TypeNode{ string(SchemaTypeArray), &containType }
}

func NewObjectTypeNode(containType TypeNode) (TypeNode) {
	return TypeNode{ string(SchemaTypeObject), &containType }
}

