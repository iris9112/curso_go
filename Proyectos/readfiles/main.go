package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "dir",
			Value: "",
			Usage: "dir for files",
		},
		cli.StringFlag{
			Name:  "output",
			Value: "",
			Usage: "Name file output",
		},
	}

	app.Action = func(c *cli.Context) error {
		dir := c.String("dir")
		output := c.String("output")

		return concatFiles(dir, output)

	}

	app.Run(os.Args)
}

//llamada al archivo: go run *.go --dir directorio --output archivoSalida
//go run *.go --dir "/home/isabel/golibs/src/github.com/iris9112/readfiles/textos" --output "file.txt"