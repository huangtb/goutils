package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/pkg/errors"
)

var RekCli *rekognition.Rekognition

func (a *Aws) InitRekognition() error {

	sess, err := session.NewSession(a.GetConfig())
	if err != nil {
		return errors.Errorf("New Rekognition Session creation error: %v", err.Error())
	}
	RekCli = rekognition.New(sess)
	return nil

}
