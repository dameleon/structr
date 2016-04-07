package main

import (
	"github.com/codegangsta/cli"
	"os"
	"path/filepath"
	"log"
	"github.com/k0kubun/pp"
)

func main() {
	app := cli.NewApp()
	app.Name = "structr"
	app.Usage = "generate struct"
	app.Action = func(c *cli.Context) {
		file := c.Args().First()
		if file == "" {
			file = "hoge.json"
		}
		if !filepath.IsAbs(file) {
			cwd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			file = filepath.Join(cwd, file)
		}
		bundler := NewJsonSchemaBundler(NewJsonReference(file))
		bundler.Bundle()
		pp.Print(bundler.GetSchema("#/definitions/address"))
		// parser := json.NewDecoder(file)
		// var s JsonSchema
		// e := parser.Decode(&s)
		// if e != nil {
		// 	println(e.Error())
		// 	return
		// }
		// node, _ := NewNode(NodeParam{ Id: "hoge", Schema: &s })
		// tpl := template.Must(template.ParseFiles("ObjectMapper.tmpl"))
		// if te := tpl.Execute(os.Stdout, node); te != nil {
		// 	println(te)
		// }
	}
	app.Run(os.Args)
}
