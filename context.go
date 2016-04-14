package main

import (
	"errors"
	"github.com/codegangsta/cli"
	"os"
	"path/filepath"
)

type Context struct {
	Config        Config
	OutputDirPath string
	Inputs        []string
}

func NewContext(configFilePath string, outDir string, args cli.Args) (Context, error) {
	var context Context
	wd, err := os.Getwd()
	if err != nil {
		return context, err
	}
	// create & register config
	if configFilePath == "" {
		return context, errors.New("config flag must be specified")
	} else if !filepath.IsAbs(configFilePath) {
		configFilePath = filepath.Join(wd, configFilePath)
	}
	config, err := NewConfig(configFilePath)
	if err != nil {
		return context, err
	}
	context.Config = config
	// register out dir if specified
	if outDir != "" && !filepath.IsAbs(outDir) {
		outDir = filepath.Join(wd, outDir)
	}
	context.OutputDirPath = outDir
	// register inputs
	inputs, err := createInputs(args, wd)
	if err != nil {
		return context, err
	}
	context.Inputs = inputs
	return context, nil
}

func (c Context) OutputsFiles() bool {
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
