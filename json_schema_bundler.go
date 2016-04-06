package main

import (
	"github.com/xeipuuv/gojsonreference"
)

type JsonSchemaBundler interface {
	Bundle()
	GetRootSchema() (JsonSchema)
	GetSchema(string) (JsonSchema)
}

type bundler struct {
	rootRef gojsonreference.JsonReference
	bundles map[string]JsonSchema
}

func NewJsonSchemaBundler(rootReference gojsonreference.JsonReference) (JsonSchemaBundler) {
	b := bundler{ rootReference, make(map[string]JsonSchema) }
	b.registerBundle(b.rootRef.String(), b.rootRef)
	return &b
}

func (b *bundler) GetRootSchema() (JsonSchema) {
	return b.GetSchema(b.rootRef.String())
}

func (b *bundler) Bundle() {
	b.bundleRecursive(b.GetRootSchema())
}

func (b *bundler) bundleRecursive(schema JsonSchema) {
	if schema.Ref != "" {
		b.registerBundle(schema.Ref, NewRelativeJsonReference(b.rootRef, NewJsonReference(schema.Ref)))
	}
	// object type
	for _, v := range schema.Properties {
		b.bundleRecursive(v)
	}
	for _, d := range schema.Dependencies {
		b.bundleRecursive(d)
	}
	// array type
	for _, i := range schema.GetItemList() {
		b.bundleRecursive(i)
	}
}

func (b *bundler) registerBundle(path string, ref gojsonreference.JsonReference) {
	b.bundles[path] = GetJsonSchemaLoader().Load(ref)
}

func (b *bundler) GetSchema(path string) (JsonSchema) {
	return b.bundles[path]
}

