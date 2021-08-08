package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/kr/pretty"
	"github.com/mewkiz/pkg/term"
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
	file, err := ParseFile(demoPath)
	if err != nil {
		return errors.WithStack(err)
	}
	pretty.Println("file:", file)
	return nil
}
