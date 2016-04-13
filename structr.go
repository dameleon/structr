package main

import (
	"github.com/codegangsta/cli"
	"os"
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
			panic(err)
		}
		bundler := NewJsonSchemaBundler(NewJsonSchemaLoader())
		bundler.AddJsonSchema(context.Inputs...)
		creator := NewJsonSchemaNodeCreator(context, bundler)
		exporter := NewExporter(context)
		for _, b := range bundler.GetBundles() {
			exporter.Export(creator.CreateStructureNode(b))
		}
	}
	app.Run(os.Args)
}
