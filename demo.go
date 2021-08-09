// Package demo provides access to demo files (recording of DevilutionX input
// events).
package demo

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/mewkiz/pkg/term"
)

var (
	// dbg is a logger with the "demo:" prefix which logs debug messages to
	// standard error.
	dbg = log.New(os.Stderr, term.MagentaBold("demo:")+" ", 0)
	// warn is a logger with the "demo:" prefix which logs warning messages
	// to standard error.
	warn = log.New(os.Stderr, term.RedBold("demo:")+" ", log.Lshortfile)
)

func init() {
	if !debug {
		dbg.SetOutput(ioutil.Discard)
	}
}

// Enable debug output.
const debug = true

// File holds the contents of a demo file.
type File struct {
	// File header.
	Hdr FileHeader `json:"file_hdr"`
	// Commands
	Commands []Command `json:"commands"`
}

// FileHeader specifies the file header of a demo file.
type FileHeader struct {
	// File format version number.
	VersionNum int `json:"version_num"`
	// Save number (e.g. "single_#.sv").
	SaveNum int `json:"save_num"`
	// Screen dimensions.
	ScreenWidth  int `json:"screen_width"`
	ScreenHeight int `json:"screen_height"`
}

// Command is a demo command.
//
// Command has one of the following underlying types:
//
//    GameTickCommand
//    RenderingCommand
//    EventCommand
type Command interface {
	// isCommand ensures that only commands may be assigned to the Command
	// interface.
	isCommand()
}

// GameTickCommand specifies a game tick command.
type GameTickCommand struct {
	// Command type; always CommandTypeGameTick (used for serialization).
	CommandType CommandType `json:"command_type"`
	// Progress to next game tick in range [0.0, 1.0].
	GameTickProgress float64 `json:"game_tick_progress"`
}

// RenderingCommand specifies a rendering command.
type RenderingCommand struct {
	// Command type; always CommandTypeRendering (used for serialization).
	CommandType CommandType `json:"command_type"`
	// Progress to next game tick in range [0.0, 1.0].
	GameTickProgress float64 `json:"game_tick_progress"`
}

// EventCommand specifies an input event command.
type EventCommand struct {
	// Command type; always CommandTypeEvent (used for serialization).
	CommandType CommandType `json:"command_type"`
	// Progress to next game tick in range [0.0, 1.0].
	GameTickProgress float64 `json:"game_tick_progress"`
	// Event type.
	DvlEventType DvlEventType `json:"event_type"`
	// Event parameters.
	WParam uint32 `json:"wparam"`
	LParam uint32 `json:"lparam"`
}

// DvlEventType specifies the event type of a message command.
type DvlEventType uint32

const (
	DvlEventTypeMouseMove            DvlEventType = 0x0200
	DvlEventTypeMouseButtonLeftDown  DvlEventType = 0x0201
	DvlEventTypeMouseButtonLeftUp    DvlEventType = 0x0202
	DvlEventTypeMouseButtonRightDown DvlEventType = 0x0204
	DvlEventTypeMouseButtonRightUp   DvlEventType = 0x0205
	DvlEventTypeKeyDown              DvlEventType = 0x0100
	DvlEventTypeKeyUp                DvlEventType = 0x0101
	DvlEventTypeChar                 DvlEventType = 0x0102
	DvlEventTypeQuit                 DvlEventType = 0x0012
	DvlEventTypeCaptureChanged       DvlEventType = 0x0215
	DvlEventTypePaint                DvlEventType = 0x000F
	DvlEventTypeQueryEndSession      DvlEventType = 0x0011

	// 0x402
	//
	// WM_DIABNEXTLVL
	DvlEventTypeTrigMsgNextDLvl = DvlEventTypeTrigMsgUser + 2 // Down to level NN
	// 0x403
	//
	// WM_DIABPREVLVL
	DvlEventTypeTrigMsgPrevDLvl = DvlEventTypeTrigMsgUser + 3 // Up to level NN
	// 0x404
	//
	// WM_DIABRTNLVL
	DvlEventTypeTrigMsgReturnDLvl = DvlEventTypeTrigMsgUser + 4 // Return to entrance of set level
	// 0x405
	//
	// WM_DIABSETLVL
	DvlEventTypeTrigMsgSetDLvl = DvlEventTypeTrigMsgUser + 5 // Enter set level
	// 0x406
	//
	// WM_DIABWARPLVL
	DvlEventTypeTrigMsgTownPortal = DvlEventTypeTrigMsgUser + 6 // Town Portal from PLAYER
	// 0x407
	//
	// WM_DIABTOWNWARP
	DvlEventTypeTrigMsgTownWarpDown = DvlEventTypeTrigMsgUser + 7 // Down to DTYPE
	// 0x408
	//
	// WM_DIABTWARPUP
	DvlEventTypeTrigMsgTownWarpUp = DvlEventTypeTrigMsgUser + 8 // Up to town
	// 0x409
	//
	// WM_DIABRETOWN
	DvlEventTypeTrigMsgRestartInTown = DvlEventTypeTrigMsgUser + 9 // Restart In Town
	// 0x40A
	//
	// WM_DIABNEWGAME
	DvlEventTypeTrigMsgNewGame = DvlEventTypeTrigMsgUser + 10 // New Game
	// 0x40B
	//
	// WM_DIABLOADGAME
	DvlEventTypeTrigMsgLoadGame = DvlEventTypeTrigMsgUser + 11 // Load Game

	DvlEventTypeTrigMsgUser DvlEventType = 0x0400 // WM_USER
)

// isCommand ensures that only commands may be assigned to the Command
// interface.
func (GameTickCommand) isCommand() {}

// isCommand ensures that only commands may be assigned to the Command
// interface.
func (RenderingCommand) isCommand() {}

// isCommand ensures that only commands may be assigned to the Command
// interface.
func (EventCommand) isCommand() {}
