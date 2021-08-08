// generated by Textmapper; DO NOT EDIT

package ast

import (
	"fmt"

	"github.com/mewspring/demo"
)

func ToDemoNode(n *Node) DemoNode {
	switch n.Type() {
	case demo.Command:
		return &Command{n}
	case demo.CommandData:
		return &CommandData{n}
	case demo.File:
		return &File{n}
	case demo.FileHeader:
		return &FileHeader{n}
	case demo.GameTickProgress:
		return &GameTickProgress{n}
	case demo.LParam:
		return &LParam{n}
	case demo.MessageType:
		return &MessageType{n}
	case demo.MsgTypeEnum:
		return &MsgTypeEnum{n}
	case demo.WParam:
		return &WParam{n}
	case demo.NoType:
		return nilInstance
	}
	panic(fmt.Errorf("ast: unknown node type %v", n.Type()))
	return nil
}