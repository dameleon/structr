package main

type NodeCreator interface {
	CreateStructureNode(bundle bundle) (StructureNode)
	CreatePropertyNode(name string, bundle bundle, isRequired bool) (PropertyNode)
	CreateTypeNode(bundle bundle, additionalKey string) (TypeNode)
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
		var bdl bundle
		if schema.HasReference() {
			bdl = creator.bundler.GetBundle(rootBundle.GetRelativeJsonReference(schema.Ref))
		} else {
			bdl = rootBundle.CreateChild(schema)
		}
		// create property
		prop := creator.CreatePropertyNode(key, bdl, rootSchema.IsRequired(key))
		properties = append(properties, prop)
		// create children
		if innerType := prop.Type.InnerType; innerType != nil && !bdl.isReferred {
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

func (creator jsonSchemaNodeCreator) CreatePropertyNode(name string, bundle bundle, isRequired bool) (PropertyNode) {
	return PropertyNode{
		name,
		creator.CreateTypeNode(bundle, name),
		isRequired,
	}
}

func (creator jsonSchemaNodeCreator) CreateTypeNode(bdl bundle, additionalKey string) (TypeNode) {
	schema := bdl.schema
	if IsPrimitiveSchemaType(schema.Type) {
		return creator.newSpecifiedTypeNode(schema.Type)
	} else if schema.Type == SchemaTypeArray {
		// NOTE: not support multiple item types
		childSchema := schema.GetItemList()[0]
		var innerBundle bundle
		if childSchema.HasReference() {
			innerBundle = creator.bundler.GetBundle(bdl.GetRelativeJsonReference(childSchema.Ref))
		} else {
			innerBundle = bdl.CreateChild(childSchema)
		}
		return creator.newArrayTypeNode(creator.CreateTypeNode(innerBundle, additionalKey))
	} else if schema.Type == SchemaTypeObject {
		var typ string
		if bdl.isReferred {
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

