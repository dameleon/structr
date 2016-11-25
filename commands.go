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
	commandEnvironment,
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
			Name:  "type, t",
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
		var creator NodeCreator
		switch bedrock.InputType {
		case INPUT_TYPE_JSON:
			bundler := NewJsonSchemaBundler(NewJsonSchemaLoader())
			if err := bundler.AddJsonSchema(bedrock.Inputs...); err != nil {
				log.Fatalln("cannot add load json schema: ", err.Error())
			}
			creator = NewJsonSchemaNodeCreator(bundler)
		case INPUT_TYPE_API_BLUEPRINT:

		default:
			log.Fatalln("unknown input type")
		}
		if err := creator.Create(NewExporter(bedrock)); err != nil {
			log.Fatalln("failed export: ", err.Error())
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

var commandEnvironment = cli.Command{
	Name:      "env",
	Usage:     "Print environment",
	UsageText: "structr env",
	Action: func(c *cli.Context) {
		env := NewEnvironment()
		fmt.Printf("Usable: json_schema: %t, API Blueprint: %t\n", env.JsonSchema, env.ApiBlueprint)
		if env.ApiBlueprint {
			fmt.Printf("drafter(API Blueprint) bin path: %s\n", env.drafterBinPath)
		}
	},
}
