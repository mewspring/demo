// generated by Textmapper; DO NOT EDIT

package ast

import (
	"github.com/mewspring/demo/dmo/selector"
)

// Interfaces.

type DemoNode interface {
	DemoNode() *Node
}

type NilNode struct{}

var nilInstance = &NilNode{}

// All types implement DemoNode.
func (n Command) DemoNode() *Node         { return n.Node }
func (n CommandTypeEnum) DemoNode() *Node { return n.Node }
func (n EventData) DemoNode() *Node       { return n.Node }
func (n EventType) DemoNode() *Node       { return n.Node }
func (n File) DemoNode() *Node            { return n.Node }
func (n FileHeader) DemoNode() *Node      { return n.Node }
func (n FloatLit) DemoNode() *Node        { return n.Node }
func (n IntLit) DemoNode() *Node          { return n.Node }
func (NilNode) DemoNode() *Node           { return nil }

// Types.

type Command struct {
	*Node
}

func (n Command) CommandType() CommandTypeEnum {
	return CommandTypeEnum{n.Child(selector.CommandTypeEnum)}
}

func (n Command) GameTickProgress() FloatLit {
	return FloatLit{n.Child(selector.FloatLit)}
}

func (n Command) EventData() (EventData, bool) {
	field := EventData{n.Child(selector.EventData)}
	return field, field.IsValid()
}

type CommandTypeEnum struct {
	*Node
}

type EventData struct {
	*Node
}

func (n EventData) EventType() EventType {
	return EventType{n.Child(selector.EventType)}
}

func (n EventData) WParam() IntLit {
	return IntLit{n.Child(selector.IntLit)}
}

func (n EventData) LParam() IntLit {
	return IntLit{n.Child(selector.IntLit).Next(selector.IntLit)}
}

type EventType struct {
	*Node
}

type File struct {
	*Node
}

func (n File) FileHeader() FileHeader {
	return FileHeader{n.Child(selector.FileHeader)}
}

func (n File) Commands() []Command {
	nodes := n.Children(selector.Command)
	var ret = make([]Command, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, Command{node})
	}
	return ret
}

type FileHeader struct {
	*Node
}

func (n FileHeader) VersionNum() IntLit {
	return IntLit{n.Child(selector.IntLit)}
}

func (n FileHeader) SaveNum() IntLit {
	return IntLit{n.Child(selector.IntLit).Next(selector.IntLit)}
}

func (n FileHeader) ScreenWidth() IntLit {
	return IntLit{n.Child(selector.IntLit).Next(selector.IntLit).Next(selector.IntLit)}
}

func (n FileHeader) ScreenHeight() IntLit {
	return IntLit{n.Child(selector.IntLit).Next(selector.IntLit).Next(selector.IntLit).Next(selector.IntLit)}
}

type FloatLit struct {
	*Node
}

type IntLit struct {
	*Node
}