package main

import (
	"github.com/xeipuuv/gojsonreference"
	"testing"
)

var ref, _ = gojsonreference.NewJsonReference("/foo/bar/baz.json")

func TestNewJsonSchemaBundle(t *testing.T) {
	noIdentifiedJsonSchemaBundle := NewJsonSchemaBundle(&ref, JsonSchema{Title: "NoIdentified"}, false)
	if noIdentifiedJsonSchemaBundle.Name != "baz" {
		t.Errorf("bundle name should be 'baz'. NAME: %S", noIdentifiedJsonSchemaBundle.Name)
	}
	if noIdentifiedJsonSchemaBundle.Ref != &ref {
		t.Error("mismatched json reference pointer")
	}
	if noIdentifiedJsonSchemaBundle.Schema.Title != "NoIdentified" {
		t.Error("json schema title should be 'NoIdentified'. TITLE: %s", noIdentifiedJsonSchemaBundle.Schema.Title)
	}
	if noIdentifiedJsonSchemaBundle.IsReferred != false {
		t.Error("is referred should be false")
	}
	if noIdentifiedJsonSchemaBundle.HasParent != false {
		t.Error("has parent should be false")
	}

	identifiedJsonSchemaBundle := NewJsonSchemaBundle(&ref, JsonSchema{Id: "Identified"}, false)
	if identifiedJsonSchemaBundle.Name != "Identified" {
		t.Errorf("bundle name should be 'Identified'. NAME: %S", identifiedJsonSchemaBundle.Name)
	}
	if identifiedJsonSchemaBundle.Schema.Id != "Identified" {
		t.Error("json schema Id should be 'Identified'")
	}
}

func TestNewNamedChildJsonSchemaBundle(t *testing.T) {
	rootJsonSchemaBundle := NewJsonSchemaBundle(&ref, JsonSchema{Id: "Root"}, false)
	bundle := NewNamedChildJsonSchemaBundle(rootJsonSchemaBundle, "Child", JsonSchema{Id: "ChildSchema"})
	if bundle.Name != "Child" {
		t.Errorf("bundle name should be 'Child'. NAME: %s", bundle.Name)
	}
	if bundle.Schema.Id != "ChildSchema" {
		t.Errorf("json schema id should be 'Child'. ID: %s", bundle.Schema.Id)
	}
	if bundle.IsReferred != rootJsonSchemaBundle.IsReferred {
		t.Error("is referred should be same to parent value")
	}
	if bundle.HasParent != true {
		t.Error("has parent should be true")
	}
}

func TestNewChildJsonSchemaBundle(t *testing.T) {
	rootJsonSchemaBundle := NewJsonSchemaBundle(&ref, JsonSchema{Id: "Root"}, false)
	bundle := NewChildJsonSchemaBundle(rootJsonSchemaBundle, JsonSchema{Id: "Child"})
	if bundle.Name != "Root" {
		t.Errorf("bundle name should be 'Root'. NAME: %s", bundle.Name)
	}
	if bundle.Schema.Id != "Child" {
		t.Errorf("json schema id should be 'Child'. ID: %s", bundle.Schema.Id)
	}
	if bundle.IsReferred != rootJsonSchemaBundle.IsReferred {
		t.Error("is referred should be same to parent value")
	}
	if bundle.HasParent != true {
		t.Error("has parent should be true")
	}
}

func TestJsonSchemaBundle_GetRelativeJsonReference(t *testing.T) {
	rootJsonSchemaBundle := NewJsonSchemaBundle(&ref, JsonSchema{Id: "Root"}, false)
	// relative path
	ref, err := rootJsonSchemaBundle.GetRelativeJsonReference("./hoge/fuga/piyo.json")
	if err != nil {
		t.Fatal(err)
	}
	if ref.GetUrl().Path != "/foo/bar/hoge/fuga/piyo.json" {
		t.Errorf("mismatched ref url. URL: %s", ref.GetUrl().Path)
	}
	// absolute path
	ref, err = rootJsonSchemaBundle.GetRelativeJsonReference("/hoge/fuga/piyo.json")
	if ref.GetUrl().Path != "/hoge/fuga/piyo.json" {
		t.Errorf("mismatched ref url. URL: %s", ref.GetUrl().Path)
	}
}

func TestJsonSchemaBundle_IsSameRef(t *testing.T) {
	rootJsonSchemaBundle := NewJsonSchemaBundle(&ref, JsonSchema{Id: "Root"}, false)
	childJsonSchemaBundle := NewChildJsonSchemaBundle(rootJsonSchemaBundle, JsonSchema{Id: "Child"})

	if rootJsonSchemaBundle.IsSameRef(childJsonSchemaBundle) != true {
		t.Error("is same ref should be true")
	}
	otherRef, _ := gojsonreference.NewJsonReference("/other/bundle.json")
	otherJsonSchemaBundle := NewJsonSchemaBundle(&otherRef, JsonSchema{}, false)
	if rootJsonSchemaBundle.IsSameRef(otherJsonSchemaBundle) != false {
		t.Error("is same ref should be false")
	}
}
