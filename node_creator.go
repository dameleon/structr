package main

import "errors"

type NodeCreator interface {
	CreateStructureNode(bundle Bundle) (StructureNode, error)
	CreatePropertyNode(name string, bundle Bundle, isRequired bool) (PropertyNode, error)
	CreateTypeNode(bundle Bundle, additionalKey string) (TypeNode, error)
}

func NewJsonSchemaNodeCreator(context Context, bundler JsonSchemaBundler) (NodeCreator) {
	return &jsonSchemaNodeCreator{ context, bundler }
}

type jsonSchemaNodeCreator struct {
	context Context
	bundler JsonSchemaBundler
}

func (creator *jsonSchemaNodeCreator) CreateStructureNode(rootBundle Bundle) (StructureNode, error) {
	rootSchema := rootBundle.Schema
	if rootSchema.Type != JsonSchemaTypeObject {
		return StructureNode{}, errors.New("root schema must be object type")
	}
	properties := []PropertyNode{}
	childrenMap := make(map[string]StructureNode)
	for key, schema := range rootSchema.Properties {
		var bdl Bundle
		if schema.HasReference() {
			// if current schema designated reference, specify referred bundle to create property
			ref, err := rootBundle.GetRelativeJsonReference(schema.Ref)
			if err != nil {
				return StructureNode{}, err
			}
			bdl = creator.bundler.GetBundle(ref)
		} else {
			bdl = rootBundle.CreateChild(schema)
		}
		// create property
		prop, err := creator.CreatePropertyNode(key, bdl, rootSchema.IsRequired(key))
		if err != nil {
			return StructureNode{}, err
		}
		properties = append(properties, prop)
		// create children
		if innerType := prop.Type.InnerType; innerType != nil && !bdl.IsReferred {
			name := innerType.EntityName()
			schema := schema
			switch schema.Type {
			case JsonSchemaTypeArray:
				schema = schema.GetItemList()[0]
				fallthrough
			case JsonSchemaTypeObject:
				if _, ok := childrenMap[name]; !ok {
					child, err := creator.CreateStructureNode(bdl.CreateChild(schema))
					if err != nil {
						return StructureNode{}, err
					}
					childrenMap[name] = child
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
	}, nil
}

func (creator *jsonSchemaNodeCreator) CreatePropertyNode(name string, bundle Bundle, isRequired bool) (PropertyNode, error) {
	typeNode, err := creator.CreateTypeNode(bundle, name)
	if err != nil {
		return PropertyNode{}, err
	}
	return PropertyNode{ name, typeNode, isRequired }, nil
}

func (creator *jsonSchemaNodeCreator) CreateTypeNode(bdl Bundle, additionalKey string) (TypeNode, error) {
	schema := bdl.Schema
	switch {
	case schema.Type.IsPrimitiveSchemaType():
		return newSpecifiedTypeNode(schema.Type.String()), nil
	case schema.Type == JsonSchemaTypeArray:
		// TODO: not support multiple item types
		childSchema := schema.GetItemList()[0]
		var innerBundle Bundle
		if childSchema.HasReference() {
			ref, err := bdl.GetRelativeJsonReference(childSchema.Ref)
			if err != nil {
				return TypeNode{}, err
			}
			innerBundle = creator.bundler.GetBundle(ref)
		} else {
			innerBundle = bdl.CreateChild(childSchema)
		}
		// create inner type recursive
		innerNode, err := creator.CreateTypeNode(innerBundle, additionalKey)
		if err != nil {
			return TypeNode{}, err
		}
		return newArrayTypeNode(innerNode), nil
	case schema.Type == JsonSchemaTypeObject:
		var typ string
		if bdl.IsReferred {
			typ = bdl.GetName()
		} else {
			typ = additionalKey
		}
		return newObjectTypeNode(newSpecifiedTypeNode(typ)), nil
	default:
		panic("undefined type")
	}
}

func newSpecifiedTypeNode(typ string) (TypeNode) {
	return TypeNode{ typ, nil }
}

func newArrayTypeNode(containType TypeNode) (TypeNode) {
	return TypeNode{ JsonSchemaTypeArray, &containType }
}

func newObjectTypeNode(containType TypeNode) (TypeNode) {
	return TypeNode{ JsonSchemaTypeObject, &containType }
}

