package main

import (
	"testing"
	"github.com/xeipuuv/gojsonreference"
	"errors"
)

type stubJsonSchemaLoader struct {
	cond bool
}

func (l stubJsonSchemaLoader) Load(ref gojsonreference.JsonReference) (JsonSchema, error) {
	if l.cond {
		switch ref.GetUrl().String() {
		case "/base":
			return JsonSchema{ Id: "BaseSchema", Ref: "/referred" }, nil
		case "/referred":
			return JsonSchema{ Id: "ReferredSchema", Ref: "/deep_referred" }, nil
		case "/deep_referred":
			return JsonSchema{ Id: "DeepReferredSchema" }, nil
		}
	}
	return JsonSchema{ Id: "ErrorJsonSchema" }, errors.New("cannot load json schema")
}

func newSucceedBundler() (JsonSchemaBundler) {
	return NewJsonSchemaBundler(stubJsonSchemaLoader{true})
}

func newFailureBundler() (JsonSchemaBundler) {
	return NewJsonSchemaBundler(stubJsonSchemaLoader{false})
}

func TestJsonSchemaBundler_AddJsonSchema(t *testing.T) {
	sb := newSucceedBundler()
	fb := newFailureBundler()

	// succeed
	if err := sb.AddJsonSchema("/base"); err != nil {
		t.Fatal("succeed bundler should be load json schema")
	}
	ref, _ := gojsonreference.NewJsonReference("/referred")
	bundle, _ := sb.GetBundle(ref)
	if bundle.Ref.GetUrl().Path != ref.GetUrl().Path {
		t.Errorf("referred json schema has invalid url URL: %s", bundle.Ref.GetUrl())
	}
	// failure
	if err := fb.AddJsonSchema("/base"); err == nil {
		t.Fatal("failure bundler should be not load json schema")
	}
}

func TestJsonSchemaBundler_GetBundles(t *testing.T) {
	sb := newSucceedBundler()
	if err := sb.AddJsonSchema("/base"); err != nil {
		t.Fatal("succeed bundler should be load json schema")
	}
	if len(sb.GetBundles()) != 3 {
		t.Errorf("mismatched bundle count. COUNT: %d", len(sb.GetBundles()))
	}
}

func TestJsonSchemaBundler_GetBundle(t *testing.T) {
	sb := newSucceedBundler()
	if err := sb.AddJsonSchema("/base"); err != nil {
		t.Fatal("succeed bundler should be load json schema")
	}
	ref, _ := gojsonreference.NewJsonReference("/base")
	bundle, ok := sb.GetBundle(ref)
	if !ok {
		t.Fatal("cannot get bundle")
	} else if bundle.Schema.Id != "BaseSchema" {
		t.Error("get bundle retures invalid schema ID: %s", bundle.Schema.Id)
	}
}

func TestJsonSchemaBundler_GetReferredBundleWalk(t *testing.T) {
	sb := newSucceedBundler()
	if err := sb.AddJsonSchema("/base"); err != nil {
		t.Fatal("succeed bundler should be load json schema")
	}
	ref, _ := gojsonreference.NewJsonReference("/base")
	bundle, _ := sb.GetBundle(ref)
	refBundle, err := sb.GetReferredBundleWalk(bundle)
	if err != nil {
		t.Fatal("cannot get referred bundle")
	} else if refBundle.Schema.Id != "DeepReferredSchema" {
		t.Error("get bundle retures invalid schema ID: %s", bundle.Schema.Id)
	}
}