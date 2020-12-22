package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/pkg/errors"
)

var DdbCli *dynamo.DB



func (a *Aws) InitDynamoDBClient() error {

	sess, err := session.NewSession(a.GetConfig())
	if err != nil {
		return errors.Errorf("New DynamoDB Session creation error: %v" , err.Error())
	}

	ddb := dynamo.New(sess)

	if _, err := ddb.ListTables().All(); err != nil {
		return errors.Errorf("New DynamoDB error: %s", err.Error())
	}
	DdbCli = ddb
	return nil
}
