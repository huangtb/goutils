package aws

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/pkg/errors"
	"time"
)

var FHoseCli *firehose.Firehose

func (a *Aws) InitFireHoseClient() error {

	cred := getCredentials(a.AccessKey,a.SecretKey)
	_, err := cred.Get()
	if err != nil {
		return errors.Errorf("New Static Credentials  error:" , err.Error())
	}

	sess := session.Must(session.NewSession(aws.NewConfig().WithRegion(a.Region).
		WithCredentials(cred)))

	f := firehose.New(sess)
	FHoseCli = f
	return nil
}

func PutToFireHose(data, topic, StreamName string) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	str := fmt.Sprintf("%d\t%s\t%s\n", time.Now().Unix(), string(b), topic)
	params := &firehose.PutRecordInput{
		DeliveryStreamName: aws.String(StreamName),
		Record: &firehose.Record{
			Data: []byte(str),
		},
	}
	_, err = FHoseCli.PutRecord(params)
	return err
}
