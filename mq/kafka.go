package mq

import (
	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)

//生产者使用sarama
//消费者使用kafka-go

var kp *KafkaProducer

type KafkaProducer struct {
	Producer sarama.SyncProducer
}

type KafkaConfig struct {
	Brokers []string
	GroupID string
	Topic   string
}

func GetKafkaProducer() *KafkaProducer {
	return kp
}

func NewConfig() *sarama.Config {
	return sarama.NewConfig()
}

func (c *KafkaConfig) NewKafkaProducer(config *sarama.Config) (*KafkaProducer, error) {
	prd := &KafkaProducer{}
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForLocal
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true

	// 使用给定代理地址和配置创建一个同步生产者
	newkp, err := sarama.NewSyncProducer(c.Brokers, config)
	if err != nil {
		return kp, errors.Errorf("Kafka addr[%s] New SyncProducer error:%v", c.Brokers, err)
	}
	prd.Producer = newkp
	kp = prd
	return prd, nil
}

func (kp *KafkaProducer) Publish(topic, value string) error {
	// 构造一个消息
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(value),
	}
	_, _, err := kp.Producer.SendMessage(msg)
	if err != nil {
		return errors.Errorf("Topic[%s] send msg[%s] failed, err:%v", topic, value, err)
	}
	return nil
}



func (c *KafkaConfig) NewKafkaReader() *kafka.Reader {
	config := kafka.ReaderConfig{
		Brokers:  c.Brokers,
		Topic:    c.Topic,
		GroupID:  c.GroupID,
		StartOffset: kafka.FirstOffset,
		MinBytes: 1e3,
		MaxBytes: 1e6,

	}
	return kafka.NewReader(config)
}
