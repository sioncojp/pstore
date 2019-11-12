package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/urfave/cli"
)

// CommandShow ... Output decrypt file
func CommandShow(c *cli.Context) error {
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

	output, err := DecryptFile(pstoreKey)
	if err != nil {
		return err
	}

	fmt.Println(string(output))

	return nil
}
