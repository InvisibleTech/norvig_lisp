package lexer

import (
	"unicode"
	"unicode/utf8"

	"github.com/InvisibleTech/norvig_lisp/token"
)

type Lexer struct {
	input           string
	inputPosition   int
	readingPosition int
	ch              rune
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readingPosition >= len(l.input) {
		l.ch = 0
	} else {
		currentCh, width := utf8.DecodeRuneInString(l.input[l.readingPosition:])
		l.ch = currentCh
		l.inputPosition = l.readingPosition
		l.readingPosition += width
	}
}

func (l *Lexer) peekChar() rune {
	if l.readingPosition >= len(l.input) {
		return 0
	} else {
		currentCh, _ := utf8.DecodeRuneInString(l.input[l.readingPosition:])
		return currentCh
	}
}

func (l *Lexer) Tokenize() []token.Token {
	var tokens []token.Token
	var tok token.Token = l.NextToken()
	for tok.Type != token.EOF {
		tokens = append(tokens, tok)
		tok = l.NextToken()
	}

	return tokens
}

func (l *Lexer) NextToken() token.Token {
	var currentToken token.Token
	l.skipWhitespace()

	switch l.ch {
	case '(':
		currentToken = newToken(token.LPAREN, string(l.ch))
		l.readChar()
	case ')':
		currentToken = newToken(token.RPAREN, string(l.ch))
		l.readChar()
	case 0:
		currentToken = newToken(token.EOF, "")
	default:
		if unicode.IsLetter(l.ch) {
			currentToken = newToken(token.IDENTIFIER, l.readIdentifier())
		} else if l.ch == '"' {
			currentToken = newToken(token.STRING, l.readString())
		} else if unicode.IsDigit(l.ch) ||
			(isSign(l.ch) && unicode.IsDigit(l.peekChar())) {
			currentToken = newToken(token.INT, l.readInteger())
		} else {
			currentToken = newToken(token.ILLEGAL, string(l.ch))
		}
	}

	return currentToken
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	startingPosition := l.inputPosition
	for unicode.IsLetter(l.ch) {
		l.readChar()
	}
	return l.input[startingPosition:l.inputPosition]
}

func (l *Lexer) readString() string {
	// Skip opening quote
	l.readChar()

	startingPosition := l.inputPosition
	for l.ch != 0 && l.ch != '"' {
		l.readChar()
	}
	endingPosition := l.inputPosition

	// Skip quote or just stay at EOF
	l.readChar()

	return l.input[startingPosition:endingPosition]
}

func (l *Lexer) readInteger() string {
	startingPosition := l.inputPosition
	if isSign(l.ch) {
		l.readChar()
	}
	for unicode.IsDigit(l.ch) {
		l.readChar()
	}
	return l.input[startingPosition:l.inputPosition]
}

func isSign(ch rune) bool {
	return ch == '+' || ch == '-'
}

func newToken(tokenType token.TokenType, value string) token.Token {
	return token.Token{Type: tokenType, Literal: value}
}
