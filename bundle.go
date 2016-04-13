package main

import (
	"github.com/xeipuuv/gojsonreference"
	"path/filepath"
	"strings"
)

type Bundle struct {
	Ref        gojsonreference.JsonReference
	Schema     JsonSchema
	IsReferred bool
}

func (b Bundle) GetRelativeJsonReference(path string) (gojsonreference.JsonReference, error) {
	ref, err := gojsonreference.NewJsonReference(path)
	if err != nil {
		return ref, err
	}
	return NewRelativeJsonReference(b.Ref, ref)
}

func (b Bundle) GetName() (string) {
	if b.Schema.Id != "" {
		return b.Schema.Id
	}
	basename := filepath.Base(b.Ref.GetUrl().String())
	return strings.Replace(basename, filepath.Ext(basename), "", 1)
}

func (b Bundle) CreateChild(schema JsonSchema) (Bundle) {
	return Bundle{ b.Ref, schema, b.IsReferred }
}
