package main

import "encoding/json"

type StringDefinition struct {
	name   string
}

type IntegerDefinition struct {
	name    string
	minimum json.Number
	maximum json.Number
}

type NumberDefinition struct {
	name string
	minimum json.Number
	maximum json.Number
	exclusiveMaximum bool
}

type BooleanDefinition struct {
	name string
}

type ObjectDefinition struct {
	name string
	properties interface{}
	required []string
}

type ArrayDefinition struct {
	name string
	items []interface{}
}

type NullDefinition struct {
	name string
}