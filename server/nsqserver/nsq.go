package nsqserver

import (
	"strings"

	"github.com/beruangcoklat/code-execution-engine/common"
	"github.com/beruangcoklat/code-execution-engine/pkg/lib/configlib"
	"github.com/beruangcoklat/code-execution-engine/pkg/lib/nsqlib"
)

func InitNSQ() error {
	cfg := configlib.Get()

	err := nsqlib.Init(nsqlib.NSQLibParam{
		NsqLookupdAddrs: strings.Split(cfg.NSQ.NsqLookupd, ","),
		NsqdAddr:        cfg.NSQ.Nsqd,
	})

	if err != nil {
		return err
	}

	nsqlib.RegisterConsumer(executeCode, common.TOPIC_EXECUTE_CODE, common.CHANNEL_EXECUTE_CODE, 5)

	return nil
}
