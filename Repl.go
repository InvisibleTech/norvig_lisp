package norvig_lisp

import (
	"github.com/InvisibleTech/norvig_lisp/lexer"
	"github.com/InvisibleTech/norvig_lisp/token"
)

func repl(lines string) {
	var l *lexer.Lexer = lexer.NewLexer(lines)
	var tok token.Token = l.NextToken()
	for tok.Type != token.EOF {
		tok = l.NextToken()
	}
}
