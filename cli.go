package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli"
)

// FlagSet ... set flag for cli
func FlagSet() *cli.App {
	app := cli.NewApp()
	app.Name = "pstore"
	app.Usage = "pstore"
	app.Version = AppVersion

	// set subcommand
	app.Commands = NewCommand()

	// sort command
	sort.Sort(cli.CommandsByName(app.Commands))

	return app
}

// NewCommand ... create cli command
func NewCommand() []cli.Command {
	return []cli.Command{
		{
			Name:  "add",
			Usage: "add data",
			Flags: WithDefaultCliFlag([]cli.Flag{}),
			Action: func(c *cli.Context) error {
				return CommandAdd(c)
			},
		},
		{
			Name:  "delete",
			Usage: "delete data",
			Flags: WithDefaultCliFlag([]cli.Flag{}),
			Action: func(c *cli.Context) error {
				if !IsFileNotEmpty(c.String("file")) {
					return nil
				}
				return CommandDelete(c)
			},
		},
		{
			Name:  "show",
			Usage: "show decode data",
			Flags: WithDefaultCliFlag([]cli.Flag{}),
			Action: func(c *cli.Context) error {
				if !IsFileNotEmpty(c.String("file")) {
					return nil
				}
				return CommandShow(c)
			},
		},
		{
			Name:  "list-kms",
			Usage: "list kms alias",
			Flags: WithDefaultCliFlag([]cli.Flag{}),
			Action: func(c *cli.Context) error {
				return CommandListKms(c)
			},
		},
	}
}

// WithDefaultCliFlag ... default flag
func WithDefaultCliFlag(additionalFlag []cli.Flag) []cli.Flag {
	defaultFlag := []cli.Flag{
		cli.StringFlag{
			Name:     "file",
			Usage:    "yaml file. (e.g. AwsProfileName/ap-northeast-1.yml)",
			Required: true,
		},
		cli.StringFlag{
			Name:  "pstore-key",
			Usage: "pstore key is encrypt data.",
			Value: "/pstore/key",
		},
	}

	return append(defaultFlag, additionalFlag...)
}

// IsFileNotEmpty ... validate file does not exist and non-zero size
func IsFileNotEmpty(file string) bool {
	fileinfo, staterr := os.Stat(file)

	if staterr != nil {
		fmt.Println("does not exist file")
		return false
	}

	// ファイルサイズを表示
	if fileinfo.Size() == 0 {
		fmt.Println("file is empty data")
		return false
	}

	return true
}
