package main

import (
	"encoding/json"
	"reflect"
)

type JsonSchema struct {
	Schema string `json:"$schema"`
	Type string `json:"type"`
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

func (schema *JsonSchema) GetItemList() ([]JsonSchema) {
	if schema.Items == nil {
		return nil
	}
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

