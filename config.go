package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	TypeTranslateMap       map[string]string `yaml:"type_translate_map"`
	OutputFilename         string            `yaml:"output_filename"`
	OutputDependencies     bool              `yaml:"output_dependencies"`
	StructureTemplate      string            `yaml:"structure_template"`
	ChildStructuresNesting string            `yaml:"child_structures_nesting"`
}

func NewConfig(filepath string) (Config, error) {
	var config Config
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		return config, err
	}
	if err := yaml.Unmarshal(f, &config); err != nil {
		return config, err
	}
	return config, nil
}
