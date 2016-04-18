package main

import (
	"testing"
	"github.com/xeipuuv/gojsonreference"
)

var ref, _ = gojsonreference.NewJsonReference("/foo/bar/baz.json")

func TestNewBundle(t *testing.T) {
	noIdentifiedBundle := NewBundle(&ref, JsonSchema{ Title: "NoIdentified" }, false)
	if noIdentifiedBundle.Name != "baz" {
		t.Errorf("bundle name should be 'baz'. NAME: %S", noIdentifiedBundle.Name)
	}
	if noIdentifiedBundle.Ref != &ref {
		t.Error("mismatched json reference pointer")
	}
	if noIdentifiedBundle.Schema.Title != "NoIdentified" {
		t.Error("json schema title should be 'NoIdentified'. TITLE: %s", noIdentifiedBundle.Schema.Title)
	}
	if noIdentifiedBundle.IsReferred != false {
		t.Error("is referred should be false")
	}
	if noIdentifiedBundle.HasParent != false {
		t.Error("has parent should be false")
	}

	identifiedBundle := NewBundle(&ref, JsonSchema{ Id: "Identified" }, false)
	if identifiedBundle.Name != "Identified" {
		t.Errorf("bundle name should be 'Identified'. NAME: %S", identifiedBundle.Name)
	}
	if identifiedBundle.Schema.Id != "Identified" {
		t.Error("json schema Id should be 'Identified'")
	}
}

func TestNewNamedChildBundle(t *testing.T) {
	rootBundle := NewBundle(&ref, JsonSchema{ Id: "Root" }, false)
	bundle := NewNamedChildBundle(rootBundle, "Child", JsonSchema{ Id: "ChildSchema" })
	if bundle.Name != "Child" {
		t.Errorf("bundle name should be 'Child'. NAME: %s", bundle.Name)
	}
	if bundle.Schema.Id != "ChildSchema" {
		t.Errorf("json schema id should be 'Child'. ID: %s", bundle.Schema.Id)
	}
	if bundle.IsReferred != rootBundle.IsReferred {
		t.Error("is referred should be same to parent value")
	}
	if bundle.HasParent != true {
		t.Error("has parent should be true")
	}
}

func TestNewChildBundle(t *testing.T) {
	rootBundle := NewBundle(&ref, JsonSchema{ Id: "Root" }, false)
	bundle := NewChildBundle(rootBundle, JsonSchema{ Id: "Child" })
	if bundle.Name != "Root" {
		t.Errorf("bundle name should be 'Root'. NAME: %s", bundle.Name)
	}
	if bundle.Schema.Id != "Child" {
		t.Errorf("json schema id should be 'Child'. ID: %s", bundle.Schema.Id)
	}
	if bundle.IsReferred != rootBundle.IsReferred {
		t.Error("is referred should be same to parent value")
	}
	if bundle.HasParent != true {
		t.Error("has parent should be true")
	}
}

func TestBundle_GetRelativeJsonReference(t *testing.T) {
	rootBundle := NewBundle(&ref, JsonSchema{ Id: "Root" }, false)
	// relative path
	ref, err := rootBundle.GetRelativeJsonReference("./hoge/fuga/piyo.json")
	if err != nil {
		t.Fatal(err)
	}
	if ref.GetUrl().Path != "/foo/bar/hoge/fuga/piyo.json" {
		t.Errorf("mismatched ref url. URL: %s", ref.GetUrl().Path)
	}
	// absolute path
	ref, err = rootBundle.GetRelativeJsonReference("/hoge/fuga/piyo.json")
	if ref.GetUrl().Path != "/hoge/fuga/piyo.json" {
		t.Errorf("mismatched ref url. URL: %s", ref.GetUrl().Path)
	}
}

func TestBundle_IsSameRef(t *testing.T) {
	rootBundle := NewBundle(&ref, JsonSchema{ Id: "Root" }, false)
	childBundle := NewChildBundle(rootBundle, JsonSchema{ Id: "Child" })

	if rootBundle.IsSameRef(childBundle) != true {
		t.Error("is same ref should be true")
	}
	otherRef, _ := gojsonreference.NewJsonReference("/other/bundle.json")
	otherBundle := NewBundle(&otherRef, JsonSchema{}, false)
	if rootBundle.IsSameRef(otherBundle) != false {
		t.Error("is same ref should be false")
	}
}