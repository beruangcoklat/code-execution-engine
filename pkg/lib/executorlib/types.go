package executorlib

import "context"

type (
	ExecutorItf interface {
		Compile(ctx context.Context) (*ExecutionResult, error)
		Run(ctx context.Context) (*ExecutionResult, error)
		Clean(ctx context.Context) error
	}

	CPPExecutor struct {
		WorkDir     string
		Code        string
		Stdin       string
		TimeLimit   int64
		MemoryLimit int64

		buildPath string
	}

	NewExecutorParam struct {
		ExecutorType int
		WorkDir      string
		Code         string
		Stdin        string
		TimeLimit    int64
		MemoryLimit  int64
	}

	ExecutionResult struct {
		Stdout              string
		Stderr              string
		IsTimeLimitExceeded bool
		IsCompileError      bool
		Error               string
	}
)
