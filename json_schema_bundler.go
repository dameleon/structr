package main

import (
	"fmt"
	"github.com/xeipuuv/gojsonreference"
)

type JsonSchemaBundler interface {
	AddJsonSchema(paths ...string) error
	GetBundles() map[string]Bundle
	GetBundle(ref gojsonreference.JsonReference) (Bundle, bool)
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
		if err := b.registerReferredBundleWalk(bdl); err != nil {
			return err
		}
		b.registerNewBundle(bdl)
	}
	return nil
}

func (b *bundler) registerReferredBundleWalk(bundle Bundle) error {
	for _, r := range bundle.Schema.GetRefList() {
		ref, err := bundle.GetRelativeJsonReference(r)
		if err != nil {
			return err
		}
		schema, err := b.loader.Load(ref)
		if err != nil {
			return err
		}
		bdl := NewBundle(&ref, schema, true)
		b.registerNewBundle(bdl)
		b.registerReferredBundleWalk(bdl)
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

func (b *bundler) GetBundle(ref gojsonreference.JsonReference) (Bundle, bool) {
	res, ok := b.bundles[ref.String()]
	return res, ok
}

func (b *bundler) GetReferredBundleWalk(bundle Bundle) (Bundle, error) {
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
