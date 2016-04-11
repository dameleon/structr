package main

import (
	"github.com/xeipuuv/gojsonreference"
)

type JsonSchemaBundler interface {
	GetRootSchema() (JsonSchema)
	GetSchema(string) (JsonSchema)
}

type bundler struct {
	rootRef *gojsonreference.JsonReference
	loader JsonSchemaLoader
}

func NewJsonSchemaBundler(loader JsonSchemaLoader, rootRef *gojsonreference.JsonReference) (JsonSchemaBundler) {
	b := &bundler{
		loader,
		rootRef,
	}
	return b
}

func (b *bundler) GetRootSchema() (JsonSchema) {
	return b.loader.Load(b.rootRef)
}

func (b *bundler) GetSchema(refString string) (JsonSchema) {
	return b.loader.Load(NewRelativeJsonReference(b.rootRef, NewJsonReference(refString)))
}

