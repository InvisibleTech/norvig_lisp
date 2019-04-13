package norvig_lisp

import (
	"github.com/InvisibleTech/norvig_lisp/lexer"
)

func repl(lines string) {
	var l *lexer.Lexer = lexer.NewLexer(lines)
	l.Tokenize()
}
