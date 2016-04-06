package main

import "log"

type SchemaType string

const (
	SchemaTypeNull SchemaType = "null"
	SchemaTypeBoolean = "boolean"
	SchemaTypeString = "string"
	SchemaTypeInteger = "integer"
	SchemaTypeNumber = "number"
	SchemaTypeObject = "object"
	SchemaTypeArray = "array"
)

var SchemaTypes = [...]SchemaType{
	SchemaTypeNull,
	SchemaTypeBoolean,
	SchemaTypeString,
	SchemaTypeInteger,
	SchemaTypeNumber,
	SchemaTypeObject,
	SchemaTypeArray,
}

func SchemaTypeFromString(str string) (SchemaType) {
	for _, v := range SchemaTypes {
		if str == string(v) {
			return v
		}
	}
	log.Fatal("undefined type")
}
