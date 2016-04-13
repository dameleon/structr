package main

import (
	"path/filepath"
	"os"
	"errors"
	"github.com/codegangsta/cli"
	"github.com/asaskevich/govalidator"
)

type Context struct {
	Config Config
	OutputDirPath string
	Inputs []string
}

func NewContext(configFilePath string, outDir string, args cli.Args) (Context, error) {
	var context Context
	cwd, err := os.Getwd()
	if err != nil {
		return context, err
	}
	// create & register config
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
	// register out dir if specified
	if outDir != "" && !filepath.IsAbs(outDir) {
		outDir = filepath.Join(cwd, outDir)
	}
	context.OutputDirPath = outDir
	// register inputs
	inputs, err := context.createInputs(args)
	if err != nil {
		return context, err
	}
	context.Inputs = inputs
	return context, nil
}

func (c Context) createInputs(args cli.Args) ([]string, error) {
	res := []string{}
	for _, arg := range args {
		if govalidator.IsURL(arg) {
			res = append(res, arg)
		} else {
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
					res = append(res, arg)
				}
			}
		}
	}
	return res, nil
}

func (c Context) OutputsFiles() (bool) {
	return c.OutputDirPath != ""
}
