package executorlib

import (
	"errors"

	"github.com/beruangcoklat/code-execution-engine/common"
)

func New(param NewExecutorParam) (ExecutorItf, error) {
	if param.ExecutorType == common.CPPType {
		return &CPPExecutor{
			WorkDir:     param.WorkDir,
			Code:        param.Code,
			Stdin:       param.Stdin,
			TimeLimit:   param.TimeLimit,
			MemoryLimit: param.MemoryLimit,
		}, nil
	}

	return nil, errors.New("Executor not found")
}
