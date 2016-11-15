package main

import (
	"github.com/xeipuuv/gojsonreference"
	"path/filepath"
	"strings"
)

type JsonSchemaBundle struct {
	Name       string
	Ref        *gojsonreference.JsonReference
	Schema     JsonSchema
	IsReferred bool
	HasParent  bool
}

func NewJsonSchemaBundle(ref *gojsonreference.JsonReference, schema JsonSchema, isReferred bool) JsonSchemaBundle {
	var name string
	if schema.Id != "" {
		name = schema.Id
	} else {
		basename := filepath.Base(ref.GetUrl().String())
		name = strings.Replace(basename, filepath.Ext(basename), "", 1)
	}
	return JsonSchemaBundle{name, ref, schema, isReferred, false}
}

func NewNamedChildJsonSchemaBundle(bundle JsonSchemaBundle, name string, schema JsonSchema) JsonSchemaBundle {
	return JsonSchemaBundle{name, bundle.Ref, schema, bundle.IsReferred, true}
}

func NewChildJsonSchemaBundle(bundle JsonSchemaBundle, schema JsonSchema) JsonSchemaBundle {
	return JsonSchemaBundle{bundle.Name, bundle.Ref, schema, bundle.IsReferred, true}
}

func (b JsonSchemaBundle) GetRelativeJsonReference(path string) (gojsonreference.JsonReference, error) {
	ref, err := gojsonreference.NewJsonReference(path)
	if err != nil {
		return ref, err
	}
	return NewRelativeJsonReference(b.Ref, &ref)
}

func (b JsonSchemaBundle) IsSameRef(bundle JsonSchemaBundle) bool {
	return b.Ref.GetUrl() == bundle.Ref.GetUrl()
}
