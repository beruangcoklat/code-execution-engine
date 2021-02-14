package executorlib

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/beruangcoklat/code-execution-engine/pkg/lib/errorlib"
)

func (ex *CPPExecutor) Compile(ctx context.Context) (*ExecutionResult, error) {
	uniqueFolder := fmt.Sprintf("%v", time.Now().UnixNano())
	path := filepath.Join(ex.WorkDir, uniqueFolder)
	ex.buildPath = path

	var err error

	err = os.MkdirAll(path, 0777)
	if err != nil {
		return nil, errorlib.AddTrace(err)
	}

	err = ioutil.WriteFile(filepath.Join(path, "main.cpp"), []byte(ex.Code), 0644)
	if err != nil {
		return nil, errorlib.AddTrace(err)
	}

	var outb, errb bytes.Buffer

	cmd := exec.CommandContext(ctx, "bash", "-c", "gcc main.cpp -o main")
	cmd.Dir = path
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err = cmd.Run()

	return &ExecutionResult{
		Stdout:         outb.String(),
		Stderr:         errb.String(),
		Error:          errorlib.ToString(err),
		IsCompileError: err != nil,
	}, nil
}

func (ex *CPPExecutor) Run(ctx context.Context) (*ExecutionResult, error) {
	var outb, errb bytes.Buffer
	var err error

	command := fmt.Sprintf("ulimit -Sv %v; %v", ex.MemoryLimit, filepath.Join(ex.buildPath, "main"))
	cmd := exec.CommandContext(ctx, "bash", "-c", command)
	if len(ex.Stdin) > 0 {
		cmd.Stdin = strings.NewReader(ex.Stdin)
	}
	cmd.Dir = ex.buildPath
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err = cmd.Start()
	if err != nil {
		return nil, errorlib.AddTrace(err)
	}

	done := make(chan error)
	defer func() {
		close(done)
	}()

	go func() {
		done <- cmd.Wait()
	}()

	isTle := false

	select {
	case err = <-done:
		//
	case <-time.After(time.Duration(ex.TimeLimit) * time.Millisecond):
		cmd.Process.Kill()
		isTle = true
	}

	return &ExecutionResult{
		Stdout:              outb.String(),
		Stderr:              errb.String(),
		IsTimeLimitExceeded: isTle,
		Error:               errorlib.ToString(err),
	}, nil
}

func (ex *CPPExecutor) Clean(ctx context.Context) error {
	err := os.RemoveAll(ex.buildPath)
	if err != nil {
		return errorlib.AddTrace(err)
	}
	return nil
}
