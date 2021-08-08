package main

import (
	"io"
	"io/ioutil"
	"time"

	"github.com/mewspring/demo/ast"
	"github.com/pkg/errors"
)

// ParseFile parses the given demo file.
func ParseFile(path string) (*File, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ParseBytes(path, buf)
}

// Parse parses the given demo file, reading from r. An optional path to the
// source file may be specified for error reporting.
func Parse(path string, r io.Reader) (*File, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ParseBytes(path, buf)
}

// ParseBytes parses the given demo file, reading from b. An optional path to
// the source file may be specified for error reporting.
func ParseBytes(path string, b []byte) (*File, error) {
	content := string(b)
	return ParseString(path, content)
}

// ParseString parses the given demo file, reading from content. An optional
// path to the source file may be specified for error reporting.
func ParseString(path, content string) (*File, error) {
	parseStart := time.Now()
	tree, err := ast.Parse(path, content)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse %q into an AST", path)
	}
	dbg.Println("parsing into AST took:", time.Since(parseStart))
	root := ast.ToDemoNode(tree.Root())
	return translate(root.(*ast.File)), nil
}
