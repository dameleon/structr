package main

import (
	"testing"
	"github.com/dameleon/structr/fixtures"
)

func TestJsonSchemaNodeCreator_CreateStructureNode_WithBasicJsonSchema(t *testing.T) {
	bundler := NewJsonSchemaBundler(NewJsonSchemaLoader())
	bundler.AddJsonSchema(fixtures.BasicJsonSchemaRef.GetUrl().Path)
	creator := NewJsonSchemaNodeCreator(bundler)
	bundle, ok := bundler.GetBundle(fixtures.BasicJsonSchemaRef)
	if ok != true {
		t.Fatal("cannot get root bundle")
	}
	node, err := creator.CreateStructureNode("Basic", bundle)
	if err != nil {
		t.Fatal(err)
	}
	if node.Name != "Basic" {
		t.Errorf("node name should be 'Basic'. NAME: %s", node.Name)
	}
	if node.Parent != nil {
		t.Error("parent should be nil")
	}
	if len(node.Properties) != 7 {
		t.Errorf("properties count should be 7. LEN: %d", len(node.Properties))
	}
	matcher := newPropertiesMatcher(t, node.Properties)
	for _, typ := range JsonSchemaTypes {
		matcher.has(typ.String()+"Type", typ.String())
	}
	if len(node.Children) != 2 {
		t.Errorf("children count should be 2. LEN: %d", len(node.Children))
	}
	// children
	childrenMap := make(map[string]StructureNode)
	for _, child := range node.Children {
		childrenMap[child.Name] = child
	}
	// objectType
	{
		matcher := newPropertiesMatcher(t, childrenMap["objectType"].Properties)
		matcher.has("foo", "string")
		matcher.has("bar", "integer")
		matcher.has("baz", "number")
		matcher.has("nested", "object", "nested")
		// nested object type
		{
			nestedChildrenMap := make(map[string]StructureNode)
			for _, child := range childrenMap["objectType"].Children {
				nestedChildrenMap[child.Name] = child
			}
			matcher := newPropertiesMatcher(t, nestedChildrenMap["nested"].Properties)
			matcher.has("foo", "string")
			matcher.has("bar", "integer")
			matcher.has("baz", "number")
		}
	}
	// arrayType
	{
		matcher := newPropertiesMatcher(t, childrenMap["arrayType"].Properties)
		matcher.has("foo", "string")
		matcher.has("bar", "integer")
		matcher.has("baz", "number")
		matcher.has("nested", "array", "string")
	}
}

func TestJsonSchemaNodeCreator_CreateStructureNode_WithInternalReferenceSchema(t *testing.T) {
	bundler := NewJsonSchemaBundler(NewJsonSchemaLoader())
	bundler.AddJsonSchema(fixtures.InternalReferenceSchemaRef.GetUrl().Path)
	creator := NewJsonSchemaNodeCreator(bundler)
	bundle, ok := bundler.GetBundle(fixtures.InternalReferenceSchemaRef)
	if ok != true {
		t.Fatal("cannot get root bundle")
	}
	node, err := creator.CreateStructureNode("InternalReferenceSchema", bundle)
	if err != nil {
		t.Fatal(err)
	}
	if len(node.Children) != 0 {
		t.Errorf("children count should be 0. LEN: %d", len(node.Children))
	}
	if len(bundler.GetBundles()) != 3 {
		t.Errorf("bundle count should be 3. LEN: %d", len(bundler.GetBundles()))
	}
	// refObjectType
	{
		ref, _ := bundle.GetRelativeJsonReference("#/definitions/refObjectType")
		bundle, _ := bundler.GetBundle(ref)
		node, err := creator.CreateStructureNode("refObjectType", bundle)
		if err != nil {
			t.Fatal(err)
		}
		if len(node.Properties) != 4 {
			t.Errorf("properties count should be 4. LEN: %d", len(node.Properties))
		}
		matcher := newPropertiesMatcher(t, node.Properties)
		matcher.has("foo", "string")
		matcher.has("bar", "integer")
		matcher.has("baz", "number")
		matcher.has("nested", "object", "nested")
		if len(node.Children) != 1 {
			t.Errorf("children count should be 1. LEN: %d", len(node.Children))
		}
	}
	// refArrayType
	{
		ref, _ := bundle.GetRelativeJsonReference("#/definitions/refArrayType")
		bundle, _ := bundler.GetBundle(ref)
		node, err := creator.CreateStructureNode("refArrayType", bundle)
		if err != nil {
			t.Fatal(err)
		}
		if len(node.Properties) != 4 {
			t.Errorf("properties count should be 4. LEN: %d", len(node.Properties))
		}
		matcher := newPropertiesMatcher(t, node.Properties)
		matcher.has("foo", "string")
		matcher.has("bar", "integer")
		matcher.has("baz", "number")
		matcher.has("nested", "array", "object", "refObjectType")
		if len(node.Children) != 0 {
			t.Errorf("children count should be 0. LEN: %d", len(node.Children))
		}
	}
}

func TestJsonSchemaNodeCreator_CreateStructureNode_WithExternalReferenceSchema(t *testing.T) {
	bundler := NewJsonSchemaBundler(NewJsonSchemaLoader())
	bundler.AddJsonSchema(fixtures.ExternalReferenceSchemaRef.GetUrl().Path)
	creator := NewJsonSchemaNodeCreator(bundler)
	bundle, ok := bundler.GetBundle(fixtures.ExternalReferenceSchemaRef)
	if ok != true {
		t.Fatal("cannot get root bundle")
	}
	node, err := creator.CreateStructureNode("ExternalReferenceSchema", bundle)
	if err != nil {
		t.Fatal(err)
	}
	if len(node.Children) != 0 {
		t.Errorf("children count should be 0. LEN: %d", len(node.Children))
	}
	if len(bundler.GetBundles()) != 8 {
		t.Errorf("bundle count should be 8. LEN: %d", len(bundler.GetBundles()))
	}
	// refObjectType
	{
		ref, _ := bundle.GetRelativeJsonReference("./type/objects.json")
		bundle, _ := bundler.GetBundle(ref)
		node, err := creator.CreateStructureNode("Objects", bundle)
		if err != nil {
			t.Fatal(err)
		}
		if len(node.Properties) != 3 {
			t.Errorf("properties count should be 3. LEN: %d", len(node.Properties))
		}
		matcher := newPropertiesMatcher(t, node.Properties)
		matcher.has("foo", "string")
		matcher.has("bar", "integer")
		matcher.has("baz", "number")
	}
}

func newPropertiesMatcher(t *testing.T, properties []PropertyNode) *propertiesMatcher {
	props := make(map[string]PropertyNode)
	for _, p := range properties {
		props[p.Name] = p
	}
	return &propertiesMatcher{t, props}
}

type propertiesMatcher struct {
	t     *testing.T
	props map[string]PropertyNode
}

func (m *propertiesMatcher) has(name string, types ...string) {
	prop, ok := m.props[name]
	if !ok {
		m.t.Errorf("undefined property. NAME: %s", name)
	}
	typeNode := &prop.Type
	for _, typ := range types {
		if typeNode == nil {
			m.t.Fatal("type node is nil")
		}
		if typeNode.Name != typ {
			m.t.Errorf("mismatched type name. NAME: %s", typeNode.Name)
		}
		typeNode = typeNode.InnerType
	}
}
