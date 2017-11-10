package styles

import (
	"github.com/alecthomas/chroma"
)

// Dark Low Contrast

// Lucius style.
var Lucius = Register(chroma.MustNewStyle("lucius", chroma.StyleEntries{
	chroma.Background:         "#bcbcbc bg:#303030",
	chroma.Comment:            "#9e9e9e",
	chroma.CommentSpecial:     "#afaf00 bg:#5f5f00",
	chroma.Keyword:            "#5fafd7", // Duplicates marked with *K
	chroma.KeywordDeclaration: "#5fafd7", // *K
	chroma.KeywordNamespace:   "#5fafd7", // *K
	chroma.KeywordType:        "#5fafaf", // Type
	chroma.Operator:           "#5fafd7", // *K
	chroma.OperatorWord:       "#5fafd7", // *K
	chroma.NameClass:          "#5fafd7", // *K
	chroma.NameBuiltin:        "#5fafd7", // *K
	chroma.NameException:      "#5fafd7", // *K
	chroma.NameVariable:       "#5fafd7", // *K
	chroma.LiteralString:      "#afaf87", // Literal
	chroma.LiteralNumber:      "#afaf87", // Literal
	//chroma.GenericHeading:     "bold #00afd7",      //!
	chroma.GenericSubheading: "bold #800080",      //!
	chroma.GenericDeleted:    "#875f5f",           // DiffDelete (KeywordNamespace, Operator)
	chroma.GenericInserted:   "#5f875f",           // DiffAdd (NameException)
	chroma.GenericError:      "underline #ff5f5f", //! Spellbad ()
	chroma.GenericEmph:       "italic",            //! SpellCap (not sure)
	chroma.GenericStrong:     "bold",              // SpellLocal
	//chroma.GenericPrompt:      "bold #5fafd7",      //! SpellCap (not sure)
	//chroma.GenericOutput:    "#bcbcbc", // shell prompt
	//chroma.GenericTraceback: "#870000", // error traceback
	chroma.Error: "#d75f5f bg:#870000", // generic error
}))

//chroma.Error:              "border:#d75f5f bg:#870000",
