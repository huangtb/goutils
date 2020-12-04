package mq

import (
	"github.com/nsqio/go-nsq"
	"github.com/pkg/errors"
)

var np *NsqProducer

type NsqConfig struct {
	Addr string	`json:"addr" yaml:"addr"`
	Topic string `json:"topic" yaml:"topic"`
	Channel string `json:"channel" yaml:"channel"`
}

type NsqProducer struct {
	Producer *nsq.Producer
}

func GetNsqProducer() *NsqProducer {
	return np
}

func (c *NsqConfig) NewNsqProducer() (p NsqProducer,err error) {
	p.Producer, err = nsq.NewProducer(c.Addr, nsq.NewConfig())
	if err != nil {
		return p,errors.Errorf("New nsq producer error: %v", err)
	}
	err = p.Producer.Ping()
	if nil != err {
		p.Producer.Stop()
		return p,errors.Errorf("Ping nsq error: %v", err)
	}
	np = &p
	return p,nil
}

//发布消息
func (np *NsqProducer) PublishMsg(topic string, data []byte) error {
	if np == nil {
		return errors.Errorf("Nsq producer is nil")
	}
	if data == nil {
		return errors.Errorf("publish nil message")
	}
	return np.Producer.Publish(topic, data)
}



func (c *NsqConfig) NewNsqConsumer(handler nsq.Handler) (*nsq.Consumer,error) {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(c.Topic, c.Channel,config)
	if err != nil {
		return nil,errors.Errorf("Topic:%s ,Channel:%s new consumer error:%s", c.Topic, c.Channel, err.Error())
	}
	// 设置消息处理函数
	consumer.AddHandler(handler)
	if err := consumer.ConnectToNSQD(c.Addr); err != nil {

		return nil,errors.Errorf("Topic:%s ,Channel:%s connect to nsq error:%s", c.Topic, c.Channel, err.Error())
	}
	return consumer,nil
}
