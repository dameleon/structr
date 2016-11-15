package main

import (
	"errors"
	"github.com/codegangsta/cli"
	"os"
	"path/filepath"
)

type Bedrock struct {
	Config        Config
	OutputDirPath string
	Inputs        []string
	InputType     InputType
}

func NewBedrock(configFilePath string, inputType string, outDir string, args cli.Args) (Bedrock, error) {
	var b Bedrock
	wd, err := os.Getwd()
	if err != nil {
		return b, err
	}
	b.InputType = StringToInputMode(inputType)
	// create & register config
	if configFilePath == "" {
		return b, errors.New("config flag must be specified")
	} else if !filepath.IsAbs(configFilePath) {
		configFilePath = filepath.Join(wd, configFilePath)
	}
	config, err := NewConfig(configFilePath)
	if err != nil {
		return b, err
	}
	b.Config = config
	// register out dir if specified
	if outDir != "" && !filepath.IsAbs(outDir) {
		outDir = filepath.Join(wd, outDir)
	}
	b.OutputDirPath = outDir
	// register inputs
	inputs, err := createInputs(args, b.InputType.extNames(), wd)
	if err != nil {
		return b, err
	}
	b.Inputs = inputs
	return b, nil
}

func (c Bedrock) OutputsFiles() bool {
	return c.OutputDirPath != ""
}

func createInputs(args cli.Args, allowExts []string, wd string) ([]string, error) {
	res := []string{}
	for _, arg := range args {
		files, err := filepath.Glob(arg)
		if err != nil {
			return res, err
		}
		for _, file := range files {
			info, err := os.Stat(file)
			if err != nil {
				return res, err
			}
			if !info.IsDir() && contains(allowExts, filepath.Ext(file)) {
				res = append(res, filepath.Join(wd, file))
			}
		}
	}
	return res, nil
}

func contains(list []string, target string) bool {
	for _, val := range list {
		if val == target {
			return true
		}
	}
	return false
}