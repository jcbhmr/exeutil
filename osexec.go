package exeutil

import (
	"bytes"
	"context"
	"errors"
	"io"
)

type MemCmd struct {
	Args      []string
	Env       []string
	Dir       string
	Stdin     io.Reader
	Stdout    io.Writer
	Stderr    io.Writer
	Err       error
	Cancel    func() error
	WaitDelay func() error
	memCmd
}

func (m *MemCmd) CombinedOutput() ([]byte, error) {
	if m.Stdout != nil {
		return nil, errors.New("exeutil: Stdout already set")
	}
	if m.Stderr != nil {
		return nil, errors.New("exeutil: Stderr already set")
	}
	var b bytes.Buffer
	m.Stdout = &b
	m.Stderr = &b
	err := m.Run()
	return b.Bytes(), err
}

func (m *MemCmd) Environ() []string {
	return m.Env
}

func (m *MemCmd) Output() ([]byte, error) {
	if m.Stdout != nil {
		return nil, errors.New("exec: Stdout already set")
	}
	var stdout bytes.Buffer
	m.Stdout = &stdout

	captureErr := m.Stderr == nil
	if captureErr {
		m.Stderr = &prefixSuffixSaver{N: 32 << 10}
	}

	err := m.Run()
	if err != nil && captureErr {
		if ee, ok := err.(*ExitError); ok {
			ee.Stderr = m.Stderr.(*prefixSuffixSaver).Bytes()
		}
	}
	return stdout.Bytes(), err
}

func MemCommand(bytes []byte, argv0 string, args ...string) *MemCmd {
	panic("unimplemented")
}

func MemCommandContext(ctx context.Context, bytes []byte, argv0 string, args ...string) *MemCmd {
	panic("unimplemented")
}
