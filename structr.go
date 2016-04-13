package main

import (
	"github.com/codegangsta/cli"
	"os"
	"log"
)

func main() {
	app := cli.NewApp()
	app.Name = "structr"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "config",
			Value: "",
			Usage: "configuration for structr",
		},
		cli.StringFlag{
			Name: "outDir",
			Value: "",
			Usage: "output directory for generated structure",
		},
	}
	app.Action = func(c *cli.Context) {
		args := c.Args()
		if len(args) < 1 {
			cli.ShowAppHelp(c)
			os.Exit(1)
		}
		context, err := NewContext(c.String("config"), c.String("outDir"), args)
		if err != nil {
			log.Fatalln("initialize error: ", err.Error())
		}
		bundler := NewJsonSchemaBundler(NewJsonSchemaLoader())
		if err := bundler.AddJsonSchema(context.Inputs...); err != nil {
			log.Fatalln("cannot add load json schema: ", err.Error())
		}
		creator := NewJsonSchemaNodeCreator(context, bundler)
		exporter := NewExporter(context)
		for _, b := range bundler.GetBundles() {
			structure, err := creator.CreateStructureNode(b)
			if err != nil {
				log.Fatalln("cannot create structure node: ", err.Error())
			}
			if err := exporter.Export(structure); err != nil {
				log.Fatalln("cannot export structure node: ", err.Error())
			}
		}
	}
	app.Run(os.Args)
}
