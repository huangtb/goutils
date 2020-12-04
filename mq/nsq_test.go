package mq

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"testing"
)

var (
	nsqAddr    = "IP:4150"
	nsqTopic   = "nsq_topic_ebk_login"
	nsqChannel = "nsq_channel_login"
)

func TestNewNsqProducer(t *testing.T) {
	c := NsqConfig{
		Addr:    nsqAddr,
		Topic:   nsqTopic,
		Channel: nsqChannel,
	}
	c.NewNsqProducer()
	GetNsqProducer().PublishMsg(nsqTopic, []byte("7777"))
}

type loginHandler struct {
}

func (loginHandler) HandleMessage(message *nsq.Message) error {
	fmt.Printf("message.Body>>%s \n", message.Body)
	return nil
}

func TestNsqConsumer(t *testing.T) {
	c := NsqConfig{
		Addr:    nsqAddr,
		Topic:   nsqTopic,
		Channel: nsqChannel,
	}
	_, err := c.NewNsqConsumer(loginHandler{})
	if err != nil {
		panic(err)
	}
	<-make(chan bool)
}
