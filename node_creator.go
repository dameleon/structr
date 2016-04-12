package main

import "strings"

type NodeCreator interface {
	CreateStructureNode(name string, schema JsonSchema) (StructureNode)
	CreatePropertyNode(name string, schema JsonSchema, isRequired bool, refString string) (PropertyNode)
	CreateTypeNode(schema JsonSchema, additionalKey string, refString string) (TypeNode)
}

func NewJsonSchemaNodeCreator(context Context, bundler JsonSchemaBundler) (NodeCreator) {
	return jsonSchemaNodeCreator{ context, bundler }
}

type jsonSchemaNodeCreator struct {
	context Context
	bundler JsonSchemaBundler
}

func (creator jsonSchemaNodeCreator) CreateStructureNode(name string, rootSchema JsonSchema) (StructureNode) {
	if rootSchema.Type != SchemaTypeObject {
		panic("root schema must be type of object")
	}
	properties := []PropertyNode{}
	childrenMap := make(map[string]StructureNode)
	for key, schema := range rootSchema.Properties {
		refString := ""
		if schema.Ref != "" {
			refString = schema.Ref
			schema = creator.bundler.GetSchema(schema.Ref)
		}
		// create property
		prop := creator.CreatePropertyNode(key, schema, rootSchema.IsRequired(key), refString)
		properties = append(properties, prop)
		// create children
		if innerType := prop.Type.InnerType; innerType != nil && refString == "" {
			name := innerType.EntityName()
			s := schema
			if s.Type == SchemaTypeArray {
				s = schema.GetItemList()[0]
			}
			if s.Type == SchemaTypeObject {
				if _, ok := childrenMap[name]; !ok {
					childrenMap[name] = creator.CreateStructureNode(name, s)
				}
			}
		}
	}
	children := []StructureNode{}
	for _, v := range childrenMap {
		children = append(children, v)
	}
	return StructureNode{
		name,
		properties,
		children,
	}
}

func (creator jsonSchemaNodeCreator) CreatePropertyNode(name string, schema JsonSchema, isRequired bool, refString string) (PropertyNode) {
	return PropertyNode{
		name,
		creator.CreateTypeNode(schema, name, refString),
		isRequired,
	}
}

func (creator jsonSchemaNodeCreator) CreateTypeNode(schema JsonSchema, additionalKey string, refString string) (TypeNode) {
	if IsPrimitiveSchemaType(schema.Type) {
		return creator.newSpecifiedTypeNode(schema.Type)
	} else if schema.Type == SchemaTypeArray {
		// NOTE: not support multiple item types
		item := schema.GetItemList()[0]
		return creator.newArrayTypeNode(creator.CreateTypeNode(item, additionalKey, item.Ref))
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

