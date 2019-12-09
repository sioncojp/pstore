package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/urfave/cli"
)

// CommandDelete ... Delete SSM and file data
func CommandDelete(c *cli.Context) error {
	NewVariable(c)
	d := &Data{}

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

	deleteData := &DeleteData{}
	if err := deleteData.DeletePrompt(config); err != nil {
		return err
	}

	// delete parameter
	var isNotExistParameterStore bool
	if _, err := client.DeleteParameter(deleteData.Name); err != nil {
		isNotExistParameterStore = true
	}
	if isNotExistParameterStore {
		if err := deleteData.ForceDeletePrompt(config); err != nil {
			return err
		}
	}

	// encrypt data to yml
	if err := d.DecryptFileAndWriteFileWithDeleteData(pstoreKey, deleteData.Name); err != nil {
		return err
	}
	if err := EncryptFile(pstoreKey); err != nil {
		return err
	}


	fmt.Println()
	fmt.Println("finished!!")

	return nil
}
