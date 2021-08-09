package demo

import (
	"fmt"
	"strconv"

	"github.com/mewspring/demo/dmo/ast"
)

// translate translates the given AST file into a corresponding demo file.
func translate(root *ast.File) *File {
	file := &File{
		Hdr: translateFileHeader(root.FileHeader()),
	}
	for _, astCommand := range root.Commands() {
		command := translateCommand(astCommand)
		file.Commands = append(file.Commands, command)
	}
	return file
}

// translateFileHeader translates the given AST file header into a corresponding
// demo file header.
func translateFileHeader(astHdr ast.FileHeader) FileHeader {
	return FileHeader{
		VersionNum:   intLit(astHdr.VersionNum()),
		SaveNum:      intLit(astHdr.SaveNum()),
		ScreenWidth:  intLit(astHdr.ScreenWidth()),
		ScreenHeight: intLit(astHdr.ScreenHeight()),
	}
}

// translateCommand translates the given AST command into a corresponding demo
// command.
func translateCommand(astCommand ast.Command) Command {
	commandType := commandTypeFromString(astCommand.CommandType().Text())
	switch commandType {
	case CommandTypeGameTick:
		return GameTickCommand{
			CommandType:      CommandTypeGameTick,
			GameTickProgress: floatLit(astCommand.GameTickProgress()),
		}
	case CommandTypeRendering:
		return RenderingCommand{
			CommandType:      CommandTypeRendering,
			GameTickProgress: floatLit(astCommand.GameTickProgress()),
		}
	case CommandTypeEvent:
		eventData, ok := astCommand.EventData()
		if !ok {
			panic(fmt.Errorf("missing event data for event command"))
		}
		return EventCommand{
			CommandType:      CommandTypeEvent,
			GameTickProgress: floatLit(astCommand.GameTickProgress()),
			DvlEventType:     dvlEventTypeFromString(eventData.EventType().Text()),
			WParam:           uint32(intLit(eventData.WParam())),
			LParam:           uint32(intLit(eventData.LParam())),
		}
	default:
		panic(fmt.Errorf("support for command type %d not yet implemented", commandType))
	}
}

// CommandType specifies the type of a demo command.
type CommandType uint8

// Command types.
const (
	CommandTypeGameTick  = 0
	CommandTypeRendering = 1
	CommandTypeEvent     = 2
)

// commandTypeFromString converts the given string to the corresponding command
// type enum.
func commandTypeFromString(s string) CommandType {
	switch s {
	case "0":
		return CommandTypeGameTick
	case "1":
		return CommandTypeRendering
	case "2":
		return CommandTypeEvent
	default:
		panic(fmt.Errorf("support for command type %q not yet implemented", s))
	}
}

// dvlEventTypeFromString converts the given string to the corresponding event
// type enum.
func dvlEventTypeFromString(s string) DvlEventType {
	x, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		panic(fmt.Errorf("unable to parse event type %q; %+v", s, err))
	}
	return DvlEventType(x)
}

// intLit converts the given integer literal to the corresponding integer.
func intLit(n ast.IntLit) int {
	s := n.Text()
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Errorf("unable to parse integer literal %q; %+v", s, err))
	}
	return x
}

// floatLit converts the given floating-point literal to the corresponding
// float.
func floatLit(n ast.FloatLit) float64 {
	s := n.Text()
	x1, err := strconv.ParseFloat(s, 64)
	if err != nil {
		x2, err2 := strconv.Atoi(s)
		if err2 != nil {
			panic(fmt.Errorf("unable to parse floating-point literal %q; %+v; %+v", s, err, err2))
		}
		return float64(x2)
	}
	return x1
}
