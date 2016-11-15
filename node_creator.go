package main

type NodeCreator interface {
	CreateStructureNode(name string, bundle JsonSchemaBundle) (StructureNode, error)
	CreatePropertyNode(name string, bundle JsonSchemaBundle, isRequired bool) (PropertyNode, error)
	CreateTypeNode(bundle JsonSchemaBundle, additionalKey string) (TypeNode, error)
	ExportTo(exporter Exporter) error
}
