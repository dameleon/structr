package main

type NodeCreator interface {
	CreateStructureNode(bundle Bundle) (StructureNode)
	CreatePropertyNode(name string, bundle Bundle, isRequired bool) (PropertyNode)
	CreateTypeNode(bundle Bundle, additionalKey string) (TypeNode)
}

func NewJsonSchemaNodeCreator(context Context, bundler JsonSchemaBundler) (NodeCreator) {
	return jsonSchemaNodeCreator{ context, bundler }
}

type jsonSchemaNodeCreator struct {
	context Context
	bundler JsonSchemaBundler
}

func (creator jsonSchemaNodeCreator) CreateStructureNode(rootBundle Bundle) (StructureNode) {
	rootSchema := rootBundle.Schema
	if rootSchema.Type != SchemaTypeObject {
		panic("root schema must be type of object")
	}
	properties := []PropertyNode{}
	childrenMap := make(map[string]StructureNode)
	for key, schema := range rootSchema.Properties {
		var bdl Bundle
		if schema.HasReference() {
			// if current schema designated reference, specify referred bundle to create property
			bdl = creator.bundler.GetBundle(rootBundle.GetRelativeJsonReference(schema.Ref))
		} else {
			bdl = rootBundle.CreateChild(schema)
		}
		// create property
		prop := creator.CreatePropertyNode(key, bdl, rootSchema.IsRequired(key))
		properties = append(properties, prop)
		// create children
		if innerType := prop.Type.InnerType; innerType != nil && !bdl.IsReferred {
			name := innerType.EntityName()
			schema := schema
			if schema.Type == SchemaTypeArray {
				schema = schema.GetItemList()[0]
			}
			if schema.Type == SchemaTypeObject {
				if _, ok := childrenMap[name]; !ok {
					childrenMap[name] = creator.CreateStructureNode(bdl.CreateChild(schema))
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

func (creator jsonSchemaNodeCreator) CreatePropertyNode(name string, bundle Bundle, isRequired bool) (PropertyNode) {
	return PropertyNode{
		name,
		creator.CreateTypeNode(bundle, name),
		isRequired,
	}
}

func (creator jsonSchemaNodeCreator) CreateTypeNode(bdl Bundle, additionalKey string) (TypeNode) {
	schema := bdl.Schema
	if IsPrimitiveSchemaType(schema.Type) {
		return creator.newSpecifiedTypeNode(schema.Type)
	} else if schema.Type == SchemaTypeArray {
		// TODO: not support multiple item types
		childSchema := schema.GetItemList()[0]
		var innerBundle Bundle
		if childSchema.HasReference() {
			innerBundle = creator.bundler.GetBundle(bdl.GetRelativeJsonReference(childSchema.Ref))
		} else {
			innerBundle = bdl.CreateChild(childSchema)
		}
		// create inner type recursive
		return creator.newArrayTypeNode(creator.CreateTypeNode(innerBundle, additionalKey))
	} else if schema.Type == SchemaTypeObject {
		var typ string
		if bdl.IsReferred {
			typ = bdl.GetName()
		} else {
			typ = additionalKey
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

