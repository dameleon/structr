package main

import (
	"github.com/xeipuuv/gojsonreference"
	"path/filepath"
	"strings"
)

type Bundle struct {
	Name       string
	Ref        *gojsonreference.JsonReference
	Schema     JsonSchema
	IsReferred bool
	HasParent  bool
}

func NewBundle(ref *gojsonreference.JsonReference, schema JsonSchema, isReferred bool) Bundle {
	var name string
	if schema.Id != "" {
		name = schema.Id
	} else {
		basename := filepath.Base(ref.GetUrl().String())
		name = strings.Replace(basename, filepath.Ext(basename), "", 1)
	}
	return Bundle{name, ref, schema, isReferred, false}
}

func NewNamedChildBundle(bundle Bundle, name string, schema JsonSchema) Bundle {
	return Bundle{name, bundle.Ref, schema, bundle.IsReferred, true}
}

func NewChildBundle(bundle Bundle, schema JsonSchema) Bundle {
	return Bundle{bundle.Name, bundle.Ref, schema, bundle.IsReferred, true}
}

func (b Bundle) GetRelativeJsonReference(path string) (gojsonreference.JsonReference, error) {
	ref, err := gojsonreference.NewJsonReference(path)
	if err != nil {
		return ref, err
	}
	return NewRelativeJsonReference(b.Ref, &ref)
}

func (b Bundle) IsSameRef(bundle Bundle) bool {
	return b.Ref.GetUrl() == bundle.Ref.GetUrl()
}
