package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

// KmsClient ... Store KMS client with a session
type KmsClient struct {
	Session *kms.KMS
}

// NewKmsClient ... The function to create KMS client with config settings.
func NewKmsClient() (*KmsClient, error) {
	c := &KmsClient{}

	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           AwsProfile,
		Config:            aws.Config{Region: aws.String(AwsRegion)},
	})

	if err != nil {
		return nil, fmt.Errorf("failed kms NewSessionWithOptions %+v", err)
	}

	c.Session = kms.New(sess, aws.NewConfig().WithMaxRetries(10))
	return c, nil
}

// ListAliases ... Gets a list of aliases.
func (k *KmsClient) ListAliases() (*kms.ListAliasesOutput, error) {
	param := &kms.ListAliasesInput{}
	result, err := k.Session.ListAliases(param)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// ListKmsAliasNamesToSlice ... Gets a list of alias names to slice.
func (k *KmsClient) ListKmsAliasNamesToSlice() ([]string, error) {
	var result []string
	list, err := k.ListAliases()
	if err != nil {
		return result, err
	}

	for _, v := range list.Aliases {
		result = append(result, aws.StringValue(v.AliasName))
	}

	return result, nil
}
