package main

import (
	"encoding/json"
	"reflect"
	"strings"
	"github.com/xeipuuv/gojsonpointer"
)

type JsonSchema struct {
	Schema string `json:"$schema"`
	Type JsonSchemaType `json:"type"`
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Default interface{} `json:"default"`
	Enum []interface{} `json:"enum"`
	Definitions map[string]JsonSchema `json:"definitions"`
	Ref string `json:"$ref"`

	// string
	MinLength json.Number `json:"minLength"`
	MaxLength json.Number `json:"maxLength"`
	Pattern string `json:"pattern"`
	Format string `json:"format"`

	// numeric
	MultipleOf json.Number `json:"multipleOf"`
	Minimum json.Number `json:"minimum"`
	Maximum json.Number `json:"maximum"`
	ExclusiveMaximum bool `json:"exclusiveMaximum"`

	// object
	Properties map[string]JsonSchema `json:"properties"`
	AdditionalProperties interface{} `json:"additionalProperties"`
	Required []string `json:"required"`
	MinProperties json.Number `json:"minProperties"`
	MaxProperties json.Number `json:"maxProperties"`
	Dependencies map[string]JsonSchema `json:"dependencies"`

	// array
	Items interface{} `json:"items"`
	AdditionalItems bool `json:"additionalItems"`
	MinItems json.Number `json:"minItems"`
	MaxItems json.Number `json:"maxItems"`
	UniqueItems bool `json:"uniqueItems"`

	// combining
	AllOf []JsonSchema `json:"allOf"`
	AnyOf []JsonSchema `json:"anyOf"`
	OneOf []JsonSchema `json:"oneOf"`
	Not []JsonSchema `json:"not"`
}

func (schema JsonSchema) GetItemList() ([]JsonSchema) {
	if schema.Items == nil {
		return nil
	}
	// schema.items defined type of Object or array
	j, _ := json.Marshal(schema.Items)
	var res []JsonSchema
	switch reflect.ValueOf(schema.Items).Kind() {
	case reflect.Array, reflect.Slice:
		json.Unmarshal(j, &res)
	case reflect.Map:
		var s JsonSchema
		json.Unmarshal(j, &s)
		res = []JsonSchema{ s }
	}
	return res
}

func (schema JsonSchema) IsRequired(key string) (bool) {
	if schema.Required == nil {
		return false
	}
	for _, r := range schema.Required {
		if key == r {
			return true
		}
	}
	return false
}

func (schema JsonSchema) GetRefSchema(pointer *gojsonpointer.JsonPointer) (JsonSchema) {
	path := strings.Split(pointer.String(), "/")
	s := schema
	i := 0
	for i < len(path) {
		p := path[i]
		if p != "" && p == "definitions" {
			i++
			s = s.Definitions[path[i]]
		}
		i++
	}
	return s
}

func (schema JsonSchema) GetRefList() ([]string) {
	res := []string{}
	switch schema.Type {
	case JsonSchemaTypeObject:
		for _, s := range schema.Properties {
			res = append(res, s.GetRefList()...)
		}
	case JsonSchemaTypeArray:
		for _, s := range schema.GetItemList() {
			res = append(res, s.GetRefList()...)
		}
	default:
		if schema.Ref != "" {
			res = append(res, schema.Ref)
		}
	}
	return res
}

func (schema JsonSchema) HasReference() (bool) {
	return schema.Ref != ""
}