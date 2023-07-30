package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/moskitos80/gonc/cmd"
	_ "github.com/moskitos80/gonc/cmd/grpc"
	_ "github.com/moskitos80/gonc/cmd/http"
)

var errCmdNotProvided = errors.New("command not provided")

func main() {

	c, err := dispatchCommand(os.Stdout, os.Stderr, os.Args)

	if err != nil {
		if !errors.Is(err, flag.ErrHelp) {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}

	err = c.Run(os.Args[2:])
	if err != nil {
		if !errors.Is(err, flag.ErrHelp) {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
}

var helpTpl string = `usage:
	%s sub-command [options]

sub-command:
	http - A HTTP client
	grpc - A gRPC client
`

func dispatchCommand(out, errout io.Writer, args []string) (cmd.Cmd, error) {
	appName := filepath.Base(args[0])

	fs := flag.NewFlagSet(appName, flag.ContinueOnError)
	fs.SetOutput(errout)
	fs.Usage = func() {
		fmt.Fprintf(errout, helpTpl, appName)
		fs.PrintDefaults()
		fmt.Fprintln(errout)
	}

	if err := fs.Parse(args[1:]); err != nil {
		return nil, err
	}

	if fs.NArg() == 0 {
		return nil, errCmdNotProvided
	}

	return cmd.Get(args[1], out, errout)
}
