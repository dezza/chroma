package lexers_test

import (
	"io/ioutil"
	"testing"

	"github.com/dezza/assert"

	"github.com/dezza/chroma/formatters"
	"github.com/dezza/chroma/lexers"
	"github.com/dezza/chroma/styles"
)

func TestCompileAllRegexes(t *testing.T) {
	for _, lexer := range lexers.Registry.Lexers {
		it, err := lexer.Tokenise(nil, "")
		assert.NoError(t, err, "%s failed", lexer.Config().Name)
		err = formatters.NoOp.Format(ioutil.Discard, styles.SwapOff, it)
		assert.NoError(t, err, "%s failed", lexer.Config().Name)
	}
}
