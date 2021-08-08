// generated by Textmapper; DO NOT EDIT

package demo

import (
	"fmt"
)

var tmNonterminals = [...]string{
	"Command_optlist",
	"File",
	"FileHeader",
	"VersionNum",
	"SaveNum",
	"ScreenWidth",
	"ScreenHeight",
	"Command",
	"MsgTypeEnum",
	"GameTickProgress",
	"CommandData",
	"MessageType",
	"WParam",
	"LParam",
}

func symbolName(sym int32) string {
	if sym < int32(NumTokens) {
		return Token(sym).String()
	}
	if i := int(sym) - int(NumTokens); i < len(tmNonterminals) {
		return tmNonterminals[i]
	}
	return fmt.Sprintf("nonterminal(%d)", sym)
}

var tmAction = []int32{
	-1, 4, 1, -1, -3, -1, 10, 0, -1, 5, -1, -1, -1, 11, -1, 6, -1, 9, -1, -1, 13,
	-1, -1, 7, -1, 8, -1, 3, 14, -1, -1, 15, 12, -1, -2,
}

var tmLalr = []int32{
	3, -1, 0, 2, -1, -2,
}

var tmGoto = []int32{
	0, 2, 2, 8, 24, 26, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
	68,
}

var tmFromTo = []int8{
	33, 34, 14, 17, 21, 25, 24, 27, 0, 1, 4, 6, 5, 9, 12, 15, 18, 20, 19, 23, 26,
	28, 30, 31, 11, 13, 3, 5, 8, 11, 10, 12, 14, 18, 16, 19, 22, 26, 29, 30, 2,
	4, 0, 33, 0, 2, 0, 3, 5, 10, 12, 16, 19, 24, 4, 7, 4, 8, 11, 14, 18, 21, 18,
	22, 26, 29, 30, 32,
}

var tmRuleLen = []int8{
	2, 0, 2, 8, 1, 1, 1, 1, 6, 4, 1, 1, 5, 1, 1, 1,
}

var tmRuleSymbol = []int32{
	6, 6, 7, 8, 9, 10, 11, 12, 13, 13, 14, 15, 16, 17, 18, 19,
}

var tmRuleType = [...]NodeType{
	0,                // Command_optlist : Command_optlist Command
	0,                // Command_optlist :
	File,             // File : FileHeader Command_optlist
	FileHeader,       // FileHeader : VersionNum ',' SaveNum ',' ScreenWidth ',' ScreenHeight _new_line
	0,                // VersionNum : int_lit_tok
	0,                // SaveNum : int_lit_tok
	0,                // ScreenWidth : int_lit_tok
	0,                // ScreenHeight : int_lit_tok
	Command,          // Command : MsgTypeEnum ',' GameTickProgress ',' CommandData _new_line
	Command,          // Command : MsgTypeEnum ',' GameTickProgress _new_line
	MsgTypeEnum,      // MsgTypeEnum : int_lit_tok
	GameTickProgress, // GameTickProgress : float_lit_tok
	CommandData,      // CommandData : MessageType ',' WParam ',' LParam
	MessageType,      // MessageType : int_lit_tok
	WParam,           // WParam : int_lit_tok
	LParam,           // LParam : int_lit_tok
}