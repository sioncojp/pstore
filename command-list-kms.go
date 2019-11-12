package main

import (
	"fmt"

	"github.com/urfave/cli"
)

// CommandListKms ... Output kms aliases list
func CommandListKms(c *cli.Context) error {
	NewVariable(c)
	client, err := NewKmsClient()
	if err != nil {
		return err
	}

	kmsAliasList, err := client.ListKmsAliasNamesToSlice()
	if err != nil {
		return err
	}

	for _, v := range kmsAliasList {
		fmt.Println(v)
	}

	return nil
}
