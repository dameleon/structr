package main

import (
	"github.com/xeipuuv/gojsonreference"
)

type JsonSchemaBundler interface {
	AddJsonSchema(paths ...string) error
	GetBundles() map[string]Bundle
	GetBundle(ref gojsonreference.JsonReference) Bundle
	GetReferredBundleWalk(bundle Bundle) (Bundle, error)
}

type bundler struct {
	loader  JsonSchemaLoader
	bundles map[string]Bundle
}

func NewJsonSchemaBundler(loader JsonSchemaLoader) JsonSchemaBundler {
	b := &bundler{
		loader:  loader,
		bundles: make(map[string]Bundle),
	}
	return b
}

func (b *bundler) AddJsonSchema(paths ...string) error {
	for _, path := range paths {
		ref, err := gojsonreference.NewJsonReference(path)
		if err != nil {
			return err
		}
		schema, err := b.loader.Load(ref)
		if err != nil {
			return err
		}
		bdl := NewBundle(&ref, schema, false)
		for _, r := range bdl.Schema.GetRefList() {
			ref, err := bdl.GetRelativeJsonReference(r)
			if err != nil {
				return err
			}
			schema, err := b.loader.Load(ref)
			if err != nil {
				return err
			}
			b.registerNewBundle(NewBundle(&ref, schema, true))
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

func (b *bundler) GetBundles() map[string]Bundle {
	return b.bundles
}

func (b *bundler) GetBundle(ref gojsonreference.JsonReference) Bundle {
	if bdl, ok := b.bundles[ref.String()]; ok {
		return bdl
	}
	panic("undefined bundle")
}

func (b *bundler) GetReferredBundleWalk(bundle Bundle) (Bundle, error) {
	if !bundle.Schema.HasReference() {
		return bundle, nil
	}
	ref, err := bundle.GetRelativeJsonReference(bundle.Schema.Ref)
	if err != nil {
		return bundle, err
	}
	refBundle := b.GetBundle(ref)
	return b.GetReferredBundleWalk(refBundle)
}
