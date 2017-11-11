package lexers

import (
	. "github.com/dezza/chroma" // nolint
)

// Crystal lexer.
var Crystal = Register(MustNewLexer(
	&Config{
		Name:      "Crystal",
		Aliases:   []string{"cr", "crystal"},
		Filenames: []string{"*.cr"},
		MimeTypes: []string{"text/x-crystal"},
		DotAll:    true,
	},
	Rules{
		"root": {
			{`#.*?$`, CommentSingle, nil},
			{Words(``, `\b`, `abstract`, `asm`, `as`, `begin`, `break`, `case`, `do`, `else`, `elsif`, `end`, `ensure`, `extend`, `ifdef`, `if`, `include`, `instance_sizeof`, `next`, `of`, `pointerof`, `private`, `protected`, `rescue`, `return`, `require`, `sizeof`, `super`, `then`, `typeof`, `unless`, `until`, `when`, `while`, `with`, `yield`), Keyword, nil},
			{Words(``, `\b`, `true`, `false`, `nil`), KeywordConstant, nil},
			{`(module|lib)(\s+)([a-zA-Z_]\w*(?:::[a-zA-Z_]\w*)*)`, ByGroups(Keyword, Text, NameNamespace), nil},
			{`(def|fun|macro)(\s+)((?:[a-zA-Z_]\w*::)*)`, ByGroups(Keyword, Text, NameNamespace), Push("funcname")},
			{"def(?=[*%&^`~+-/\\[<>=])", Keyword, Push("funcname")},
			{`(class|struct|union|type|alias|enum)(\s+)((?:[a-zA-Z_]\w*::)*)`, ByGroups(Keyword, Text, NameNamespace), Push("classname")},
			{`(self|out|uninitialized)\b|(is_a|responds_to)\?`, KeywordPseudo, nil},
			{Words(``, `\b`, `debugger`, `record`, `pp`, `assert_responds_to`, `spawn`, `parallel`, `getter`, `setter`, `property`, `delegate`, `def_hash`, `def_equals`, `def_equals_and_hash`, `forward_missing_to`), NameBuiltinPseudo, nil},
			{`getter[!?]|property[!?]|__(DIR|FILE|LINE)__\b`, NameBuiltinPseudo, nil},
			{Words(`(?<!\.)`, `\b`, `Object`, `Value`, `Struct`, `Reference`, `Proc`, `Class`, `Nil`, `Symbol`, `Enum`, `Void`, `Bool`, `Number`, `Int`, `Int8`, `Int16`, `Int32`, `Int64`, `UInt8`, `UInt16`, `UInt32`, `UInt64`, `Float`, `Float32`, `Float64`, `Char`, `String`, `Pointer`, `Slice`, `Range`, `Exception`, `Regex`, `Mutex`, `StaticArray`, `Array`, `Hash`, `Set`, `Tuple`, `Deque`, `Box`, `Process`, `File`, `Dir`, `Time`, `Channel`, `Concurrent`, `Scheduler`, `abort`, `at_exit`, `caller`, `delay`, `exit`, `fork`, `future`, `get_stack_top`, `gets`, `lazy`, `loop`, `main`, `p`, `print`, `printf`, `puts`, `raise`, `rand`, `read_line`, `sleep`, `sprintf`, `system`, `with_color`), NameBuiltin, nil},
			{"(?<!\\w)(<<-?)([\"`\\']?)([a-zA-Z_]\\w*)(\\2)(.*?\\n)", StringHeredoc, nil},
			{`(<<-?)("|\')()(\2)(.*?\n)`, StringHeredoc, nil},
			{`__END__`, CommentPreproc, Push("end-part")},
			{`(?:^|(?<=[=<>~!:])|(?<=(?:\s|;)when\s)|(?<=(?:\s|;)or\s)|(?<=(?:\s|;)and\s)|(?<=\.index\s)|(?<=\.scan\s)|(?<=\.sub\s)|(?<=\.sub!\s)|(?<=\.gsub\s)|(?<=\.gsub!\s)|(?<=\.match\s)|(?<=(?:\s|;)if\s)|(?<=(?:\s|;)elsif\s)|(?<=^when\s)|(?<=^index\s)|(?<=^scan\s)|(?<=^sub\s)|(?<=^gsub\s)|(?<=^sub!\s)|(?<=^gsub!\s)|(?<=^match\s)|(?<=^if\s)|(?<=^elsif\s))(\s*)(/)`, ByGroups(Text, LiteralStringRegex), Push("multiline-regex")},
			{`(?<=\(|,|\[)/`, LiteralStringRegex, Push("multiline-regex")},
			{`(\s+)(/)(?![\s=])`, ByGroups(Text, LiteralStringRegex), Push("multiline-regex")},
			{`(0o[0-7]+(?:_[0-7]+)*(?:_?[iu][0-9]+)?)\b(\s*)([/?])?`, ByGroups(LiteralNumberOct, Text, Operator), nil},
			{`(0x[0-9A-Fa-f]+(?:_[0-9A-Fa-f]+)*(?:_?[iu][0-9]+)?)\b(\s*)([/?])?`, ByGroups(LiteralNumberHex, Text, Operator), nil},
			{`(0b[01]+(?:_[01]+)*(?:_?[iu][0-9]+)?)\b(\s*)([/?])?`, ByGroups(LiteralNumberBin, Text, Operator), nil},
			{`((?:0(?![0-9])|[1-9][\d_]*)(?:\.\d[\d_]*)(?:e[+-]?[0-9]+)?(?:_?f[0-9]+)?)(\s*)([/?])?`, ByGroups(LiteralNumberFloat, Text, Operator), nil},
			{`((?:0(?![0-9])|[1-9][\d_]*)(?:\.\d[\d_]*)?(?:e[+-]?[0-9]+)(?:_?f[0-9]+)?)(\s*)([/?])?`, ByGroups(LiteralNumberFloat, Text, Operator), nil},
			{`((?:0(?![0-9])|[1-9][\d_]*)(?:\.\d[\d_]*)?(?:e[+-]?[0-9]+)?(?:_?f[0-9]+))(\s*)([/?])?`, ByGroups(LiteralNumberFloat, Text, Operator), nil},
			{`(0\b|[1-9][\d]*(?:_\d+)*(?:_?[iu][0-9]+)?)\b(\s*)([/?])?`, ByGroups(LiteralNumberInteger, Text, Operator), nil},
			{`@@[a-zA-Z_]\w*`, NameVariableClass, nil},
			{`@[a-zA-Z_]\w*`, NameVariableInstance, nil},
			{`\$\w+`, NameVariableGlobal, nil},
			{"\\$[!@&`\\'+~=/\\\\,;.<>_*$?:\"^-]", NameVariableGlobal, nil},
			{`\$-[0adFiIlpvw]`, NameVariableGlobal, nil},
			{`::`, Operator, nil},
			Include("strings"),
			{`\?(\\[MC]-)*(\\([\\befnrtv#"\']|x[a-fA-F0-9]{1,2}|[0-7]{1,3})|\S)(?!\w)`, LiteralStringChar, nil},
			{`[A-Z][A-Z_]+\b`, NameConstant, nil},
			{`\{%`, LiteralStringInterpol, Push("in-macro-control")},
			{`\{\{`, LiteralStringInterpol, Push("in-macro-expr")},
			{`(@\[)(\s*)([A-Z]\w*)`, ByGroups(Operator, Text, NameDecorator), Push("in-attr")},
			{Words(`(\.|::)`, ``, `!=`, `!~`, `!`, `%`, `&&`, `&`, `**`, `*`, `+`, `-`, `/`, `<=>`, `<<`, `<=`, `<`, `===`, `==`, `=~`, `=`, `>=`, `>>`, `>`, `[]=`, `[]?`, `[]`, `^`, `||`, `|`, `~`), ByGroups(Operator, NameOperator), nil},
			{"(\\.|::)([a-zA-Z_]\\w*[!?]?|[*%&^`~+\\-/\\[<>=])", ByGroups(Operator, Name), nil},
			{`[a-zA-Z_]\w*(?:[!?](?!=))?`, Name, nil},
			{`(\[|\]\??|\*\*|<=>?|>=|<<?|>>?|=~|===|!~|&&?|\|\||\.{1,3})`, Operator, nil},
			{`[-+/*%=<>&!^|~]=?`, Operator, nil},
			{`[(){};,/?:\\]`, Punctuation, nil},
			{`\s+`, Text, nil},
		},
		"funcname": {
			{"(?:([a-zA-Z_]\\w*)(\\.))?([a-zA-Z_]\\w*[!?]?|\\*\\*?|[-+]@?|[/%&|^`~]|\\[\\]=?|<<|>>|<=?>|>=?|===?)", ByGroups(NameClass, Operator, NameFunction), Pop(1)},
			Default(Pop(1)),
		},
		"classname": {
			{`[A-Z_]\w*`, NameClass, nil},
			{`(\()(\s*)([A-Z_]\w*)(\s*)(\))`, ByGroups(Punctuation, Text, NameClass, Text, Punctuation), nil},
			Default(Pop(1)),
		},
		"in-intp": {
			{`\{`, LiteralStringInterpol, Push()},
			{`\}`, LiteralStringInterpol, Pop(1)},
			Include("root"),
		},
		"string-intp": {
			{`#\{`, LiteralStringInterpol, Push("in-intp")},
		},
		"string-escaped": {
			{`\\([\\befnstv#"\']|x[a-fA-F0-9]{1,2}|[0-7]{1,3})`, LiteralStringEscape, nil},
		},
		"string-intp-escaped": {
			Include("string-intp"),
			Include("string-escaped"),
		},
		"interpolated-regex": {
			Include("string-intp"),
			{`[\\#]`, LiteralStringRegex, nil},
			{`[^\\#]+`, LiteralStringRegex, nil},
		},
		"interpolated-string": {
			Include("string-intp"),
			{`[\\#]`, LiteralStringOther, nil},
			{`[^\\#]+`, LiteralStringOther, nil},
		},
		"multiline-regex": {
			Include("string-intp"),
			{`\\\\`, LiteralStringRegex, nil},
			{`\\/`, LiteralStringRegex, nil},
			{`[\\#]`, LiteralStringRegex, nil},
			{`[^\\/#]+`, LiteralStringRegex, nil},
			{`/[imsx]*`, LiteralStringRegex, Pop(1)},
		},
		"end-part": {
			{`.+`, CommentPreproc, Pop(1)},
		},
		"in-macro-control": {
			{`\{%`, LiteralStringInterpol, Push()},
			{`%\}`, LiteralStringInterpol, Pop(1)},
			{`for\b|in\b`, Keyword, nil},
			Include("root"),
		},
		"in-macro-expr": {
			{`\{\{`, LiteralStringInterpol, Push()},
			{`\}\}`, LiteralStringInterpol, Pop(1)},
			Include("root"),
		},
		"in-attr": {
			{`\[`, Operator, Push()},
			{`\]`, Operator, Pop(1)},
			Include("root"),
		},
		"strings": {
			{`\:@{0,2}[a-zA-Z_]\w*[!?]?`, LiteralStringSymbol, nil},
			{Words(`\:@{0,2}`, ``, `!=`, `!~`, `!`, `%`, `&&`, `&`, `**`, `*`, `+`, `-`, `/`, `<=>`, `<<`, `<=`, `<`, `===`, `==`, `=~`, `=`, `>=`, `>>`, `>`, `[]=`, `[]?`, `[]`, `^`, `||`, `|`, `~`), LiteralStringSymbol, nil},
			{`:'(\\\\|\\'|[^'])*'`, LiteralStringSymbol, nil},
			{`'(\\\\|\\'|[^']|\\[^'\\]+)'`, LiteralStringChar, nil},
			{`:"`, LiteralStringSymbol, Push("simple-sym")},
			{`([a-zA-Z_]\w*)(:)(?!:)`, ByGroups(LiteralStringSymbol, Punctuation), nil},
			{`"`, LiteralStringDouble, Push("simple-string")},
			{"(?<!\\.)`", LiteralStringBacktick, Push("simple-backtick")},
			{`%\{`, LiteralStringOther, Push("cb-intp-string")},
			{`%[wi]\{`, LiteralStringOther, Push("cb-string")},
			{`%r\{`, LiteralStringRegex, Push("cb-regex")},
			{`%\[`, LiteralStringOther, Push("sb-intp-string")},
			{`%[wi]\[`, LiteralStringOther, Push("sb-string")},
			{`%r\[`, LiteralStringRegex, Push("sb-regex")},
			{`%\(`, LiteralStringOther, Push("pa-intp-string")},
			{`%[wi]\(`, LiteralStringOther, Push("pa-string")},
			{`%r\(`, LiteralStringRegex, Push("pa-regex")},
			{`%<`, LiteralStringOther, Push("ab-intp-string")},
			{`%[wi]<`, LiteralStringOther, Push("ab-string")},
			{`%r<`, LiteralStringRegex, Push("ab-regex")},
			{`(%r([\W_]))((?:\\\2|(?!\2).)*)(\2[imsx]*)`, String, nil},
			{`(%[wi]([\W_]))((?:\\\2|(?!\2).)*)(\2)`, String, nil},
			{`(?<=[-+/*%=<>&!^|~,(])(\s*)(%([\t ])(?:(?:\\\3|(?!\3).)*)\3)`, ByGroups(Text, LiteralStringOther, None), nil},
			{`^(\s*)(%([\t ])(?:(?:\\\3|(?!\3).)*)\3)`, ByGroups(Text, LiteralStringOther, None), nil},
			{`(%([\[{(<]))((?:\\\2|(?!\2).)*)(\2)`, String, nil},
		},
		"simple-string": {
			Include("string-intp-escaped"),
			{`[^\\"#]+`, LiteralStringDouble, nil},
			{`[\\#]`, LiteralStringDouble, nil},
			{`"`, LiteralStringDouble, Pop(1)},
		},
		"simple-sym": {
			Include("string-escaped"),
			{`[^\\"#]+`, LiteralStringSymbol, nil},
			{`[\\#]`, LiteralStringSymbol, nil},
			{`"`, LiteralStringSymbol, Pop(1)},
		},
		"simple-backtick": {
			Include("string-intp-escaped"),
			{"[^\\\\`#]+", LiteralStringBacktick, nil},
			{`[\\#]`, LiteralStringBacktick, nil},
			{"`", LiteralStringBacktick, Pop(1)},
		},
		"cb-intp-string": {
			{`\\[\{]`, LiteralStringOther, nil},
			{`\{`, LiteralStringOther, Push()},
			{`\}`, LiteralStringOther, Pop(1)},
			Include("string-intp-escaped"),
			{`[\\#{}]`, LiteralStringOther, nil},
			{`[^\\#{}]+`, LiteralStringOther, nil},
		},
		"cb-string": {
			{`\\[\\{}]`, LiteralStringOther, nil},
			{`\{`, LiteralStringOther, Push()},
			{`\}`, LiteralStringOther, Pop(1)},
			{`[\\#{}]`, LiteralStringOther, nil},
			{`[^\\#{}]+`, LiteralStringOther, nil},
		},
		"cb-regex": {
			{`\\[\\{}]`, LiteralStringRegex, nil},
			{`\{`, LiteralStringRegex, Push()},
			{`\}[imsx]*`, LiteralStringRegex, Pop(1)},
			Include("string-intp"),
			{`[\\#{}]`, LiteralStringRegex, nil},
			{`[^\\#{}]+`, LiteralStringRegex, nil},
		},
		"sb-intp-string": {
			{`\\[\[]`, LiteralStringOther, nil},
			{`\[`, LiteralStringOther, Push()},
			{`\]`, LiteralStringOther, Pop(1)},
			Include("string-intp-escaped"),
			{`[\\#\[\]]`, LiteralStringOther, nil},
			{`[^\\#\[\]]+`, LiteralStringOther, nil},
		},
		"sb-string": {
			{`\\[\\\[\]]`, LiteralStringOther, nil},
			{`\[`, LiteralStringOther, Push()},
			{`\]`, LiteralStringOther, Pop(1)},
			{`[\\#\[\]]`, LiteralStringOther, nil},
			{`[^\\#\[\]]+`, LiteralStringOther, nil},
		},
		"sb-regex": {
			{`\\[\\\[\]]`, LiteralStringRegex, nil},
			{`\[`, LiteralStringRegex, Push()},
			{`\][imsx]*`, LiteralStringRegex, Pop(1)},
			Include("string-intp"),
			{`[\\#\[\]]`, LiteralStringRegex, nil},
			{`[^\\#\[\]]+`, LiteralStringRegex, nil},
		},
		"pa-intp-string": {
			{`\\[\(]`, LiteralStringOther, nil},
			{`\(`, LiteralStringOther, Push()},
			{`\)`, LiteralStringOther, Pop(1)},
			Include("string-intp-escaped"),
			{`[\\#()]`, LiteralStringOther, nil},
			{`[^\\#()]+`, LiteralStringOther, nil},
		},
		"pa-string": {
			{`\\[\\()]`, LiteralStringOther, nil},
			{`\(`, LiteralStringOther, Push()},
			{`\)`, LiteralStringOther, Pop(1)},
			{`[\\#()]`, LiteralStringOther, nil},
			{`[^\\#()]+`, LiteralStringOther, nil},
		},
		"pa-regex": {
			{`\\[\\()]`, LiteralStringRegex, nil},
			{`\(`, LiteralStringRegex, Push()},
			{`\)[imsx]*`, LiteralStringRegex, Pop(1)},
			Include("string-intp"),
			{`[\\#()]`, LiteralStringRegex, nil},
			{`[^\\#()]+`, LiteralStringRegex, nil},
		},
		"ab-intp-string": {
			{`\\[<]`, LiteralStringOther, nil},
			{`<`, LiteralStringOther, Push()},
			{`>`, LiteralStringOther, Pop(1)},
			Include("string-intp-escaped"),
			{`[\\#<>]`, LiteralStringOther, nil},
			{`[^\\#<>]+`, LiteralStringOther, nil},
		},
		"ab-string": {
			{`\\[\\<>]`, LiteralStringOther, nil},
			{`<`, LiteralStringOther, Push()},
			{`>`, LiteralStringOther, Pop(1)},
			{`[\\#<>]`, LiteralStringOther, nil},
			{`[^\\#<>]+`, LiteralStringOther, nil},
		},
		"ab-regex": {
			{`\\[\\<>]`, LiteralStringRegex, nil},
			{`<`, LiteralStringRegex, Push()},
			{`>[imsx]*`, LiteralStringRegex, Pop(1)},
			Include("string-intp"),
			{`[\\#<>]`, LiteralStringRegex, nil},
			{`[^\\#<>]+`, LiteralStringRegex, nil},
		},
	},
))
