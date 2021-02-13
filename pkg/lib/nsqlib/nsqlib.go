package nsqlib

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/nsqio/go-nsq"
)

var (
	single *nsqLib
)

func Init(param NSQLibParam) error {
	cfg := nsq.NewConfig()

	producer, err := nsq.NewProducer(param.NsqdAddr, cfg)
	if err != nil {
		return err
	}

	single = &nsqLib{
		NSQLookupds: param.NsqLookupdAddrs,
		Cfg:         cfg,
		Producer:    producer,
	}

	return nil
}

func RegisterConsumer(handler NSQHandler, topic, channel string, concurrency int) error {
	q, err := nsq.NewConsumer(topic, channel, single.Cfg)
	if err != nil {
		return err
	}

	q.AddConcurrentHandlers(nsq.HandlerFunc(handler), concurrency)

	err = q.ConnectToNSQLookupds(single.NSQLookupds)
	if err != nil {
		return err
	}

	return nil
}

func Publish(topic string, value interface{}) error {
	body, err := jsoniter.Marshal(value)
	if err != nil {
		return err
	}

	return single.Producer.Publish(topic, body)
}
