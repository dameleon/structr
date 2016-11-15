package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

var Commands = []cli.Command{
	commandGenerate,
	commandTemplate,
}

var commandGenerate = cli.Command{
	Name:      "generate",
	Usage:     "Generate structure definition(s)",
	UsageText: "structr generate -c YOUR_CONFIGURATION_FILE [filepath...]",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "",
			Usage: "(REQUIRED) configuration file for structr",
		},
		cli.StringFlag{
			Name: "type, t",
			Value: "json",
			Usage: "input file type",
		},
		cli.StringFlag{
			Name:  "outDir",
			Value: "",
			Usage: "output directory for generated structure",
		},
	},
	Action: func(c *cli.Context) {
		args := c.Args()
		if len(args) < 1 {
			cli.ShowAppHelp(c)
			os.Exit(1)
		}
		bedrock, err := NewBedrock(c.String("config"), c.String("type"), c.String("outDir"), args)
		if err != nil {
			log.Fatalln("initialize error: ", err.Error())
		}
		bundler := NewJsonSchemaBundler(NewJsonSchemaLoader())
		if err := bundler.AddJsonSchema(bedrock.Inputs...); err != nil {
			log.Fatalln("cannot add load json schema: ", err.Error())
		}
		creator := NewJsonSchemaNodeCreator(bundler)
		exporter := NewExporter(bedrock)
		for _, b := range bundler.GetBundles() {
			if !b.Schema.HasStructure() {
				continue
			}
			structure, err := creator.CreateStructureNode(b.Name, b)
			if err != nil {
				log.Fatalln("cannot create structure node: ", err.Error())
			}
			if err := exporter.Export(structure); err != nil {
				log.Fatalln("cannot export structure node: ", err.Error())
			}
		}
	},
}

var commandTemplate = cli.Command{
	Name:      "template",
	Usage:     "Output template of configuration file",
	UsageText: "structr template > YOUR_CONFIGURATION_NAME.yml",
	Action: func(c *cli.Context) {
		bytes, err := Asset("resources/config.yml")
		if err != nil {
			log.Fatalln("cannot load configuration template: ", err)
		}
		fmt.Print(string(bytes[:]))
	},
}
