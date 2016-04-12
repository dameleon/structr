package main

import (
	"strings"
)

type NodeCreator interface {
	CreateStructureNode(bundle bundle) (StructureNode)
	CreatePropertyNode(name string, bundle bundle, isRequired bool, refString string) (PropertyNode)
	CreateTypeNode(bundle bundle, additionalKey string, refString string) (TypeNode)
}

func NewJsonSchemaNodeCreator(context Context, bundler JsonSchemaBundler) (NodeCreator) {
	return jsonSchemaNodeCreator{ context, bundler }
}

type jsonSchemaNodeCreator struct {
	context Context
	bundler JsonSchemaBundler
}

func (creator jsonSchemaNodeCreator) CreateStructureNode(rootBundle bundle) (StructureNode) {
	rootSchema := rootBundle.schema
	if rootSchema.Type != SchemaTypeObject {
		panic("root schema must be type of object")
	}
	properties := []PropertyNode{}
	childrenMap := make(map[string]StructureNode)
	for key, schema := range rootSchema.Properties {
		refString := schema.Ref
		if refString != "" {
			schema = creator.bundler.GetReferredBundle(rootBundle.GetRelativeJsonReference(refString)).schema
		}
		bundle := rootBundle.CreateChild(schema)
		// create property
		prop := creator.CreatePropertyNode(key, bundle, rootSchema.IsRequired(key), refString)
		properties = append(properties, prop)
		// create children
		if innerType := prop.Type.InnerType; innerType != nil && refString == "" {
			name := innerType.EntityName()
			schema := schema
			if schema.Type == SchemaTypeArray {
				schema = schema.GetItemList()[0]
			}
			if schema.Type == SchemaTypeObject {
				if _, ok := childrenMap[name]; !ok {
					childrenMap[name] = creator.CreateStructureNode(bundle.CreateChild(schema))
				}
			}
		}
	}
	children := []StructureNode{}
	for _, v := range childrenMap {
		children = append(children, v)
	}
	return StructureNode{
		rootBundle.GetName(),
		properties,
		children,
	}
}

func (creator jsonSchemaNodeCreator) CreatePropertyNode(name string, bundle bundle, isRequired bool, refString string) (PropertyNode) {
	return PropertyNode{
		name,
		creator.CreateTypeNode(bundle, name, refString),
		isRequired,
	}
}

func (creator jsonSchemaNodeCreator) CreateTypeNode(bundle bundle, additionalKey string, refString string) (TypeNode) {
	schema := bundle.schema
	if IsPrimitiveSchemaType(schema.Type) {
		return creator.newSpecifiedTypeNode(schema.Type)
	} else if schema.Type == SchemaTypeArray {
		// NOTE: not support multiple item types
		innerBundle := bundle.CreateChild(schema.GetItemList()[0])
		refString := innerBundle.schema.Ref
		if refString != "" {
			innerBundle = creator.bundler.GetReferredBundle(bundle.GetRelativeJsonReference(refString))
		}
		return creator.newArrayTypeNode(creator.CreateTypeNode(innerBundle, additionalKey, refString))
	} else if schema.Type == SchemaTypeObject {
		typ := additionalKey
		if schema.Id != "" {
			typ = schema.Id
		} else if refString != "" {
			refs := strings.Split(refString, "/")
			typ = refs[len(refs) - 1]
		}
		return creator.newObjectTypeNode(creator.newSpecifiedTypeNode(typ))
	} else {
		panic("undefined type")
	}
}

func (creator jsonSchemaNodeCreator) newSpecifiedTypeNode(typ string) (TypeNode) {
	return TypeNode{ typ, nil }
}

func (creator jsonSchemaNodeCreator) newArrayTypeNode(containType TypeNode) (TypeNode) {
	return TypeNode{ SchemaTypeArray, &containType }
}

func (creator jsonSchemaNodeCreator) newObjectTypeNode(containType TypeNode) (TypeNode) {
	return TypeNode{ SchemaTypeObject, &containType }
}

