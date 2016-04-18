package main

import (
	"testing"
	"github.com/xeipuuv/gojsonreference"
	"github.com/docker/machine/drivers/vmwarevsphere/errors"
)

type stubJsonSchemaLoader struct {
	cond bool
}

func (l stubJsonSchemaLoader) Load(ref gojsonreference.JsonReference) (JsonSchema, error) {
	if l.cond {
		switch ref.GetUrl() {
		case "base":
			return JsonSchema{ Id: "BaseSchema", Ref: "referred" }, nil
		case "referred":
			return JsonSchema{ Id: "ReferredSchema" }
		}
	}
	return JsonSchema{ Id: "ErrorJsonSchema" }, errors.New("load error")
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

	if err := sb.AddJsonSchema("hoge"); err != nil {

	}


}
