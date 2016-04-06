package main

import (
	"github.com/xeipuuv/gojsonreference"
	"log"
)

func NewJsonReference(path string) (gojsonreference.JsonReference) {
	ref, err := gojsonreference.NewJsonReference(path)
	if err != nil {
		log.Fatal(err)
	}
	return ref
}

func NewRelativeJsonReference(base gojsonreference.JsonReference, ref gojsonreference.JsonReference) (gojsonreference.JsonReference) {
	if ref.IsCanonical() {
		return ref
	}
	return NewJsonReference(base.GetUrl().ResolveReference(ref.GetUrl()).String())
}

