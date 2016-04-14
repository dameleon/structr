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
	conf := e.context.Config
	generator, err := NewStructGenerator(conf.StructureTemplate, conf.TypeTranslateMap)
	if err != nil {
		return err
	}
	if err := generator.Generate(os.Stdout, node); err != nil {
		return err
	}
	return nil
}

type fileExporter struct {
	context Context
}

func (e *fileExporter) Export(node StructureNode) (error) {
	conf := e.context.Config
	generator, err := NewStructGenerator(conf.StructureTemplate, conf.TypeTranslateMap)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := generator.Generate(&buf, node); err != nil {
		return err
	}
	if err := e.mkdirIfNeeded(); err != nil {
		return err
	}
	filename, err := e.getFileName(node)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join(e.context.OutputDirPath, filename), buf.Bytes(), os.ModePerm)
}

func (e *fileExporter) mkdirIfNeeded() (error) {
	info, err := os.Stat(e.context.OutputDirPath)
	if info != nil && info.IsDir() {
		return nil
	}
	if os.IsNotExist(err) {
		return os.MkdirAll(e.context.OutputDirPath, os.ModePerm)
	}
	return err

}

func (e *fileExporter) getFileName(node StructureNode) (string, error) {
	tmpl, err := NewTemplate(node.Name).Parse(e.context.Config.OutputFilename)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, node); err != nil {
		return "", err
	}
	return buf.String(), nil
}
