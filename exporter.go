package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Exporter interface {
	Export(node StructureNode) error
}

func NewExporter(b Bedrock) Exporter {
	if b.OutputsFiles() {
		return &fileExporter{b}
	}
	return &stdoutExporter{b}
}

type stdoutExporter struct {
	bedrock Bedrock
}

func (e *stdoutExporter) Export(node StructureNode) error {
	conf := e.bedrock.Config
	generator, err := NewStructGenerator(conf.StructureTemplate, conf.ChildStructuresNesting, conf.TypeTranslateMap)
	if err != nil {
		return err
	}
	str, err := generator.Generate(node)
	if err != nil {
		return err
	}
	_, err = os.Stdout.WriteString(str)
	return err
}

type fileExporter struct {
	bedrock Bedrock
}

func (e *fileExporter) Export(node StructureNode) error {
	conf := e.bedrock.Config
	generator, err := NewStructGenerator(conf.StructureTemplate, conf.ChildStructuresNesting, conf.TypeTranslateMap)
	if err != nil {
		return err
	}
	str, err := generator.Generate(node)
	if err != nil {
		return err
	}
	if err := e.mkdirIfNeeded(); err != nil {
		return err
	}
	filename, err := e.getFileName(node)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join(e.bedrock.OutputDirPath, filename), []byte(str), os.ModePerm)
}

func (e *fileExporter) mkdirIfNeeded() error {
	info, err := os.Stat(e.bedrock.OutputDirPath)
	if info != nil && info.IsDir() {
		return nil
	}
	if os.IsNotExist(err) {
		return os.MkdirAll(e.bedrock.OutputDirPath, os.ModePerm)
	}
	return err

}

func (e *fileExporter) getFileName(node StructureNode) (string, error) {
	tmpl, err := NewTemplate(node.Name).Parse(e.bedrock.Config.OutputFilename)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, node); err != nil {
		return "", err
	}
	return buf.String(), nil
}
