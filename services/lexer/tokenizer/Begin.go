package tokenizer

import (
	"strings"
	"unicode"

	"github.com/InvisibleTech/norvig_lisp/services/lexer/token"
)

func Begin(tokenizer *Tokenizer) TokenizerFunc {
	tokenizer.SkipWhitespace()

	if strings.HasPrefix(tokenizer.InputToEnd(), token.LeftParen) {
		return LeftParen
	}

	thing1, thing2 := tokenizer.TwinPeeks()
	if unicode.IsDigit(thing1) ||
		((thing1 == '+' || thing1 == '-') && unicode.IsDigit(thing2)) {
		return Number
	}

	return Symbol
}
