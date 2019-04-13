package lexer

import (
	"testing"

	"github.com/InvisibleTech/norvig_lisp/token"
)

func TestNextToken(t *testing.T) {
	input := `5
	10
		( x y )
		add
		( five +5 -5 -10



		)
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INT, "5"},
		{token.INT, "10"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "five"},
		{token.INT, "+5"},
		{token.INT, "-5"},
		{token.INT, "-10"},
		{token.RPAREN, ")"},
		{token.EOF, ""},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestTokenize(t *testing.T) {
	input := `5
	10
		( x y )
		add
		( five +5 -5 -10



		)
	`

	l := NewLexer(input)

	tokens := l.Tokenize()
	if len(tokens) != 13 {
		t.Fatalf("Tokenize did not return expected length slice, expected=%d, got %d",
			13, len(tokens))
	}
}
