package main

import (
	"github.com/codegangsta/cli"
	"os"
	"path/filepath"
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
		context, err := NewContext(c.String("config"), args[len(args) - 1], c.String("outDir"))
		if err != nil {
			panic(err)
		}
		files, err := filepath.Glob(context.InputPath)
		if err != nil {
			panic(err)
		}
		type structureHolder struct {
			main []StructureNode
			dependencies map[string]StructureNode
		}
		bundler := NewJsonSchemaBundler(NewJsonSchemaLoader())
		for _, file := range files {
			if info, _ := os.Stat(file); !info.IsDir() {
				bundler.AddJsonSchema(file)
			}
		}
		creator := NewJsonSchemaNodeCreator(context, bundler)
		for _, b := range bundler.GetBundles() {
			tmpl := NewContextualTemplate(context, b.GetName())
			tmpl.Execute(os.Stdout, creator.CreateStructureNode(b))
		}
	}
	app.Run(os.Args)
}
