package lexers

import (
	. "github.com/dezza/chroma" // nolint
)

// Cobol lexer.
var Cobol = Register(MustNewLexer(
	&Config{
		Name:            "COBOL",
		Aliases:         []string{"cobol"},
		Filenames:       []string{"*.cob", "*.COB", "*.cpy", "*.CPY"},
		MimeTypes:       []string{"text/x-cobol"},
		CaseInsensitive: true,
	},
	Rules{
		"root": {
			Include("comment"),
			Include("strings"),
			Include("core"),
			Include("nums"),
			{`[a-z0-9]([\w\-]*[a-z0-9]+)?`, NameVariable, nil},
			{`[ \t]+`, Text, nil},
		},
		"comment": {
			{`(^.{6}[*/].*\n|^.{6}|\*>.*\n)`, Comment, nil},
		},
		"core": {
			{`(^|(?<=[^\w\-]))(ALL\s+)?((ZEROES)|(HIGH-VALUE|LOW-VALUE|QUOTE|SPACE|ZERO)(S)?)\s*($|(?=[^\w\-]))`, NameConstant, nil},
			{Words(`(^|(?<=[^\w\-]))`, `\s*($|(?=[^\w\-]))`, `ACCEPT`, `ADD`, `ALLOCATE`, `CALL`, `CANCEL`, `CLOSE`, `COMPUTE`, `CONFIGURATION`, `CONTINUE`, `DATA`, `DELETE`, `DISPLAY`, `DIVIDE`, `DIVISION`, `ELSE`, `END`, `END-ACCEPT`, `END-ADD`, `END-CALL`, `END-COMPUTE`, `END-DELETE`, `END-DISPLAY`, `END-DIVIDE`, `END-EVALUATE`, `END-IF`, `END-MULTIPLY`, `END-OF-PAGE`, `END-PERFORM`, `END-READ`, `END-RETURN`, `END-REWRITE`, `END-SEARCH`, `END-START`, `END-STRING`, `END-SUBTRACT`, `END-UNSTRING`, `END-WRITE`, `ENVIRONMENT`, `EVALUATE`, `EXIT`, `FD`, `FILE`, `FILE-CONTROL`, `FOREVER`, `FREE`, `GENERATE`, `GO`, `GOBACK`, `IDENTIFICATION`, `IF`, `INITIALIZE`, `INITIATE`, `INPUT-OUTPUT`, `INSPECT`, `INVOKE`, `I-O-CONTROL`, `LINKAGE`, `LOCAL-STORAGE`, `MERGE`, `MOVE`, `MULTIPLY`, `OPEN`, `PERFORM`, `PROCEDURE`, `PROGRAM-ID`, `RAISE`, `READ`, `RELEASE`, `RESUME`, `RETURN`, `REWRITE`, `SCREEN`, `SD`, `SEARCH`, `SECTION`, `SET`, `SORT`, `START`, `STOP`, `STRING`, `SUBTRACT`, `SUPPRESS`, `TERMINATE`, `THEN`, `UNLOCK`, `UNSTRING`, `USE`, `VALIDATE`, `WORKING-STORAGE`, `WRITE`), KeywordReserved, nil},
			{Words(`(^|(?<=[^\w\-]))`, `\s*($|(?=[^\w\-]))`, `ACCESS`, `ADDRESS`, `ADVANCING`, `AFTER`, `ALL`, `ALPHABET`, `ALPHABETIC`, `ALPHABETIC-LOWER`, `ALPHABETIC-UPPER`, `ALPHANUMERIC`, `ALPHANUMERIC-EDITED`, `ALSO`, `ALTER`, `ALTERNATEANY`, `ARE`, `AREA`, `AREAS`, `ARGUMENT-NUMBER`, `ARGUMENT-VALUE`, `AS`, `ASCENDING`, `ASSIGN`, `AT`, `AUTO`, `AUTO-SKIP`, `AUTOMATIC`, `AUTOTERMINATE`, `BACKGROUND-COLOR`, `BASED`, `BEEP`, `BEFORE`, `BELL`, `BLANK`, `BLINK`, `BLOCK`, `BOTTOM`, `BY`, `BYTE-LENGTH`, `CHAINING`, `CHARACTER`, `CHARACTERS`, `CLASS`, `CODE`, `CODE-SET`, `COL`, `COLLATING`, `COLS`, `COLUMN`, `COLUMNS`, `COMMA`, `COMMAND-LINE`, `COMMIT`, `COMMON`, `CONSTANT`, `CONTAINS`, `CONTENT`, `CONTROL`, `CONTROLS`, `CONVERTING`, `COPY`, `CORR`, `CORRESPONDING`, `COUNT`, `CRT`, `CURRENCY`, `CURSOR`, `CYCLE`, `DATE`, `DAY`, `DAY-OF-WEEK`, `DE`, `DEBUGGING`, `DECIMAL-POINT`, `DECLARATIVES`, `DEFAULT`, `DELIMITED`, `DELIMITER`, `DEPENDING`, `DESCENDING`, `DETAIL`, `DISK`, `DOWN`, `DUPLICATES`, `DYNAMIC`, `EBCDIC`, `ENTRY`, `ENVIRONMENT-NAME`, `ENVIRONMENT-VALUE`, `EOL`, `EOP`, `EOS`, `ERASE`, `ERROR`, `ESCAPE`, `EXCEPTION`, `EXCLUSIVE`, `EXTEND`, `EXTERNAL`, `FILE-ID`, `FILLER`, `FINAL`, `FIRST`, `FIXED`, `FLOAT-LONG`, `FLOAT-SHORT`, `FOOTING`, `FOR`, `FOREGROUND-COLOR`, `FORMAT`, `FROM`, `FULL`, `FUNCTION`, `FUNCTION-ID`, `GIVING`, `GLOBAL`, `GROUP`, `HEADING`, `HIGHLIGHT`, `I-O`, `ID`, `IGNORE`, `IGNORING`, `IN`, `INDEX`, `INDEXED`, `INDICATE`, `INITIAL`, `INITIALIZED`, `INPUT`, `INTO`, `INTRINSIC`, `INVALID`, `IS`, `JUST`, `JUSTIFIED`, `KEY`, `LABEL`, `LAST`, `LEADING`, `LEFT`, `LENGTH`, `LIMIT`, `LIMITS`, `LINAGE`, `LINAGE-COUNTER`, `LINE`, `LINES`, `LOCALE`, `LOCK`, `LOWLIGHT`, `MANUAL`, `MEMORY`, `MINUS`, `MODE`, `MULTIPLE`, `NATIONAL`, `NATIONAL-EDITED`, `NATIVE`, `NEGATIVE`, `NEXT`, `NO`, `NULL`, `NULLS`, `NUMBER`, `NUMBERS`, `NUMERIC`, `NUMERIC-EDITED`, `OBJECT-COMPUTER`, `OCCURS`, `OF`, `OFF`, `OMITTED`, `ON`, `ONLY`, `OPTIONAL`, `ORDER`, `ORGANIZATION`, `OTHER`, `OUTPUT`, `OVERFLOW`, `OVERLINE`, `PACKED-DECIMAL`, `PADDING`, `PAGE`, `PARAGRAPH`, `PLUS`, `POINTER`, `POSITION`, `POSITIVE`, `PRESENT`, `PREVIOUS`, `PRINTER`, `PRINTING`, `PROCEDURE-POINTER`, `PROCEDURES`, `PROCEED`, `PROGRAM`, `PROGRAM-POINTER`, `PROMPT`, `QUOTE`, `QUOTES`, `RANDOM`, `RD`, `RECORD`, `RECORDING`, `RECORDS`, `RECURSIVE`, `REDEFINES`, `REEL`, `REFERENCE`, `RELATIVE`, `REMAINDER`, `REMOVAL`, `RENAMES`, `REPLACING`, `REPORT`, `REPORTING`, `REPORTS`, `REPOSITORY`, `REQUIRED`, `RESERVE`, `RETURNING`, `REVERSE-VIDEO`, `REWIND`, `RIGHT`, `ROLLBACK`, `ROUNDED`, `RUN`, `SAME`, `SCROLL`, `SECURE`, `SEGMENT-LIMIT`, `SELECT`, `SENTENCE`, `SEPARATE`, `SEQUENCE`, `SEQUENTIAL`, `SHARING`, `SIGN`, `SIGNED`, `SIGNED-INT`, `SIGNED-LONG`, `SIGNED-SHORT`, `SIZE`, `SORT-MERGE`, `SOURCE`, `SOURCE-COMPUTER`, `SPECIAL-NAMES`, `STANDARD`, `STANDARD-1`, `STANDARD-2`, `STATUS`, `SUM`, `SYMBOLIC`, `SYNC`, `SYNCHRONIZED`, `TALLYING`, `TAPE`, `TEST`, `THROUGH`, `THRU`, `TIME`, `TIMES`, `TO`, `TOP`, `TRAILING`, `TRANSFORM`, `TYPE`, `UNDERLINE`, `UNIT`, `UNSIGNED`, `UNSIGNED-INT`, `UNSIGNED-LONG`, `UNSIGNED-SHORT`, `UNTIL`, `UP`, `UPDATE`, `UPON`, `USAGE`, `USING`, `VALUE`, `VALUES`, `VARYING`, `WAIT`, `WHEN`, `WITH`, `WORDS`, `YYYYDDD`, `YYYYMMDD`), KeywordPseudo, nil},
			{Words(`(^|(?<=[^\w\-]))`, `\s*($|(?=[^\w\-]))`, `ACTIVE-CLASS`, `ALIGNED`, `ANYCASE`, `ARITHMETIC`, `ATTRIBUTE`, `B-AND`, `B-NOT`, `B-OR`, `B-XOR`, `BIT`, `BOOLEAN`, `CD`, `CENTER`, `CF`, `CH`, `CHAIN`, `CLASS-ID`, `CLASSIFICATION`, `COMMUNICATION`, `CONDITION`, `DATA-POINTER`, `DESTINATION`, `DISABLE`, `EC`, `EGI`, `EMI`, `ENABLE`, `END-RECEIVE`, `ENTRY-CONVENTION`, `EO`, `ESI`, `EXCEPTION-OBJECT`, `EXPANDS`, `FACTORY`, `FLOAT-BINARY-16`, `FLOAT-BINARY-34`, `FLOAT-BINARY-7`, `FLOAT-DECIMAL-16`, `FLOAT-DECIMAL-34`, `FLOAT-EXTENDED`, `FORMAT`, `FUNCTION-POINTER`, `GET`, `GROUP-USAGE`, `IMPLEMENTS`, `INFINITY`, `INHERITS`, `INTERFACE`, `INTERFACE-ID`, `INVOKE`, `LC_ALL`, `LC_COLLATE`, `LC_CTYPE`, `LC_MESSAGES`, `LC_MONETARY`, `LC_NUMERIC`, `LC_TIME`, `LINE-COUNTER`, `MESSAGE`, `METHOD`, `METHOD-ID`, `NESTED`, `NONE`, `NORMAL`, `OBJECT`, `OBJECT-REFERENCE`, `OPTIONS`, `OVERRIDE`, `PAGE-COUNTER`, `PF`, `PH`, `PROPERTY`, `PROTOTYPE`, `PURGE`, `QUEUE`, `RAISE`, `RAISING`, `RECEIVE`, `RELATION`, `REPLACE`, `REPRESENTS-NOT-A-NUMBER`, `RESET`, `RESUME`, `RETRY`, `RF`, `RH`, `SECONDS`, `SEGMENT`, `SELF`, `SEND`, `SOURCES`, `STATEMENT`, `STEP`, `STRONG`, `SUB-QUEUE-1`, `SUB-QUEUE-2`, `SUB-QUEUE-3`, `SUPER`, `SYMBOL`, `SYSTEM-DEFAULT`, `TABLE`, `TERMINAL`, `TEXT`, `TYPEDEF`, `UCS-4`, `UNIVERSAL`, `USER-DEFAULT`, `UTF-16`, `UTF-8`, `VAL-STATUS`, `VALID`, `VALIDATE`, `VALIDATE-STATUS`), Error, nil},
			{`(^|(?<=[^\w\-]))(PIC\s+.+?(?=(\s|\.\s))|PICTURE\s+.+?(?=(\s|\.\s))|(COMPUTATIONAL)(-[1-5X])?|(COMP)(-[1-5X])?|BINARY-C-LONG|BINARY-CHAR|BINARY-DOUBLE|BINARY-LONG|BINARY-SHORT|BINARY)\s*($|(?=[^\w\-]))`, KeywordType, nil},
			{`(\*\*|\*|\+|-|/|<=|>=|<|>|==|/=|=)`, Operator, nil},
			{`([(),;:&%.])`, Punctuation, nil},
			{`(^|(?<=[^\w\-]))(ABS|ACOS|ANNUITY|ASIN|ATAN|BYTE-LENGTH|CHAR|COMBINED-DATETIME|CONCATENATE|COS|CURRENT-DATE|DATE-OF-INTEGER|DATE-TO-YYYYMMDD|DAY-OF-INTEGER|DAY-TO-YYYYDDD|EXCEPTION-(?:FILE|LOCATION|STATEMENT|STATUS)|EXP10|EXP|E|FACTORIAL|FRACTION-PART|INTEGER-OF-(?:DATE|DAY|PART)|INTEGER|LENGTH|LOCALE-(?:DATE|TIME(?:-FROM-SECONDS)?)|LOG(?:10)?|LOWER-CASE|MAX|MEAN|MEDIAN|MIDRANGE|MIN|MOD|NUMVAL(?:-C)?|ORD(?:-MAX|-MIN)?|PI|PRESENT-VALUE|RANDOM|RANGE|REM|REVERSE|SECONDS-FROM-FORMATTED-TIME|SECONDS-PAST-MIDNIGHT|SIGN|SIN|SQRT|STANDARD-DEVIATION|STORED-CHAR-LENGTH|SUBSTITUTE(?:-CASE)?|SUM|TAN|TEST-DATE-YYYYMMDD|TEST-DAY-YYYYDDD|TRIM|UPPER-CASE|VARIANCE|WHEN-COMPILED|YEAR-TO-YYYY)\s*($|(?=[^\w\-]))`, NameFunction, nil},
			{`(^|(?<=[^\w\-]))(true|false)\s*($|(?=[^\w\-]))`, NameBuiltin, nil},
			{`(^|(?<=[^\w\-]))(equal|equals|ne|lt|le|gt|ge|greater|less|than|not|and|or)\s*($|(?=[^\w\-]))`, OperatorWord, nil},
		},
		"strings": {
			{`"[^"\n]*("|\n)`, LiteralStringDouble, nil},
			{`'[^'\n]*('|\n)`, LiteralStringSingle, nil},
		},
		"nums": {
			{`\d+(\s*|\.$|$)`, LiteralNumberInteger, nil},
			{`[+-]?\d*\.\d+(E[-+]?\d+)?`, LiteralNumberFloat, nil},
			{`[+-]?\d+\.\d*(E[-+]?\d+)?`, LiteralNumberFloat, nil},
		},
	},
))
