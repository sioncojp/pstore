package main

import (
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
)

const AppVersion = "1.0.0"

// AwsProfile ... local aws profile. ~/.aws/config
var AwsProfile string

// AwsRegion ... aws region
var AwsRegion string

// FilePath ... toml file path
var FilePath string

// NewVariable ... initialize global variables
func NewVariable(c *cli.Context) {
	FilePath = c.String("file")

	d, f := filepath.Split(filepath.Clean(c.String("file")))
	AwsProfile = strings.Split(d, "/")[len(strings.Split(d, "/"))-2]
	AwsRegion = strings.Split(f, ".")[0]
}
