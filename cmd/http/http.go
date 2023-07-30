package http

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/moskitos80/gonc/cmd"
)

func init() {
	cmd.Register("http", New)
}

func New(out, errout io.Writer) cmd.Cmd {
	return &command{
		out:    out,
		errout: errout,
	}
}

var methMap = map[string]string{
	"GET":  http.MethodGet,
	"POST": http.MethodPost,
	"HEAD": http.MethodHead,
}

type command struct {
	out    io.Writer
	errout io.Writer

	verb string
	body string
	url  string
}

var ErrIncorrectCType = errors.New("incorrect content-type provided")
var ErrURLNotProvided = errors.New("the mandatory argument URL not provided")
var ErrIncorrectMethod = errors.New("incorrect http method provided")

func (c *command) Run(args []string) error {

	if err := c.parseArgs(args); err != nil {
		return err
	}

	fmt.Fprintf(c.out, "%#v\n", c)

	return nil
}

func (c *command) parseArgs(args []string) error {
	ctype := ""
	verb := ""
	fs := flag.NewFlagSet("http", flag.ContinueOnError)
	fs.StringVar(&verb, "verb", http.MethodGet, "request method [GET|POST|HEAD]")
	fs.StringVar(&c.body, "body", "", "request body depends content-type")
	fs.StringVar(&ctype, "type", "plain", "request body's content-type e.g. json | urlencoded | multipart | plain")
	fs.SetOutput(c.errout)
	fs.Usage = func() {
		fmt.Fprint(c.errout, `
usage:
	http [options] URL
`)
		fs.PrintDefaults()
		fmt.Fprintln(c.errout)
	}

	if err := fs.Parse(args); err != nil {
		return err
	}

	if fs.NArg() < 1 {
		return ErrURLNotProvided
	}

	c.url = fs.Arg(0)
	c.verb = strings.ToUpper(verb)

	var ok bool
	if _, ok = methMap[c.verb]; !ok {
		return ErrIncorrectMethod
	}

	return nil
}
