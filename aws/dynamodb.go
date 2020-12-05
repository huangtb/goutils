package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/pkg/errors"
)

var DdbCli *dynamo.DB

type Aws struct {
	Region    string
	Endpoint  string
	AccessKey string
	SecretKey string
}

func (a *Aws) InitDynamoDBClient() error {
	ses, err := session.NewSession()
	if err != nil {
		return errors.Errorf("New DynamoDB Session error: %s", err.Error())
	}
	ddb := dynamo.New(ses, &aws.Config{
		Region:      aws.String(a.Region),
		Endpoint:    aws.String(a.Endpoint),
		Credentials: credentials.NewStaticCredentials(a.AccessKey, a.SecretKey, ""),
	})
	if _, err := ddb.ListTables().All(); err != nil {
		return errors.Errorf("New DynamoDB error: %s", err.Error())
	}
	DdbCli = ddb
	return nil
}
