package main

import (
	"github.com/xeipuuv/gojsonreference"
	"path/filepath"
	"strings"
)

type JsonSchemaBundler interface {
	AddJsonSchema(path string)
	GetBundles() (map[string]bundle)
	GetBundle(ref gojsonreference.JsonReference) (bundle)
}

type bundler struct {
	loader JsonSchemaLoader
	bundles map[string]bundle
}

func NewJsonSchemaBundler(loader JsonSchemaLoader) (JsonSchemaBundler) {
	b := &bundler{
		loader,
		make(map[string]bundle),
	}
	return b
}

func (b *bundler) AddJsonSchema(path string) {
	ref := NewJsonReference(path)
	bdl := bundle { ref, b.loader.Load(ref), false }
	for _, r := range bdl.schema.GetRefList() {
		ref := bdl.GetRelativeJsonReference(r)
		b.registerNewBundle(bundle{ ref, b.loader.Load(ref), true })
	}
	b.registerNewBundle(bdl)
}

func (b *bundler) registerNewBundle(bdl bundle) {
	if _, ok := b.bundles[bdl.ref.String()]; ok {
		return
	}
	b.bundles[bdl.ref.String()] = bdl
}

func (b *bundler) GetBundles() (map[string]bundle) {
	return b.bundles
}

func (b *bundler) GetBundle(ref gojsonreference.JsonReference) (bundle) {
	if bdl, ok := b.bundles[ref.String()]; ok {
		return bdl
	}
	panic("undefined bundle")
}

type bundle struct {
	ref gojsonreference.JsonReference
	schema JsonSchema
	isReferred bool
}

func (b bundle) GetRelativeJsonReference(path string) (gojsonreference.JsonReference) {
	return NewRelativeJsonReference(b.ref, NewJsonReference(path))
}

func (b bundle) GetName() (string) {
	if b.schema.Id != "" {
		return b.schema.Id
	}
	basename := filepath.Base(b.ref.GetUrl().String())
	return strings.Replace(basename, filepath.Ext(basename), "", 1)
}

func (b bundle) CreateChild(schema JsonSchema) (bundle) {
	return bundle{ b.ref, schema, b.isReferred }
}