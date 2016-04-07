package main

import (
	"github.com/xeipuuv/gojsonreference"
	"path"
	"strings"
)

type JsonSchemaBundler interface {
	Bundle()
	GetRootBundle() (*bundle)
	GetBundle(string) (*bundle)
}

type bundler struct {
	rootRef *gojsonreference.JsonReference
	bundles map[string]JsonSchema
}

func NewJsonSchemaBundler(rootReference *gojsonreference.JsonReference) (JsonSchemaBundler) {
	b := bundler{ rootReference, make(map[string]JsonSchema) }
	b.registerBundle(b.rootRef.String(), b.rootRef)
	return &b
}

func (b *bundler) GetRootBundle() (*bundle) {
	return b.GetBundle(b.rootRef.String())
}

func (b *bundler) Bundle() {
	b.bundleRecursive(b.GetRootBundle().schema)
}

func (b *bundler) bundleRecursive(schema *JsonSchema) {
	if schema.Ref != "" {
		b.registerBundle(schema.Ref, &NewRelativeJsonReference(b.rootRef, NewJsonReference(schema.Ref)))
	}
	// object type
	for _, v := range schema.Properties {
		b.bundleRecursive(&v)
	}
	for _, d := range schema.Dependencies {
		b.bundleRecursive(&d)
	}
	// array type
	for _, i := range schema.GetItemList() {
		b.bundleRecursive(&i)
	}
}

func (b *bundler) registerBundle(path string, ref *gojsonreference.JsonReference) {
	b.bundles[path] = bundle {
		ref,
		&GetJsonSchemaLoader().Load(ref),
	}
}

func (b *bundler) GetBundle(path string) (*bundle) {
	return &b.bundles[path]
}

type bundle struct {
	ref *gojsonreference.JsonReference
	schema *JsonSchema
}

func (b *bundle) GetBundleName() (string) {
	if b.schema.Id {
		return b.schema.Id
	}
	// filename + pointer
	p := b.ref.GetUrl().Path
	n := []string{
		strings.Replace(path.Base(p), path.Ext(p), "", 1),
		strings.Split(b.ref.GetPointer().String(), "/"),
	}
	return
}


