package styles

import (
	"github.com/dezza/chroma"
)

// Pastie style.
var Pastie = Register(chroma.MustNewStyle("pastie", chroma.StyleEntries{
	chroma.TextWhitespace:        "#bbbbbb",
	chroma.Comment:               "#888888",
	chroma.CommentPreproc:        "bold #cc0000",
	chroma.CommentSpecial:        "bg:#fff0f0 bold #cc0000",
	chroma.LiteralString:         "bg:#fff0f0 #dd2200",
	chroma.LiteralStringRegex:    "bg:#fff0ff #008800",
	chroma.LiteralStringOther:    "bg:#f0fff0 #22bb22",
	chroma.LiteralStringSymbol:   "#aa6600",
	chroma.LiteralStringInterpol: "#3333bb",
	chroma.LiteralStringEscape:   "#0044dd",
	chroma.OperatorWord:          "#008800",
	chroma.Keyword:               "bold #008800",
	chroma.KeywordPseudo:         "nobold",
	chroma.KeywordType:           "#888888",
	chroma.NameClass:             "bold #bb0066",
	chroma.NameException:         "bold #bb0066",
	chroma.NameFunction:          "bold #0066bb",
	chroma.NameProperty:          "bold #336699",
	chroma.NameNamespace:         "bold #bb0066",
	chroma.NameBuiltin:           "#003388",
	chroma.NameVariable:          "#336699",
	chroma.NameVariableClass:     "#336699",
	chroma.NameVariableInstance:  "#3333bb",
	chroma.NameVariableGlobal:    "#dd7700",
	chroma.NameConstant:          "bold #003366",
	chroma.NameTag:               "bold #bb0066",
	chroma.NameAttribute:         "#336699",
	chroma.NameDecorator:         "#555555",
	chroma.NameLabel:             "italic #336699",
	chroma.LiteralNumber:         "bold #0000DD",
	chroma.GenericHeading:        "#333",
	chroma.GenericSubheading:     "#666",
	chroma.GenericDeleted:        "bg:#ffdddd #000000",
	chroma.GenericInserted:       "bg:#ddffdd #000000",
	chroma.GenericError:          "#aa0000",
	chroma.GenericEmph:           "italic",
	chroma.GenericStrong:         "bold",
	chroma.GenericPrompt:         "#555555",
	chroma.GenericOutput:         "#888888",
	chroma.GenericTraceback:      "#aa0000",
	chroma.Error:                 "bg:#e3d2d2 #a61717",
	chroma.Background:            " bg:#ffffff",
}))
