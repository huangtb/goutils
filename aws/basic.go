package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type Aws struct {
	Region    string `json:"region" yaml:"region"`
	Endpoint  string `json:"endpoint" yaml:"endpoint"`
	AccessKey string `json:"access_key" yaml:"access_key"`
	SecretKey string `json:"secret_key" yaml:"secret_key"`
}


func (a *Aws) GetConfig() *aws.Config {
	config := aws.Config{
		Region: aws.String(a.Region),
	}

	if a.Endpoint != "" {
		config.Endpoint = aws.String(a.Endpoint)
	}

	if a.AccessKey != "" && a.AccessKey != "" {
		config.Credentials = credentials.NewStaticCredentials(a.AccessKey, a.AccessKey, "")
	}

	return &config
}
