package main

import (
	"fmt"
)

func NewJsonSchemaNodeCreator(bundler JsonSchemaBundler) NodeCreator {
	return &jsonSchemaNodeCreator{bundler}
}

type jsonSchemaNodeCreator struct {
	bundler JsonSchemaBundler
}

func (creator *jsonSchemaNodeCreator) createStructureNode(name string, rootBundle JsonSchemaBundle) (StructureNode, error) {
	rootBundle, err := creator.bundler.GetReferredBundleWalk(rootBundle)
	if err != nil {
		return StructureNode{}, err
	}
	rootSchema := rootBundle.Schema
	if rootSchema.Type == JsonSchemaTypeArray {
		return creator.createStructureNode(name, NewChildJsonSchemaBundle(rootBundle, rootSchema.GetItemList()[0]))
	} else if rootSchema.Type != JsonSchemaTypeObject {
		return StructureNode{}, fmt.Errorf("root schema must be object type. TYPE: %s", rootSchema.Type.String())
	}
	node := StructureNode{
		Name:       name,
		Properties: []PropertyNode{},
		Children:   []StructureNode{},
	}
	childrenMap := make(map[string]StructureNode)
	for key, schema := range rootSchema.Properties {
		bdl, err := creator.bundler.GetReferredBundleWalk(NewNamedChildJsonSchemaBundle(rootBundle, key, schema))
		if err != nil {
			return node, err
		}
		// create property
		prop, err := creator.createPropertyNode(key, bdl, rootSchema.IsRequired(key))
		if err != nil {
			return StructureNode{}, err
		}
		node.Properties = append(node.Properties, prop)
		// create children
		if innerType := prop.Type.InnerType; innerType != nil {
			name := innerType.EntityName()
			bdl := bdl
			for bdl.Schema.Type == JsonSchemaTypeArray {
				bdl, err = creator.bundler.GetReferredBundleWalk(NewChildJsonSchemaBundle(bdl, bdl.Schema.GetItemList()[0]))
				if err != nil {
					return node, err
				}
			}
			_, ok := childrenMap[name]
			if bdl.Schema.HasStructure() && rootBundle.IsSameRef(bdl) && bdl.HasParent && !ok {
				child, err := creator.createStructureNode(name, bdl)
				if err != nil {
					return StructureNode{}, err
				}
				childrenMap[name] = child
			}
		}
	}
	for _, v := range childrenMap {
		v.Parent = &node
		node.Children = append(node.Children, v)
	}
	return node, nil
}

func (creator *jsonSchemaNodeCreator) createPropertyNode(name string, bundle JsonSchemaBundle, isRequired bool) (PropertyNode, error) {
	typeNode, err := creator.createTypeNode(bundle, name)
	if err != nil {
		return PropertyNode{}, err
	}
	return PropertyNode{
		Name:       name,
		Type:       typeNode,
		IsRequired: isRequired,
	}, nil
}

func (creator *jsonSchemaNodeCreator) createTypeNode(bdl JsonSchemaBundle, additionalKey string) (TypeNode, error) {
	schema := bdl.Schema
	switch {
	case schema.Type.IsPrimitiveSchemaType():
		return newSpecifiedTypeNode(schema.Type.String()), nil
	case schema.Type == JsonSchemaTypeArray:
		// TODO: not support multiple item types
		innerBundle, err := creator.bundler.GetReferredBundleWalk(NewChildJsonSchemaBundle(bdl, bdl.Schema.GetItemList()[0]))
		if err != nil {
			return TypeNode{}, err
		}
		// create inner type recursive
		innerNode, err := creator.createTypeNode(innerBundle, additionalKey)
		if err != nil {
			return TypeNode{}, err
		}
		return newArrayTypeNode(innerNode), nil
	case schema.Type == JsonSchemaTypeObject:
		return newObjectTypeNode(newSpecifiedTypeNode(bdl.Name)), nil
	default:
		panic("undefined type")
	}
}

func newSpecifiedTypeNode(typ string) TypeNode {
	return TypeNode{typ, nil}
}

func newArrayTypeNode(containType TypeNode) TypeNode {
	return TypeNode{JsonSchemaTypeArray, &containType}
}

func newObjectTypeNode(containType TypeNode) TypeNode {
	return TypeNode{JsonSchemaTypeObject, &containType}
}

func (creator *jsonSchemaNodeCreator) Create(exporter Exporter) error {
	for _, b := range creator.bundler.GetBundles() {
		if !b.Schema.HasStructure() {
			continue
		}
		structure, err := creator.createStructureNode(b.Name, b)
		if err != nil {
			return err
		}
		if err := exporter.Export(structure); err != nil {
			return err
		}
	}
	return nil
}
