package main

import (
	"bytes"
	"path/filepath"
	"os"
	"io/ioutil"
)

type Exporter interface {
	Export(node StructureNode) (error)
}

func NewExporter(context Context) (Exporter) {
	if context.OutputsFiles() {
		return &fileExporter{ context }
	}
	return &stdoutExporter{ context }
}

type stdoutExporter struct {
	context Context
}

func (e *stdoutExporter) Export(node StructureNode) (error) {
	tmpl := NewContextualTemplate(e.context, node.Name)
	if err := tmpl.Execute(os.Stdout, node); err != nil {
		return err
	}
	return nil
}

type fileExporter struct {
	context Context
}

func (e *fileExporter) Export(node StructureNode) (error) {
	tmpl := NewContextualTemplate(e.context, node.Name)
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, node); err != nil {
		return err
	}
	ioutil.WriteFile(filepath.Join(e.context.OutputDirPath, e.getFileName(node)), buf.Bytes(), os.ModePerm)
	return nil
}

func (e *fileExporter) getFileName(node StructureNode) (string) {
	tmpl := NewCommonTemplate(node.Name, e.context.Config.OutputFilename)
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, node); err != nil {
		panic(err)
	}
	return buf.String()
}
