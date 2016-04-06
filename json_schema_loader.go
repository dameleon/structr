package main

import (
	"github.com/xeipuuv/gojsonreference"
	"io/ioutil"
	"log"
	"encoding/json"
)

type JsonSchemaLoader interface {
	Load() ([]byte, error)
}

func NewJsonSchemaLoader(ref gojsonreference.JsonReference) (JsonSchemaLoader, error) {
	// localfile...?
	if ref.HasFileScheme || ref.GetUrl().Scheme == "" {
		return LocalJsonSchemaLoader{ ref }
	}
	return RemoteJsonSchemaLoader{ ref }
}

type RemoteJsonSchemaLoader struct {
	ref gojsonreference.JsonReference
}

func (l *RemoteJsonSchemaLoader) Load() ([]byte, error) {
	// TBD
	return []byte{}, nil
}

type LocalJsonSchemaLoader struct {
	ref gojsonreference.JsonReference
}

func (l *LocalJsonSchemaLoader) loadJson() ([]byte) {
	file, err := ioutil.ReadFile(l.ref.GetUrl().Path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func (l *LocalJsonSchemaLoader) Load() (JsonSchema) {
	var s JsonSchema
	if e := json.Unmarshal(l.loadJson(), s); e != nil {
		log.Fatal(e)
	}
	return s
}