package aws

import (
	"github.com/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

var ddbMgr *dynamoMgr

type dynamoMgr struct {
	client *dynamo.DB
}

func GetDynamoDB() *dynamoMgr {
	return ddbMgr
}

type Aws struct {
	Region    string
	Endpoint  string
	AccessKey string
	SecretKey string
}

func (a *Aws) NewDynamoDBMgr() (ddb dynamoMgr,err error) {
	ses, err := session.NewSession()
	if err != nil {
		return ddb,errors.Errorf("New DynamoDB Session error: %s" ,err.Error())
	}
	ddb.client = dynamo.New(ses, &aws.Config{
		Region:      aws.String(a.Region),
		Endpoint:    aws.String(a.Endpoint),
		Credentials: credentials.NewStaticCredentials(a.AccessKey, a.SecretKey, ""),
	})
	if _, err := ddb.client.ListTables().All(); err != nil {
		return ddb,errors.Errorf("New DynamoDB error: %s",err.Error())
	}
	ddbMgr = &ddb
	return ddb,nil
}

