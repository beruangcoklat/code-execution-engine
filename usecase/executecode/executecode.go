package executecode

import (
	"context"

	"github.com/beruangcoklat/code-execution-engine/common"
	"github.com/beruangcoklat/code-execution-engine/pkg/lib/configlib"
	"github.com/beruangcoklat/code-execution-engine/pkg/lib/errorlib"
	"github.com/beruangcoklat/code-execution-engine/pkg/lib/executorlib"
	"github.com/beruangcoklat/code-execution-engine/pkg/lib/nsqlib"
)

func New() *Usecase {
	return &Usecase{}
}

func (uc *Usecase) ExecuteCode(ctx context.Context, payload ExecuteCodePayload) error {
	cfg := configlib.Get()

	executor, err := executorlib.New(executorlib.NewExecutorParam{
		ExecutorType: payload.LanguageID,
		Code:         payload.Code,
		Stdin:        payload.Stdin,
		TimeLimit:    payload.TimeLimit,
		MemoryLimit:  payload.MemoryLimit,
		WorkDir:      cfg.Executor.WorkDir,
	})
	if err != nil {
		return errorlib.AddTrace(err)
	}

	defer executor.Clean(ctx)

	compileResult, err := executor.Compile(ctx)
	if err != nil {
		return errorlib.AddTrace(err)
	}
	if compileResult.IsCompileError {
		err = uc.publishResult(compileResult)
		if err != nil {
			return errorlib.AddTrace(err)
		}
		return nil
	}

	runResult, err := executor.Run(ctx)
	if err != nil {
		return errorlib.AddTrace(err)
	}

	err = uc.publishResult(runResult)
	if err != nil {
		return errorlib.AddTrace(err)
	}

	return nil
}

func (uc *Usecase) publishResult(value interface{}) error {
	err := nsqlib.Publish(common.TOPIC_EXECUTION_RESULT, value)
	if err != nil {
		return errorlib.AddTrace(err)
	}
	return nil
}
