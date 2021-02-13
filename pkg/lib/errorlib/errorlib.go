package errorlib

import (
	"fmt"
	"runtime"
	"strings"
)

var (
	prefixToBeRemoved string

	TimeLimitError = &CustomError{
		errString: "Error Time Limit Exceeded",
	}

	MemoryLimitError = &CustomError{
		errString: "Error Memory Limit Exceeded",
	}
)

func Init(prefix string) {
	prefixToBeRemoved = prefix
}

func AddTrace(err interface{}) *CustomError {
	_, file, line, _ := runtime.Caller(1)
	trace := fmt.Sprintf("%v [%v]", file, line)
	trace = strings.TrimPrefix(trace, prefixToBeRemoved)
	customErr, ok := err.(*CustomError)
	if !ok {
		customErr = &CustomError{errString: err.(error).Error()}
	}
	customErr.Traces = append(customErr.Traces, trace)
	return customErr
}

func (ce *CustomError) Error() string {
	return ce.errString
}

func ToString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
