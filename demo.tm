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

# --- [ Literals ] -------------------------------------------------------------

IntLit -> IntLit
	: int_lit_tok
;

FloatLit -> FloatLit
	: float_lit_tok
	| int_lit_tok     # to support short form float for 0 and 1.
;

# --- [ File header ] ----------------------------------------------------------

# FileHeader
#    // File format version.
#    VersionNum
#    // Save number (e.g. "single_#.sv").
#    SaveNum
#    // Screen dimensions.
#    ScreenWidth, ScreenHeight
FileHeader -> FileHeader
	: VersionNum=IntLit ',' SaveNum=IntLit ',' ScreenWidth=IntLit ',' ScreenHeight=IntLit _new_line
;

# --- [ Commands ] -------------------------------------------------------------

# Command
#    // Progress to next game tick in range [0.0, 1.0].
#    GameTickProgress
Command -> Command
	: CommandType=CommandTypeEnum ',' GameTickProgress=FloatLit EventData=(',' EventData)? _new_line
;

CommandTypeEnum -> CommandTypeEnum
	: int_lit_tok
;

EventData -> EventData
	: EventType=EventType ',' WParam=IntLit ',' LParam=IntLit
;

EventType -> EventType
	: int_lit_tok # uint32
;
