package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/pkg/errors"
)

var SqsCli *sqs.SQS

func (a *Aws) InitSqsClient() error {

	cred := getCredentials(a.AccessKey,a.SecretKey)
	_, err := cred.Get()
	if err != nil {
		return errors.Errorf("New Static Credentials  error:" , err.Error())
	}

	ses, err := session.NewSession(&aws.Config{
		Region:      aws.String(a.Region),
		Credentials: cred,
		MaxRetries:  aws.Int(5),
	})
	if err != nil {
		return errors.Errorf("New aws session error:%v", err.Error())
	}

	sc := sqs.New(ses)
	if sc.Client == nil {
		return errors.Errorf("New sqs client error")
	}
	SqsCli = sc
	return nil
}

func SendToSQS(queueUrl, message string) (*sqs.SendMessageOutput, error) {
	sendParams := &sqs.SendMessageInput{
		MessageBody: aws.String(message),  // Required
		QueueUrl:    aws.String(queueUrl), // Required
	}
	output, err := SqsCli.SendMessage(sendParams)
	if err != nil {
		return nil, errors.Errorf("New sqs client error:%v", err.Error())
	}
	return output, nil
}

type Handler interface {
	HandleMessage(queueUrl string,messages []*sqs.Message) error
}

type Consumer struct {
	QueueUrl string
}

func NewConsumer(queueUrl string) *Consumer {
	return &Consumer{
		QueueUrl: queueUrl,
	}
}

func (c *Consumer) AddHandler(input *sqs.ReceiveMessageInput, handler Handler) {
	go c.handlerLoop(input, handler)
}

func (c *Consumer) NewInputParams() *sqs.ReceiveMessageInput {
	return &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(c.QueueUrl),
		MaxNumberOfMessages: aws.Int64(3),  //单次最大接收消息数
		WaitTimeSeconds:     aws.Int64(10), //长轮询
	}
}

func (c *Consumer) handlerLoop(input *sqs.ReceiveMessageInput, handler Handler) error {
	for {
		output, _ := SqsCli.ReceiveMessage(input)
		err := handler.HandleMessage(c.QueueUrl,output.Messages)
		if err != nil {
			return errors.Errorf("New sqs client error:%v", err.Error())
		}
	}
}
