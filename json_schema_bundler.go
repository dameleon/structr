package main

import (
	"fmt"
	"github.com/xeipuuv/gojsonreference"
)

type JsonSchemaBundler interface {
	AddJsonSchema(paths ...string) error
	GetBundles() map[string]JsonSchemaBundle
	GetBundle(ref gojsonreference.JsonReference) (JsonSchemaBundle, bool)
	GetReferredBundleWalk(bundle JsonSchemaBundle) (JsonSchemaBundle, error)
}

type bundler struct {
	loader  JsonSchemaLoader
	bundles map[string]JsonSchemaBundle
}

func NewJsonSchemaBundler(loader JsonSchemaLoader) JsonSchemaBundler {
	b := &bundler{
		loader:  loader,
		bundles: make(map[string]JsonSchemaBundle),
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
		bdl := NewJsonSchemaBundle(&ref, schema, false)
		if err := b.registerReferredBundleWalk(bdl); err != nil {
			return err
		}
		b.registerNewJsonSchemaBundle(bdl)
	}
	return nil
}

func (b *bundler) registerReferredBundleWalk(bundle JsonSchemaBundle) error {
	for _, r := range bundle.Schema.GetRefList() {
		ref, err := bundle.GetRelativeJsonReference(r)
		if err != nil {
			return err
		}
		schema, err := b.loader.Load(ref)
		if err != nil {
			return err
		}
		bdl := NewJsonSchemaBundle(&ref, schema, true)
		b.registerNewJsonSchemaBundle(bdl)
		b.registerReferredBundleWalk(bdl)
	}
	return nil
}

func (b *bundler) registerNewJsonSchemaBundle(bdl JsonSchemaBundle) {
	if _, ok := b.bundles[bdl.Ref.String()]; ok {
		return
	}
	b.bundles[bdl.Ref.String()] = bdl
}

func (b *bundler) GetBundles() map[string]JsonSchemaBundle {
	return b.bundles
}

func (b *bundler) GetBundle(ref gojsonreference.JsonReference) (JsonSchemaBundle, bool) {
	res, ok := b.bundles[ref.String()]
	return res, ok
}

func (b *bundler) GetReferredBundleWalk(bundle JsonSchemaBundle) (JsonSchemaBundle, error) {
	if !bundle.Schema.HasReference() {
		return bundle, nil
	}
	ref, err := bundle.GetRelativeJsonReference(bundle.Schema.Ref)
	if err != nil {
		return bundle, err
	}
	refBundle, ok := b.GetBundle(ref)
	if !ok {
		return refBundle, fmt.Errorf("undefined bundle. REF: %s", ref.String())
	}
	return b.GetReferredBundleWalk(refBundle)
}
