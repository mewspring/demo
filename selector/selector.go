// generated by Textmapper; DO NOT EDIT

package selector

import (
	"github.com/mewspring/demo"
)

type Selector func(nt demo.NodeType) bool

var (
	Any             = func(t demo.NodeType) bool { return true }
	Command         = func(t demo.NodeType) bool { return t == demo.Command }
	CommandTypeEnum = func(t demo.NodeType) bool { return t == demo.CommandTypeEnum }
	EventData       = func(t demo.NodeType) bool { return t == demo.EventData }
	EventType       = func(t demo.NodeType) bool { return t == demo.EventType }
	File            = func(t demo.NodeType) bool { return t == demo.File }
	FileHeader      = func(t demo.NodeType) bool { return t == demo.FileHeader }
	FloatLit        = func(t demo.NodeType) bool { return t == demo.FloatLit }
	IntLit          = func(t demo.NodeType) bool { return t == demo.IntLit }
)

func OneOf(types ...demo.NodeType) Selector {
	if len(types) == 0 {
		return func(demo.NodeType) bool { return false }
	}
	const bits = 32
	max := 1
	for _, t := range types {
		if int(t) > max {
			max = int(t)
		}
	}
	size := (max + bits) / bits
	bitarr := make([]uint32, size)
	for _, t := range types {
		bitarr[uint(t)/bits] |= 1 << (uint(t) % bits)
	}
	return func(t demo.NodeType) bool {
		i := uint(t) / bits
		return int(i) < len(bitarr) && bitarr[i]&(1<<(uint(t)%bits)) != 0
	}
}
