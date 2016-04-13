package main

import (
	"github.com/xeipuuv/gojsonreference"
)

type JsonSchemaBundler interface {
	AddJsonSchema(paths ...string) (error)
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

func (b *bundler) AddJsonSchema(paths ...string) (error) {
	for _, path := range paths {
		ref, err := gojsonreference.NewJsonReference(path)
		if err != nil {
			return err
		}
		schema, err := b.loader.Load(ref)
		if err != nil {
			return err
		}
		bdl := Bundle{ ref, schema, false }
		for _, r := range bdl.Schema.GetRefList() {
			ref, err := bdl.GetRelativeJsonReference(r)
			if err != nil {
				return err
			}
			schema, err := b.loader.Load(ref)
			if err != nil {
				return err
			}
			b.registerNewBundle(Bundle{ ref, schema, true })
		}
		b.registerNewBundle(bdl)
	}
	return nil
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
