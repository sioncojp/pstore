package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/urfave/cli"
)

// CommandAdd ... Add SSM and file data
func CommandAdd(c *cli.Context) error {
	// initialize
	NewVariable(c)
	d := &Data{}
	client, err := NewSsmClient()
	if err != nil {
		return err
	}

	clientKms, err := NewKmsClient()
	if err != nil {
		return err
	}

	// get pstore key
	p, err := client.GetParameter(c.String("pstore-key"))
	if err != nil {
		return err
	}

	pstoreKey := aws.StringValue(p.Parameter.Value)

	// prompt
	kmsAliasList, err := clientKms.ListKmsAliasNamesToSlice()
	if err != nil {
		return err
	}

	if err := d.AddPrompt(kmsAliasList); err != nil {
		return err
	}

	// put
	if _, err := client.PutParameter(d.Name, d.Value, d.Type, d.KmsAlias, d.Description); err != nil {
		return err
	}

	// encrypt data to yml
	if err := d.DecryptFileAndWriteFileWithAddData(pstoreKey); err != nil {
		return err
	}
	if err := EncryptFile(pstoreKey); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("finished!!")

	return nil
}
