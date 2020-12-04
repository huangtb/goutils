package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
	"testing"
)

const (
	QURL      = "https://sqs.us-west-2.amazonaws.com/你的账号/你的队列"
	Region    = "你的 Region"
	AccessKey = "你的 AccessKey"
	SecretKey = "你的 SecretKey"
)

func TestSendToSQS(t *testing.T) {
	aws := Aws{
		Region:    Region,
		AccessKey: AccessKey,
		SecretKey: SecretKey,
	}
	err := aws.InitSqsClient()
	if err != nil {
		panic(err)
	}

	resp, _ := SendToSQS(QURL, "test SendToSQS ")
	log.Printf("resp: %+v", resp)

}

func TestConsumer(t *testing.T) {
	aws := Aws{
		Region:    Region,
		AccessKey: AccessKey,
		SecretKey: SecretKey,
	}
	err := aws.InitSqsClient()
	if err != nil {
		panic(err)
	}
	HandleApplyInfoMessage(QURL)

}

type ApplyMessage struct {
}

func (m *ApplyMessage) HandleMessage(messages []*sqs.Message) error {
	log.Printf("Receive message:%s", messages)
	return nil

}

func HandleApplyInfoMessage(queueUrl string) {

	consumer := NewConsumer(queueUrl)
	input := consumer.NewInputParams()
	input.WaitTimeSeconds = aws.Int64(20)

	consumer.AddHandler(input, &ApplyMessage{})

	<-make(chan bool)
}
