package lexers

import (
	. "github.com/dezza/chroma" // nolint
)

// Apl lexer.
var Apl = Register(MustNewLexer(
	&Config{
		Name:      "APL",
		Aliases:   []string{"apl"},
		Filenames: []string{"*.apl"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {
			{`\s+`, Text, nil},
			{`[⍝#].*$`, CommentSingle, nil},
			{`\'((\'\')|[^\'])*\'`, LiteralStringSingle, nil},
			{`"(("")|[^"])*"`, LiteralStringDouble, nil},
			{`[⋄◇()]`, Punctuation, nil},
			{`[\[\];]`, LiteralStringRegex, nil},
			{`⎕[A-Za-zΔ∆⍙][A-Za-zΔ∆⍙_¯0-9]*`, NameFunction, nil},
			{`[A-Za-zΔ∆⍙][A-Za-zΔ∆⍙_¯0-9]*`, NameVariable, nil},
			{`¯?(0[Xx][0-9A-Fa-f]+|[0-9]*\.?[0-9]+([Ee][+¯]?[0-9]+)?|¯|∞)([Jj]¯?(0[Xx][0-9A-Fa-f]+|[0-9]*\.?[0-9]+([Ee][+¯]?[0-9]+)?|¯|∞))?`, LiteralNumber, nil},
			{`[\.\\/⌿⍀¨⍣⍨⍠⍤∘]`, NameAttribute, nil},
			{`[+\-×÷⌈⌊∣|⍳?*⍟○!⌹<≤=>≥≠≡≢∊⍷∪∩~∨∧⍱⍲⍴,⍪⌽⊖⍉↑↓⊂⊃⌷⍋⍒⊤⊥⍕⍎⊣⊢⍁⍂≈⌸⍯↗]`, Operator, nil},
			{`⍬`, NameConstant, nil},
			{`[⎕⍞]`, NameVariableGlobal, nil},
			{`[←→]`, KeywordDeclaration, nil},
			{`[⍺⍵⍶⍹∇:]`, NameBuiltinPseudo, nil},
			{`[{}]`, KeywordType, nil},
		},
	},
))
