package main

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
//    Rendering
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
	EventType EventType `json:"event_type"`
	// Event parameters.
	WParam int32 `json:"wparam"`
	LParam int32 `json:"lparam"`
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
