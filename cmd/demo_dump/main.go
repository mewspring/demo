package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mewkiz/pkg/jsonutil"
	"github.com/mewkiz/pkg/term"
	"github.com/mewspring/demo"
	"github.com/pkg/errors"
)

var (
	// dbg is a logger with the "demo_dump:" prefix which logs debug messages to
	// standard error.
	dbg = log.New(os.Stderr, term.MagentaBold("demo_dump:")+" ", 0)
	// warn is a logger with the "demo_dump:" prefix which logs warning messages
	// to standard error.
	warn = log.New(os.Stderr, term.RedBold("demo_dump:")+" ", log.Lshortfile)
)

func init() {
	if !debug {
		dbg.SetOutput(ioutil.Discard)
	}
}

// Enable debug output.
const debug = true

func main() {
	flag.Parse()
	for _, demoPath := range flag.Args() {
		if err := parseDemo(demoPath); err != nil {
			warn.Fatalf("%+v", err)
		}
	}
}

// parseDemo parses the given demo file.
func parseDemo(demoPath string) error {
	dbg.Printf("parsing %q", demoPath)
	file, err := demo.ParseFile(demoPath)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := jsonutil.Write(os.Stdout, file); err != nil {
		return errors.WithStack(err)
	}
	fmt.Fprintln(os.Stdout) // add trailing newline.
	return nil
}
