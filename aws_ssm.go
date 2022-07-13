package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// SsmClient ... Store SSM client with a session
type SsmClient struct {
	Session *ssm.SSM
}

// NewSsmClient ... The function to create SSM client with config settings.
func NewSsmClient() (*SsmClient, error) {
	c := &SsmClient{}

	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           AwsProfile,
		Config:            aws.Config{Region: aws.String(AwsRegion)},
	})

	if err != nil {
		return nil, fmt.Errorf("failed ssm NewSessionWithOptions %+v", err)
	}

	c.Session = ssm.New(sess, aws.NewConfig().WithMaxRetries(10))
	return c, nil
}

// GetParameter ... get parameter
func (s *SsmClient) GetParameter(name string) (*ssm.GetParameterOutput, error) {
	param := &ssm.GetParameterInput{
		Name:           aws.String(name),
		WithDecryption: aws.Bool(true),
	}
	result, err := s.Session.GetParameter(param)
	if err != nil {
		return nil, fmt.Errorf("failed ssm get param: %+v", err)
	}

	return result, nil
}

// PutParameter ... put parameter
func (s *SsmClient) PutParameter(name, value, ssmtype, kmsAlias, description string) (*ssm.PutParameterOutput, error) {
	param := &ssm.PutParameterInput{
		Name:        aws.String(name),
		Value:       aws.String(value),
		Type:        aws.String(ssmtype),
		Description: aws.String(description),
		Overwrite:   aws.Bool(true),
	}

	if ssmtype == "SecureString" {
		param.KeyId = aws.String(kmsAlias)
	}

	result, err := s.Session.PutParameter(param)
	if err != nil {
		return nil, fmt.Errorf("failed ssm put param: %+v", err)
	}

	return result, nil
}

// DeleteParameter ... delete parameter
func (s *SsmClient) DeleteParameter(name string) (*ssm.DeleteParameterOutput, error) {
	param := &ssm.DeleteParameterInput{
		Name: aws.String(name),
	}
	result, err := s.Session.DeleteParameter(param)
	if err != nil {
		return nil, fmt.Errorf("failed ssm delete param: %+v", err)
	}

	return result, nil
}
