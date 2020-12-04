package aws

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
	"time"
)

var fHose *Firehose

type Firehose struct {
	client *firehose.Firehose
}

func (a *Aws) NewFireHose() (*Firehose, error) {
	sess := session.Must(session.NewSession(aws.NewConfig().WithRegion(a.Region).
		WithCredentials(credentials.NewStaticCredentials(a.AccessKey,
			a.SecretKey, ""))))
	fHose.client = firehose.New(sess)
	return fHose, nil
}

func (f *Firehose) PutData(data, topic, StreamName string) error {
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
	_, err = f.client.PutRecord(params)
	return err
}
