package main

const (
	SchemaTypeNull = "null"
	SchemaTypeBoolean = "boolean"
	SchemaTypeString = "string"
	SchemaTypeInteger = "integer"
	SchemaTypeNumber = "number"
	SchemaTypeObject = "object"
	SchemaTypeArray = "array"
)

var SchemaTypes = [...]string{
	SchemaTypeNull,
	SchemaTypeBoolean,
	SchemaTypeString,
	SchemaTypeInteger,
	SchemaTypeNumber,
	SchemaTypeObject,
	SchemaTypeArray,
}

func IsPrimitiveSchemaType(target string) (bool) {
	switch target {
	case SchemaTypeNull, SchemaTypeBoolean, SchemaTypeString, SchemaTypeInteger, SchemaTypeNumber:
		return true
	}
	return false
}

