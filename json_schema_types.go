package main

type JsonSchemaType string

const (
	JsonSchemaTypeNull    JsonSchemaType = "null"
	JsonSchemaTypeBoolean                = "boolean"
	JsonSchemaTypeString                 = "string"
	JsonSchemaTypeInteger                = "integer"
	JsonSchemaTypeNumber                 = "number"
	JsonSchemaTypeObject                 = "object"
	JsonSchemaTypeArray                  = "array"
)

var JsonSchemaTypes = [...]JsonSchemaType{
	JsonSchemaTypeNull,
	JsonSchemaTypeBoolean,
	JsonSchemaTypeString,
	JsonSchemaTypeInteger,
	JsonSchemaTypeNumber,
	JsonSchemaTypeObject,
	JsonSchemaTypeArray,
}

func (t JsonSchemaType) IsPrimitiveSchemaType() bool {
	switch t {
	case JsonSchemaTypeNull, JsonSchemaTypeBoolean, JsonSchemaTypeString, JsonSchemaTypeInteger, JsonSchemaTypeNumber:
		return true
	}
	return false
}

func (t JsonSchemaType) String() string {
	return string(t)
}
