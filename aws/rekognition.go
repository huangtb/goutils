package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/pkg/errors"
)

var RekCli *rekognition.Rekognition

func (a *Aws) NewRekognition() error {

	cred := credentials.NewStaticCredentials(a.AccessKey, a.SecretKey, "")
	sess, err := session.NewSession(&aws.Config{
		Credentials: cred,
		Region:      aws.String(a.Region),
	})
	if err != nil {
		errors.Errorf("New rekognition session err: %v", err.Error())
	}
	rek := rekognition.New(sess)
	if rek == nil {
		errors.Errorf("New rekognition client is nil : %v ", err.Error())
	}
	return nil
}
