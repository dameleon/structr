package main

import (
	"testing"
	"encoding/json"
	"github.com/xeipuuv/gojsonpointer"
	"github.com/dameleon/structr/fixtures"
)

func TestJsonSchema_GetItemList_ByObject(t *testing.T) {
	var schema JsonSchema
	itemsByObject := `{ "items": { "type": "string" } }`
	json.Unmarshal([]byte(itemsByObject), &schema)
	items := schema.GetItemList()
	if len(items) != 1 {
		t.Error("mismatched item length. LEN: %d", len(items))
	}
	if items[0].Type != JsonSchemaTypeString {
		t.Error("mismatched item type. TYPE: %s", items[0].Type)
	}
}

func TestJsonSchema_GetItemList_ByArray(t *testing.T) {
	var schema JsonSchema
	itemsByArray := `{ "items": [{ "type": "string" }, { "type": "integer" }] }`
	json.Unmarshal([]byte(itemsByArray), &schema)
	items := schema.GetItemList()
	if len(items) != 2 {
		t.Error("mismatched item length. LEN: %d", len(items))
	}
	if items[0].Type != JsonSchemaTypeString {
		t.Error("mismatched item type. TYPE: %s", items[0].Type)
	}
	if items[1].Type != JsonSchemaTypeInteger {
		t.Error("mismatched item type. TYPE: %s", items[0].Type)
	}
}

func TestJsonSchema_IsRequired(t *testing.T) {
	var schema JsonSchema
	json.Unmarshal([]byte(`{ "required": ["hoge", "fuga"] }`), &schema)

	if schema.IsRequired("hoge") != true {
		t.Error("key 'hoge' should be required")
	}
	if schema.IsRequired("piyo") != false {
		t.Error("key 'piyo' should be not required")
	}
}

func TestJsonSchema_GetRefSchema(t *testing.T) {
	var schema JsonSchema
	json.Unmarshal([]byte(`{ "definitions": { "hoge": { "type": "string" } } }`), &schema)
	p, _ := gojsonpointer.NewJsonPointer("/definitions/hoge")
	refSchema := schema.GetRefSchema(&p)
	if refSchema.Type != JsonSchemaTypeString {
		t.Error("ref schema should be type of string")
	}
}

func TestJsonSchema_GetRefList(t *testing.T) {
	schema, _ := NewJsonSchemaLoader().Load(fixtures.InternalReferenceSchemaRef)
	refList := schema.GetRefList()
	if len(refList) != 2 {
		t.Errorf("ref list count should be 2. LEN: %d", len(refList))
	}
}

func TestJsonSchema_HasReference(t *testing.T) {
	{
		var schema JsonSchema
		json.Unmarshal([]byte(`{ "$ref": "#/definitions/hoge" }`), &schema)
		if schema.HasReference() != true {
			t.Error("has reference should be true")
		}
	}
	{
		var schema JsonSchema
		json.Unmarshal([]byte(`{ "$ref": "" }`), &schema)
		if schema.HasReference() != false {
			t.Error("has reference should be false")
		}
	}
}

func TestJsonSchema_HasStructure(t *testing.T) {
	// object
	{
		var schema JsonSchema
		json.Unmarshal([]byte(`{ "type": "object" }`), &schema)
		if schema.HasStructure() != true {
			t.Error("has structure should be true")
		}

	}
	// object with array
	{
		var schema JsonSchema
		json.Unmarshal([]byte(`{ "type": "array", "items": { "type": "object" } }`), &schema)
		if schema.HasStructure() != true {
			t.Error("has structure should be true")
		}
	}
	// other
	{
		var schema JsonSchema
		json.Unmarshal([]byte(`{ "type": "string" }`), &schema)
		if schema.HasStructure() != false {
			t.Error("has structure should be false")
		}
	}
}
