language demo(go);

lang = "demo"
package = "github.com/mewspring/demo"
eventBased = true
eventFields = true
eventAST = true

# ### [ Lexical part ] #########################################################

:: lexer

_decimal_digit = /[0-9]/

_new_line : /[\n]/

# === [ Integer literals ] =====================================================

int_lit_tok : /[-]?[0-9]+/

_decimals = /{_decimal_digit}+/

# === [ Floating-point literals ] ==============================================

float_lit_tok : /{_frac_lit}/

_frac_lit = /{_sign}?{_decimals}[\.]{_decimal_digit}*/

_sign = /[+-]/

# === [ Tokens ] ===============================================================

',' : /[,]/

# ### [ Syntax part ] ##########################################################

:: parser

%input File;

# === [ Demo file ] ============================================================

File -> File
	: FileHeader Commands=Command*
;

# --- [ File header ] ----------------------------------------------------------

FileHeader -> FileHeader
	: VersionNum ',' SaveNum ',' ScreenWidth ',' ScreenHeight _new_line
;

# File format version.
VersionNum
	: int_lit_tok
;

# Save number (e.g. "single_#.sv")
SaveNum
	: int_lit_tok
;

# Screen dimensions.
ScreenWidth
	: int_lit_tok
;

ScreenHeight
	: int_lit_tok
;

# --- [ Commands ] -------------------------------------------------------------

Command -> Command
	: MsgTypeEnum ',' GameTickProgress CommandData=(',' CommandData)? _new_line
;

MsgTypeEnum -> MsgTypeEnum
	: int_lit_tok
;

# Progress to next game tick in range [0.0, 1.0].
GameTickProgress -> GameTickProgress
	: float_lit_tok
;

CommandData -> CommandData
	: MessageType ',' WParam ',' LParam
;

MessageType -> MessageType
	: int_lit_tok # uint32
;

WParam -> WParam
	: int_lit_tok # int32
;

LParam -> LParam
	: int_lit_tok # int32
;
