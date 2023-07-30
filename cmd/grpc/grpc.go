package grpc

import (
	"fmt"
	"io"

	"github.com/moskitos80/gonc/cmd"
)

func init() {
	cmd.Register("grpc", New)
}

func New(out, errout io.Writer) cmd.Cmd {
	return &command{
		out:    out,
		errout: errout,
	}
}

type command struct {
	out    io.Writer
	errout io.Writer
}

func (c *command) Run(args []string) error {
	fmt.Fprintln(c.out, "Result of grpc cmd run")
	return nil
}
