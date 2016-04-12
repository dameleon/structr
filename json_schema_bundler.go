package main

import (
	"github.com/xeipuuv/gojsonreference"
	"path/filepath"
	"strings"
)

type JsonSchemaBundler interface {
	AddJsonSchema(path string)
	GetReferredBundle(ref gojsonreference.JsonReference) (bundle)
	GetBundles() map[string]bundle
	GetReferredBundles() map[string]bundle
}

type bundler struct {
	loader JsonSchemaLoader
	bundles map[string]bundle
	referredBundles map[string]bundle
}

func NewJsonSchemaBundler(loader JsonSchemaLoader) (JsonSchemaBundler) {
	b := &bundler{
		loader,
		make(map[string]bundle),
		make(map[string]bundle),
	}
	return b
}

func (b *bundler) AddJsonSchema(path string) {
	ref := NewJsonReference(path)
	if _, ok := b.bundles[ref.String()]; ok {
		return
	}
	bandle := bundle { ref, b.loader.Load(ref) }
	for _, r := range bandle.schema.GetRefList() {
		b.addReferredJsonSchema(bandle.GetRelativeJsonReference(r))
	}
	b.bundles[ref.String()] = bandle
}

func (b *bundler) addReferredJsonSchema(ref gojsonreference.JsonReference) {
	if _, ok := b.referredBundles[ref.String()]; ok {
		return
	}
	b.referredBundles[ref.String()] = bundle{ ref, b.loader.Load(ref) }
}

func (b *bundler) GetReferredBundle(ref gojsonreference.JsonReference) (bundle) {
	if bundle, ok := b.referredBundles[ref.String()]; ok {
		return bundle
	}
	panic("undefined bundle")
}

func (b *bundler) GetBundles() ([]bundle) {
	return b.bundles
}

func (b *bundler) GetReferredBundles() ([]bundle) {
	return b.referredBundles
}

func NewBundle(ref gojsonreference.JsonReference, schema JsonSchema) (bundle) {
	return bundle { ref, schema }
}

type bundle struct {
	ref gojsonreference.JsonReference
	schema JsonSchema
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