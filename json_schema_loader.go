package main

import (
	"github.com/xeipuuv/gojsonreference"
	"io/ioutil"
	"log"
	"encoding/json"
)

type jsonSchemaLoader struct {
	pool map[string]JsonSchema
}

func (l* jsonSchemaLoader) Load(ref gojsonreference.JsonReference) (JsonSchema) {
	path := getUniquePath(ref)
	if _, ok := l.pool[path]; !ok {
		var schema JsonSchema
		if ref.HasFullFilePath || ref.HasFileScheme {
			schema = l.loadFromFile(ref)
		} else {
			schema = l.loadFromRemote(ref)
		}
		l.pool[path] = schema
	}
	if ref.GetUrl().Fragment != "" {
		return l.pool[path].GetRefSchema(ref.GetPointer())
	}
	return l.pool[path]
}

func (l jsonSchemaLoader) loadFromRemote(ref gojsonreference.JsonReference) (JsonSchema) {
	// TODO
	return JsonSchema{}
}

func (l jsonSchemaLoader) loadFromFile(ref gojsonreference.JsonReference) (JsonSchema) {
	j, err := ioutil.ReadFile(ref.GetUrl().Path)
	if err != nil {
		log.Fatal(err)
	}
	var s JsonSchema
	if e := json.Unmarshal(j, &s); e != nil {
		log.Fatal(e)
	}
	return s
}

var instance *jsonSchemaLoader

func GetJsonSchemaLoader() *jsonSchemaLoader {
	if instance == nil {
		instance = &jsonSchemaLoader{ make(map[string]JsonSchema) }
	}
	return instance
}

func getUniquePath(ref gojsonreference.JsonReference) (string) {
	url := ref.GetUrl()
	return url.RequestURI()
}