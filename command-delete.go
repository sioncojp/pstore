package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/urfave/cli"
)

// CommandDelete ... Delete SSM and file data
func CommandDelete(c *cli.Context) error {
	NewVariable(c)

	client, err := NewSsmClient()
	if err != nil {
		return err
	}

	p, err := client.GetParameter(c.String("pstore-key"))
	if err != nil {
		return err
	}

	pstoreKey := aws.StringValue(p.Parameter.Value)

	config := &Config{}
	if err := config.Load(pstoreKey); err != nil {
		return err
	}

	d := &DeleteData{}
	if err := d.DeletePrompt(config); err != nil {
		return err
	}

	// delete parameter
	if _, err := client.DeleteParameter(d.Name); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("finished!!")

	return nil
}
