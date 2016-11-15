package main

type NodeCreator interface {
	Create(exporter Exporter) error
}
