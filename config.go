package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	TypeTranslateMap    map[string]string `yaml:"type_translate_map"`
	NestedDependency    bool `yaml:"nested_dependency"`
	ArrayTypeDefinition string `yaml:"array_type_definition"`
	OutputFilename      string `yaml:"output_filename"`
	StructureTemplate   string `yaml:"structure_template"`
}

func NewConfig(filepath string) (Config, error) {
	config := Config{}
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		return config, err
	}
	if err := yaml.Unmarshal(f, &config); err != nil {
		return config, err
	}
	return config, nil
}
