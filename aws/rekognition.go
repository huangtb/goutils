package aws

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

var rekClient *rekognition.Rekognition

func GetRekClient() *rekognition.Rekognition {
	return rekClient
}

func (a *Aws) NewRekognition() error{

	cred := credentials.NewStaticCredentials(a.AccessKey, a.SecretKey, "")
	sess, err := session.NewSession(&aws.Config{
		Credentials: cred,
		Region:      aws.String(a.Region),
	})
	if err != nil {
		errors.New("New rekognition session err:" +err.Error())
	}
	rek := rekognition.New(sess)
	if rek == nil {
		errors.New("New rekognition client is nil" +err.Error())
	}
	return nil
}
