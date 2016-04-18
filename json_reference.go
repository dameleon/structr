package main

import (
	"github.com/xeipuuv/gojsonreference"
)

func NewRelativeJsonReference(base *gojsonreference.JsonReference, ref *gojsonreference.JsonReference) (gojsonreference.JsonReference, error) {
	if ref.IsCanonical() {
		return *ref, nil
	}
	return gojsonreference.NewJsonReference(base.GetUrl().ResolveReference(ref.GetUrl()).String())
}
