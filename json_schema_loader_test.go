package main

import (
	"github.com/dameleon/structr/fixtures"
	"testing"
)

func TestJsonSchemaLoader_Load_FromFile(t *testing.T) {
	loader := NewJsonSchemaLoader()
	// exists file
	schema, err := loader.Load(fixtures.BasicJsonSchemaRef)
	if err != nil {
		t.Fatal(err)
	}
	if schema.Id != "Basic" {
		t.Errorf("Invalid schema loaded. ID: %s", schema.Id)
	}
	// not exists file
	schema, err = loader.Load(fixtures.NotExistsSchemaRef)
	if err == nil {
		t.Errorf("Should not the file exists. REF: %s", fixtures.NotExistsSchemaRef)
	}
}
