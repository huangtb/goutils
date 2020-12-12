package aws

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
)

type Aws struct {
	Region    string `json:"region" yaml:"region"`
	Endpoint  string `json:"endpoint" yaml:"endpoint"`
	AccessKey string `json:"access_key" yaml:"access_key"`
	SecretKey string `json:"secret_key" yaml:"secret_key"`
}

func getCredentials(ak, sk string) *credentials.Credentials {

	if ak != "" && sk != "" {
		return credentials.NewStaticCredentials(ak, sk, "")
	}

	return credentials.NewCredentials(&ec2rolecreds.EC2RoleProvider{})
}
