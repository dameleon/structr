package main

import (
	"github.com/xeipuuv/gojsonreference"
)

type JsonSchemaBundler interface {
	AddJsonSchema(paths ...string)
	GetBundles() (map[string]Bundle)
	GetBundle(ref gojsonreference.JsonReference) (Bundle)
}

type bundler struct {
	loader JsonSchemaLoader
	bundles map[string]Bundle
}

func NewJsonSchemaBundler(loader JsonSchemaLoader) (JsonSchemaBundler) {
	b := &bundler{
		loader,
		make(map[string]Bundle),
	}
	return b
}

func (b *bundler) AddJsonSchema(paths ...string) {
	for _, path := range paths {
		ref := NewJsonReference(path)
		bdl := Bundle{ ref, b.loader.Load(ref), false }
		for _, r := range bdl.Schema.GetRefList() {
			ref := bdl.GetRelativeJsonReference(r)
			b.registerNewBundle(Bundle{ ref, b.loader.Load(ref), true })
		}
		b.registerNewBundle(bdl)
	}
}

func (b *bundler) registerNewBundle(bdl Bundle) {
	if _, ok := b.bundles[bdl.Ref.String()]; ok {
		return
	}
	b.bundles[bdl.Ref.String()] = bdl
}

func (b *bundler) GetBundles() (map[string]Bundle) {
	return b.bundles
}

func (b *bundler) GetBundle(ref gojsonreference.JsonReference) (Bundle) {
	if bdl, ok := b.bundles[ref.String()]; ok {
		return bdl
	}
	panic("undefined bundle")
}
