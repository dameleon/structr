package main

import (
	"path/filepath"
	"os"
	"errors"
)

type Context struct {
	InputPath string
	OutputDirPath string
	Config Config
}

func NewContext(configFilePath string, inputPath string, outDir string) (Context, error) {
	var context Context
	cwd, err := os.Getwd()
	if err != nil {
		return context, err
	}
	if configFilePath == "" {
		return context, errors.New("config flag must be specified")
	} else if !filepath.IsAbs(configFilePath) {
		configFilePath = filepath.Join(cwd, configFilePath)
	}
	config, err := NewConfig(configFilePath)
	if err != nil {
		return context, err
	}
	context.Config = config
	if !filepath.IsAbs(inputPath) {
		inputPath = filepath.Join(cwd, inputPath)
	}
	context.InputPath = inputPath
	if outDir != "" && !filepath.IsAbs(outDir) {
		outDir = filepath.Join(cwd, outDir)
	}
	context.OutputDirPath = outDir
	return context, nil
}
