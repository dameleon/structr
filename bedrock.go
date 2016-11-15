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
}

func NewBedrock(configFilePath string, outDir string, args cli.Args) (Bedrock, error) {
	var b Bedrock
	wd, err := os.Getwd()
	if err != nil {
		return b, err
	}
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
	inputs, err := createInputs(args, wd)
	if err != nil {
		return b, err
	}
	b.Inputs = inputs
	return b, nil
}

func (c Bedrock) OutputsFiles() bool {
	return c.OutputDirPath != ""
}

func createInputs(args cli.Args, wd string) ([]string, error) {
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
			if !info.IsDir() {
				res = append(res, filepath.Join(wd, arg))
			}
		}
	}
	return res, nil
}
