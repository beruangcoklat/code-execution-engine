package nsqlib

import "github.com/nsqio/go-nsq"

type (
	NSQHandler func(message *nsq.Message) error

	NSQLibParam struct {
		NsqLookupdAddrs []string
		NsqdAddr        string
	}

	nsqLib struct {
		NSQLookupds []string
		Cfg         *nsq.Config
		Producer    *nsq.Producer
	}
)
