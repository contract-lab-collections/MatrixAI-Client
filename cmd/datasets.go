package cmd

import (
	"MatrixAI-Client/logs"
	"fmt"
	"github.com/cavaliergopher/grab/v3"
	"github.com/urfave/cli"
)

var DatasetsCommand = cli.Command{
	Name:  "datasets",
	Usage: "Upload or download a dataset of AI models.",
	Subcommands: []cli.Command{
		{
			Name:  "download",
			Usage: "Download the dataset of the AI model.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "url, u",
					Required: true,
					Usage:    "Source of the dataset.",
				},
			},
			Action: func(c *cli.Context) error {

				url := c.String("url")
				logs.Result(fmt.Sprintf("url: %v", url))

				resp, err := grab.Get("./datasets", url)
				if err != nil {
					fmt.Printf("Error downloading file: %v\n", err)
					return err
				}

				logs.Result(fmt.Sprintf("resp name: %v", resp.Filename))
				return nil
			},
		},
	},
}
