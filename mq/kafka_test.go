package mq

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"
	"testing"
)

var (
	addr = []string{"broker1", "broker2", "broker3"}
	ktopic = "dmp-ken-test2"
	kgroup = "ebk-risk-consumer"
)

type NewKafkaMessage struct {
	Data  interface{} `json:"data"`
	MsgId string       `json:"msgId"`
	Type  string       `json:"type"`
}

type SqsData struct {
	CustomerID string `json:"customer_id" validate:"required"` //用户BVN
	DeviceID   string `json:"device_id" validate:"required"`   //设备id
	UID        string `json:"uid" validate:"required"`         //uid
	BatchNo    string `json:"batch_no,omitempty"`              //批次号
	TradeNo    string `json:"trade_no,omitempty"`              //批次号
	Country    string `json:"country"`                         //国家缩写
	Queue      string `json:"queue"`                           //SQS队列名
	ReqID      string `json:"req_id"`
}





func TestProducer(t *testing.T) {
	c := KafkaConfig{
		Brokers:addr,
		Topic: ktopic,
		GroupID:kgroup,
	}

	_, err := c.NewKafkaProducer(NewConfig())
	if err != nil {
		panic(err)
	}

	data := SqsData{
		BatchNo:"2100",
		DeviceID:"qwerty",
	}

	newMsg := NewKafkaMessage{
		Data:  data,
		MsgId: "123777",
		Type:  "thirdPartyRepayment",
	}
	newJsonMsg, err := json.Marshal(&newMsg)
	if err != nil {
		panic(err)
	}
	err = GetKafkaProducer().Publish("ebk-risk-result", string(newJsonMsg))
	if err != nil {
		fmt.Printf("发布失败:%v", err)
		return
	}
		fmt.Printf("发布成功")


}

var (
	topic = flag.String("t", "dmp-ken-test2", "kafka topic")
	group = flag.String("g", "ebk-risk-consumer2", "kafka consumer group")
)

func TestReader(t *testing.T) {


	c := KafkaConfig{
		Brokers:addr,
		Topic: ktopic,
		GroupID:kgroup,
	}

	reader := c.NewKafkaReader()
	ctx := context.Background()
	//reader.SetOffset(1)

	log.Printf("staring...")
	for {

		log.Printf("reader>>%d",reader.Offset())
		msg, err := reader.FetchMessage(ctx)
		if err != nil {
			log.Printf("fail to get msg:%v", err)
			continue
		}
		log.Printf("msg content:topic=%v,partition=%v,offset=%v,content=%v",
			msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		if strings.Contains(string(msg.Value),"65") {
			log.Printf("message 需要继续处理:%s", string(msg.Value))
			//startTime := time.Now()
			continue
		}
		err = reader.CommitMessages(ctx, msg)
		if err != nil {
			log.Printf("fail to commit msg:%v", err)
		}
		log.Printf("message处理 完成")

	}

}
