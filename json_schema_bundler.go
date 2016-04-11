package main

import (
	"github.com/xeipuuv/gojsonreference"
)

type JsonSchemaBundler interface {
	GetRootSchema() (JsonSchema)
	GetSchema(string) (JsonSchema)
}

type bundler struct {
	loader JsonSchemaLoader
	rootRef gojsonreference.JsonReference
}

func NewJsonSchemaBundler(loader JsonSchemaLoader, rootRef gojsonreference.JsonReference) (JsonSchemaBundler) {
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
	ref := NewRelativeJsonReference(b.rootRef, NewJsonReference(refString))
	schema := b.loader.Load(ref)
	if schema.Ref != "" {
		return b.GetSchema(schema.Ref)
	}
	return schema
}
