package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/pkg/errors"
)

var DdbCli *dynamo.DB



func (a *Aws) InitDynamoDBClient() error {

	ses, err := session.NewSession()
	if err != nil {
		return errors.Errorf("New DynamoDB Session error: %s", err.Error())
	}

	cred := getCredentials(a.AccessKey,a.SecretKey)
	_, err = cred.Get()
	if err != nil {
		return errors.Errorf("New Static Credentials  error:" , err.Error())
	}

	ddb := dynamo.New(ses, &aws.Config{
		Region:      aws.String(a.Region),
		Endpoint:    aws.String(a.Endpoint),
		Credentials: cred,
	})

	if _, err := ddb.ListTables().All(); err != nil {
		return errors.Errorf("New DynamoDB error: %s", err.Error())
	}
	DdbCli = ddb
	return nil
}
