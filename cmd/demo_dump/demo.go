package main

// File holds the contents of a demo file.
type File struct {
	// File header.
	Hdr FileHeader
	// Commands
	Commands []Command
}

// FileHeader specifies the file header of a demo file.
type FileHeader struct {
	// File format version number.
	VersionNum int
	// Save number (e.g. "single_#.sv").
	SaveNum int
	// Screen dimensions.
	ScreenWidth, ScreenHeight int
}

// Command is a demo command.
//
// Command has one of the following underlying types:
//
//    Rendering
type Command interface {
	// isCommand ensures that only commands may be assigned to the Command
	// interface.
	isCommand()
}

// GameTickCommand specifies a game tick command.
type GameTickCommand struct {
	// Progress to next game tick in range [0.0, 1.0].
	GameTickProgress float64
}

// RenderingCommand specifies a rendering command.
type RenderingCommand struct {
	// Progress to next game tick in range [0.0, 1.0].
	GameTickProgress float64
}

// EventCommand specifies an input event command.
type EventCommand struct {
	// Progress to next game tick in range [0.0, 1.0].
	GameTickProgress float64
	// Event type.
	EventType EventType
	// Event parameters.
	WParam, LParam int32
}

// EventType specifies the event type of a message command.
type EventType uint32

// isCommand ensures that only commands may be assigned to the Command
// interface.
func (GameTickCommand) isCommand() {}

// isCommand ensures that only commands may be assigned to the Command
// interface.
func (RenderingCommand) isCommand() {}

// isCommand ensures that only commands may be assigned to the Command
// interface.
func (EventCommand) isCommand() {}
