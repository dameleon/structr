package main

import (
	"encoding/json"
	"github.com/xeipuuv/gojsonreference"
	"io/ioutil"
)

type JsonSchemaLoader interface {
	Load(ref gojsonreference.JsonReference) (JsonSchema, error)
}

func NewJsonSchemaLoader() JsonSchemaLoader {
	return &jsonSchemaLoader{make(map[string]JsonSchema)}
}

type jsonSchemaLoader struct {
	pool map[string]JsonSchema
}

func (l *jsonSchemaLoader) Load(ref gojsonreference.JsonReference) (JsonSchema, error) {
	path := ref.GetUrl().RequestURI()
	if _, ok := l.pool[path]; !ok {
		var schema JsonSchema
		var err error
		if ref.HasFullFilePath || ref.HasFileScheme {
			schema, err = l.loadFromFile(ref)
		} else {
			schema, err = l.loadFromRemote(ref)
		}
		if err != nil {
			return schema, err
		}
		l.pool[path] = schema
	}
	if ref.GetUrl().Fragment != "" {
		return l.pool[path].GetRefSchema(ref.GetPointer()), nil
	}
	return l.pool[path], nil
}

func (l *jsonSchemaLoader) loadFromRemote(ref gojsonreference.JsonReference) (JsonSchema, error) {
	// TODO
	return JsonSchema{}, nil
}

func (l *jsonSchemaLoader) loadFromFile(ref gojsonreference.JsonReference) (JsonSchema, error) {
	var s JsonSchema
	j, err := ioutil.ReadFile(ref.GetUrl().Path)
	if err != nil {
		return s, err
	}
	if err := json.Unmarshal(j, &s); err != nil {
		return s, err
	}
	return s, nil
}
