package lexers

import (
	. "github.com/dezza/chroma" // nolint
)

// Apacheconf lexer.
var Apacheconf = Register(MustNewLexer(
	&Config{
		Name:            "ApacheConf",
		Aliases:         []string{"apacheconf", "aconf", "apache"},
		Filenames:       []string{".htaccess", "apache.conf", "apache2.conf"},
		MimeTypes:       []string{"text/x-apacheconf"},
		CaseInsensitive: true,
	},
	Rules{
		"root": {
			{`\s+`, Text, nil},
			{`(#.*?)$`, Comment, nil},
			{`(<[^\s>]+)(?:(\s+)(.*?))?(>)`, ByGroups(NameTag, Text, LiteralString, NameTag), nil},
			{`([a-z]\w*)(\s+)`, ByGroups(NameBuiltin, Text), Push("value")},
			{`\.+`, Text, nil},
		},
		"value": {
			{`\\\n`, Text, nil},
			{`$`, Text, Pop(1)},
			{`\\`, Text, nil},
			{`[^\S\n]+`, Text, nil},
			{`\d+\.\d+\.\d+\.\d+(?:/\d+)?`, LiteralNumber, nil},
			{`\d+`, LiteralNumber, nil},
			{`/([a-z0-9][\w./-]+)`, LiteralStringOther, nil},
			{`(on|off|none|any|all|double|email|dns|min|minimal|os|productonly|full|emerg|alert|crit|error|warn|notice|info|debug|registry|script|inetd|standalone|user|group)\b`, Keyword, nil},
			{`"([^"\\]*(?:\\.[^"\\]*)*)"`, LiteralStringDouble, nil},
			{`[^\s"\\]+`, Text, nil},
		},
	},
))
