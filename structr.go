package main

import (
	"github.com/codegangsta/cli"
	_ "github.com/jteeuwen/go-bindata"
	"os"
)

var Version = "1.0.0"

func main() {
	app := cli.NewApp()
	app.Name = "structr"
	app.Usage = "Generate structure definition(s) from JSON Schema"
	app.UsageText = "structr [command] [command options] [filepath...]"
	app.Author = "dameleon"
	app.Email = "dameleon@gmail.com"
	app.Version = Version
	app.Commands = Commands
	app.Run(os.Args)
}
