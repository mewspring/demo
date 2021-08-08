// generated by Textmapper; DO NOT EDIT

package demo

import (
	"fmt"
)

type NodeType int

type Listener func(t NodeType, offset, endoffset int)

const (
	NoType NodeType = iota
	File            // FileHeader Commands=(Command)*
	IntLit
	FloatLit
	FileHeader // VersionNum=IntLit SaveNum=IntLit ScreenWidth=IntLit ScreenHeight=IntLit
	Command    // CommandType=CommandTypeEnum GameTickProgress=FloatLit EventData?
	CommandTypeEnum
	EventData // EventType WParam=IntLit LParam=IntLit
	EventType
	NodeTypeMax
)

var nodeTypeStr = [...]string{
	"NONE",
	"File",
	"IntLit",
	"FloatLit",
	"FileHeader",
	"Command",
	"CommandTypeEnum",
	"EventData",
	"EventType",
}

func (t NodeType) String() string {
	if t >= 0 && int(t) < len(nodeTypeStr) {
		return nodeTypeStr[t]
	}
	return fmt.Sprintf("node(%d)", t)
}
