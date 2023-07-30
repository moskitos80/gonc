package cmd

import (
	"errors"
	"io"
)

type Cmd interface {
	Run(args []string) error
}

type CmdConstructFn func(out, outerr io.Writer) Cmd

var ErrCmdNotSupported = errors.New("command not support")

var commands = map[string]CmdConstructFn{}

func Register(name string, constructor CmdConstructFn) {
	commands[name] = constructor
}

func Get(name string, out, outerr io.Writer) (Cmd, error) {
	if command, ok := commands[name]; !ok {
		return nil, ErrCmdNotSupported
	} else {
		return command(out, outerr), nil
	}
}
