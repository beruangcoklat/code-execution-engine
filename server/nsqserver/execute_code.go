package nsqserver

import (
	"context"

	"github.com/beruangcoklat/code-execution-engine/pkg/lib/loglib"
	"github.com/beruangcoklat/code-execution-engine/server"
	"github.com/beruangcoklat/code-execution-engine/usecase/executecode"
	jsoniter "github.com/json-iterator/go"
	"github.com/nsqio/go-nsq"
)

func executeCode(msg *nsq.Message) error {
	ctx := context.Background()

	var payload executecode.ExecuteCodePayload
	var err error

	err = jsoniter.Unmarshal(msg.Body, &payload)
	if err != nil {
		msg.Finish()
		return nil
	}

	err = server.ExecutecodeUc.ExecuteCode(ctx, payload)
	if err != nil {
		loglib.ErrorTrace(err, "error execute code")
		msg.Finish()
		return nil
	}

	msg.Finish()
	return nil
}
