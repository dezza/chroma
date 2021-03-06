package lexers

import (
	. "github.com/dezza/chroma" // nolint
)

// Lighttpd Configuration File lexer.
var Lighttpd = Register(MustNewLexer(
	&Config{
		Name:      "Lighttpd configuration file",
		Aliases:   []string{"lighty", "lighttpd"},
		Filenames: []string{},
		MimeTypes: []string{"text/x-lighttpd-conf"},
	},
	Rules{
		"root": {
			{`#.*\n`, CommentSingle, nil},
			{`/\S*`, Name, nil},
			{`[a-zA-Z._-]+`, Keyword, nil},
			{`\d+\.\d+\.\d+\.\d+(?:/\d+)?`, LiteralNumber, nil},
			{`[0-9]+`, LiteralNumber, nil},
			{`=>|=~|\+=|==|=|\+`, Operator, nil},
			{`\$[A-Z]+`, NameBuiltin, nil},
			{`[(){}\[\],]`, Punctuation, nil},
			{`"([^"\\]*(?:\\.[^"\\]*)*)"`, LiteralStringDouble, nil},
			{`\s+`, Text, nil},
		},
	},
))
